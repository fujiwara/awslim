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
	"github.com/jmespath/go-jmespath"
)

var clientMethods = make(map[string]ClientMethod)

type ClientMethod func(context.Context, aws.Config, json.RawMessage) (any, error)

type CLI struct {
	Service string `arg:"" help:"service name" default:""`
	Method  string `arg:"" help:"method name" default:""`
	Input   string `arg:"" help:"input JSON" default:"{}"`
	Compact bool   `short:"c" help:"compact JSON output"`
	Query   string `short:"q" help:"JMESPath query to apply to output"`

	w io.Writer
}

func Run(ctx context.Context) error {
	var c CLI
	c.w = os.Stdout
	kong.Parse(&c)
	return c.Dispatch(ctx)
}

func (c *CLI) Dispatch(ctx context.Context) error {
	if c.Service == "" {
		return c.ListServices(ctx)
	} else if c.Method == "" {
		return c.ListMethods(ctx)
	} else {
		return c.CallMethod(ctx)
	}
}

func (c *CLI) SetWriter(w io.Writer) {
	c.w = w
}

func (c *CLI) CallMethod(ctx context.Context) error {
	method := kebabToCamel(c.Method)
	key := buildKey(c.Service, method)
	if c.Input == "help" {
		fmt.Fprintf(c.w, "See https://pkg.go.dev/github.com/aws/aws-sdk-go-v2/service/%s\n", key)
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

	if c.Query != "" {
		out, err = jmespath.Search(c.Query, out)
		if err != nil {
			return fmt.Errorf("failed to apply JMESPath query: %w", err)
		}
	}

	b, err := json.Marshal(out)
	if err != nil {
		return fmt.Errorf("failed to marshal response: %w", err)
	}
	if !c.Compact {
		var buf bytes.Buffer
		json.Indent(&buf, b, "", "  ")
		buf.WriteString("\n")
		buf.WriteTo(c.w)
	} else {
		io.WriteString(c.w, string(b))
	}
	return nil
}

func (c *CLI) ListMethods(_ context.Context) error {
	methods := make([]string, 0)
	for name := range clientMethods {
		service, method := parseKey(name)
		if service == c.Service {
			methods = append(methods, method)
		}
	}
	sort.Strings(methods)
	for _, name := range methods {
		fmt.Fprintln(c.w, name)
	}
	return nil
}

func (c *CLI) ListServices(_ context.Context) error {
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
		fmt.Fprintln(c.w, name)
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

func kebabToCamel(kebab string) string {
	parts := strings.Split(kebab, "-")
	results := make([]string, 0, len(parts))
	for _, p := range parts {
		if len(p) == 0 {
			continue
		}
		results = append(results, strings.ToUpper(p[:1])+p[1:])
	}
	return strings.Join(results, "")
}

//go:generate go run cmd/aws-sdk-client-gen/main.go cmd/aws-sdk-client-gen/gen.go
