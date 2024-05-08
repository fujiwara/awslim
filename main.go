package sdkclient

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

var awsCfg aws.Config

var clientMethods = make(map[string]func(context.Context, json.RawMessage) (any, error))

func Run(ctx context.Context) error {
	var err error
	awsCfg, err = config.LoadDefaultConfig(ctx)
	if err != nil {
		return err
	}
	switch len(os.Args) {
	case 1:
		return listServices()
	case 2:
		pkgName := os.Args[1]
		return listMethods(pkgName)
	case 3:
		pkgName := os.Args[1]
		methodName := os.Args[2]
		input := json.RawMessage(`{}`)
		return dispatchMethod(ctx, pkgName, methodName, input)
	case 4:
		pkgName := os.Args[1]
		methodName := os.Args[2]
		input := json.RawMessage(os.Args[3])
		return dispatchMethod(ctx, pkgName, methodName, input)
	default:
		return fmt.Errorf("too many args")
	}
}

func dispatchMethod(ctx context.Context, pkgName, methodName string, in json.RawMessage) error {
	fn := clientMethods[pkgName+"_"+methodName]
	if fn == nil {
		return fmt.Errorf("unknown method %s of %s", methodName, pkgName)
	}
	out, err := fn(ctx, in)
	if err != nil {
		return err
	}
	b, err := json.Marshal(out)
	if err != nil {
		return fmt.Errorf("failed to marshal response: %w", err)
	}
	var buf bytes.Buffer
	json.Indent(&buf, b, "", "  ")
	buf.WriteTo(os.Stdout)
	fmt.Fprintln(os.Stdout)
	return nil
}

func listMethods(pkgName string) error {
	methods := make([]string, 0)
	for name := range clientMethods {
		if strings.HasPrefix(name, pkgName+"_") {
			methods = append(methods, strings.TrimPrefix(name, pkgName+"_"))
		}
	}
	sort.Strings(methods)
	for _, name := range methods {
		fmt.Println(name)
	}
	return nil
}

func listServices() error {
	services := make(map[string]struct{})
	for name := range clientMethods {
		services[strings.Split(name, "_")[0]] = struct{}{}
	}
	names := make([]string, 0)
	for name := range services {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Println(name)
	}
	return nil
}

//go:generate go run cmd/aws-sdk-client-gen/main.go cmd/aws-sdk-client-gen/gen.go
