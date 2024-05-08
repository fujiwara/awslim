package main

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"

	"github.com/aws/aws-sdk-go-v2/service/ecs"
	"github.com/aws/aws-sdk-go-v2/service/firehose"
	"github.com/aws/aws-sdk-go-v2/service/kinesis"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
)

func main() {
	gen("kinesis", reflect.TypeOf(kinesis.New(kinesis.Options{})))
	gen("firehose", reflect.TypeOf(firehose.New(firehose.Options{})))
	gen("ssm", reflect.TypeOf(ssm.New(ssm.Options{})))
	gen("ecs", reflect.TypeOf(ecs.New(ecs.Options{})))
	gen("s3", reflect.TypeOf(s3.New(s3.Options{})))
}

func gen(pkgName string, clientType reflect.Type) error {
	buf := &strings.Builder{}
	fmt.Fprintln(buf, "package sdkclient")
	fmt.Fprintln(buf)
	fmt.Fprintf(buf, `import (
		"context"
		"encoding/json"
		"fmt"
		"github.com/aws/aws-sdk-go-v2/service/%s"
	)
	`, pkgName)
	var methodNames []string
	for i := 0; i < clientType.NumMethod(); i++ {
		method := clientType.Method(i)
		params := make([]string, 0)
		for j := 0; j < method.Type.NumIn(); j++ {
			params = append(params, method.Type.In(j).String())
		}
		if len(params) <= 1 {
			log.Printf("no params func %s", method.Name)
			continue
		}
		methodNames = append(methodNames, method.Name)

		fmt.Fprintf(buf, `func %s_%s(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
			svc := %s.NewFromConfig(awsCfg)
			var in %s
			if err := json.Unmarshal(b, &in); err != nil {
				return nil, fmt.Errorf("failed to unmarshal request: %%w", err)
			}
			if out, err := svc.%s(ctx, &in); err != nil {
				return nil, fmt.Errorf("failed to call %s: %%w", err)
			} else {
				o, err := json.Marshal(out)
				if err != nil {
					return nil, fmt.Errorf("failed to marshal response: %%w", err)
				}
				return o, nil
			}
		}
		`,
			pkgName,
			method.Name,
			pkgName,
			strings.TrimPrefix(params[2], "*"), // (receiver, context.Context, *Request, ...)
			method.Name,
			method.Name,
		)
		fmt.Fprintln(buf)
	}
	fmt.Fprintln(buf, `func init() {`)
	for _, name := range methodNames {
		fmt.Fprintf(buf, `clientMethods["%s_%s"] = %s_%s
		`, pkgName, name, pkgName, name)
	}
	fmt.Fprintln(buf, "}")
	f, err := os.OpenFile(pkgName+"_gen.go", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	if _, err := f.WriteString(buf.String()); err != nil {
		return err
	}
	log.Printf("generated %s_gen.go", pkgName)
	return nil
}
