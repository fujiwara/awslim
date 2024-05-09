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

var clientMethods = make(map[string]func(context.Context, aws.Config, json.RawMessage) (any, error))

func Run(ctx context.Context) error {
	var err error
	awsCfg, err := config.LoadDefaultConfig(ctx)
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
		return dispatchMethod(ctx, awsCfg, pkgName, methodName, input)
	case 4:
		pkgName := os.Args[1]
		methodName := os.Args[2]
		input := json.RawMessage(os.Args[3])
		return dispatchMethod(ctx, awsCfg, pkgName, methodName, input)
	default:
		return fmt.Errorf("too many args")
	}
}

func dispatchMethod(ctx context.Context, awsCfg aws.Config, pkgName, methodName string, in json.RawMessage) error {
	key := buildKey(pkgName, methodName)
	if bytes.Equal(in, []byte(`help`)) {
		fmt.Printf("See https://pkg.go.dev/github.com/aws/aws-sdk-go-v2/service/%s\n", key)
		return nil
	}
	fn := clientMethods[key]
	if fn == nil {
		return fmt.Errorf("unknown function %s", key)
	}
	out, err := fn(ctx, awsCfg, in)
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
		service, method := parseKey(name)
		if service == pkgName {
			methods = append(methods, method)
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
		service, _ := parseKey(name)
		services[service] = struct{}{}
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

func parseKey(key string) (string, string) {
	parts := strings.Split(key, "#")
	service := parts[0]
	method := strings.SplitN(parts[1], ".", 2)[1]
	return service, method
}

func buildKey(service, method string) string {
	return fmt.Sprintf("%s#Client.%s", service, method)
}

//go:generate go run cmd/aws-sdk-client-gen/main.go cmd/aws-sdk-client-gen/gen.go
