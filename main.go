package sdkclient

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"

	"github.com/alecthomas/kong"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

var clientMethods = make(map[string]func(context.Context, aws.Config, json.RawMessage) (any, error))

type CLI struct {
	Service string `arg:"" help:"service name" default:""`
	Method  string `arg:"" help:"method name" default:""`
	Input   string `arg:"" help:"input JSON" default:"{}"`
	Compact bool   `short:"c" help:"compact JSON output"`
}

func Run(ctx context.Context) error {
	var c CLI
	kong.Parse(&c)
	if c.Service == "" {
		return c.listServices(ctx)
	} else if c.Method == "" {
		return c.listMethods(ctx)
	} else {
		return c.dispatchMethod(ctx)
	}
}

func (c *CLI) dispatchMethod(ctx context.Context) error {
	key := buildKey(c.Service, c.Method)
	if c.Input == "help" {
		fmt.Printf("See https://pkg.go.dev/github.com/aws/aws-sdk-go-v2/service/%s\n", key)
		return nil
	}

	fn := clientMethods[key]
	if fn == nil {
		return fmt.Errorf("unknown function %s", key)
	}

	awsCfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return err
	}
	out, err := fn(ctx, awsCfg, json.RawMessage(c.Input))
	if err != nil {
		return err
	}
	b, err := json.Marshal(out)
	if err != nil {
		return fmt.Errorf("failed to marshal response: %w", err)
	}
	if !c.Compact {
		var buf bytes.Buffer
		json.Indent(&buf, b, "", "  ")
		buf.WriteTo(os.Stdout)
		fmt.Fprintln(os.Stdout)
	} else {
		io.WriteString(os.Stdout, string(b))
	}
	return nil
}

func (c *CLI) listMethods(_ context.Context) error {
	methods := make([]string, 0)
	for name := range clientMethods {
		service, method := parseKey(name)
		if service == c.Service {
			methods = append(methods, method)
		}
	}
	sort.Strings(methods)
	for _, name := range methods {
		fmt.Println(name)
	}
	return nil
}

func (c *CLI) listServices(_ context.Context) error {
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
