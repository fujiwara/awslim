package sdkclient_test

import (
	"bytes"
	"context"
	"encoding/json"
	"testing"

	sdkclient "github.com/fujiwara/awslim"
)

type PagingOutput struct {
	Next string `json:"Next,omitempty"`
}

// MockOptions mimics a service client Options struct (e.g. s3.Options).
type MockOptions struct {
	Region           string
	BaseEndpoint     *string
	UsePathStyle     bool
	RetryMaxAttempts int
	HTTPClient       any
}

func init() {
	sdkclient.SetClientMethod("foo", "List", func(_ context.Context, _ *sdkclient.ClientMethodParam) (any, error) {
		return []string{"a", "b", "c"}, nil
	})
	sdkclient.SetClientMethod("foo", "Get", func(_ context.Context, _ *sdkclient.ClientMethodParam) (any, error) {
		return struct{ Name string }{Name: "foo"}, nil
	})
	sdkclient.SetClientMethod("bar", "List", func(_ context.Context, _ *sdkclient.ClientMethodParam) (any, error) {
		return []string{"x", "y", "z"}, nil
	})
	sdkclient.SetClientMethod("bar", "Get", func(_ context.Context, _ *sdkclient.ClientMethodParam) (any, error) {
		return struct{ Name string }{Name: "bar"}, nil
	})
	sdkclient.SetClientMethod("baz", "Echo", func(_ context.Context, p *sdkclient.ClientMethodParam) (any, error) {
		var v any
		err := json.Unmarshal(p.InputBytes, &v)
		return v, err
	})
	sdkclient.SetClientMethod("baz", "Options", func(_ context.Context, p *sdkclient.ClientMethodParam) (any, error) {
		o := MockOptions{Region: "default-region"}
		if err := p.ApplyClientOptions(&o); err != nil {
			return nil, err
		}
		if p.DryRun {
			return nil, sdkclient.ErrDryRun
		}
		return o, nil
	})
	sdkclient.SetClientMethod("foo", "Options", func(_ context.Context, p *sdkclient.ClientMethodParam) (any, error) {
		o := MockOptions{Region: "default-region"}
		if err := p.ApplyClientOptions(&o); err != nil {
			return nil, err
		}
		if p.DryRun {
			return nil, sdkclient.ErrDryRun
		}
		return o, nil
	})
	sdkclient.SetClientMethod("baz", "Paging", func(_ context.Context, p *sdkclient.ClientMethodParam) (any, error) {
		var v map[string]string
		json.Unmarshal(p.InputBytes, &v)
		switch v["Start"] {
		case "":
			return PagingOutput{Next: "1"}, nil
		case "1":
			return PagingOutput{Next: "2"}, nil
		case "2":
			return PagingOutput{Next: "3"}, nil
		}
		return PagingOutput{}, nil
	})
}

type TestCase struct {
	Name    string
	Args    []string
	Expect  string
	IsError bool
	Env     map[string]string
}

var TestCases = []TestCase{
	{
		Name:   "no args (list services)",
		Args:   []string{},
		Expect: "bar\nbaz\nfoo\n",
	},
	{
		Name:   "list methods of foo",
		Args:   []string{"foo"},
		Expect: "Get\nList\nOptions\n",
	},
	{
		Name:   "list methods of bar",
		Args:   []string{"bar"},
		Expect: "Get\nList\n",
	},
	{
		Name:   "list methods of baz",
		Args:   []string{"baz"},
		Expect: "Echo\nOptions\nPaging\n",
	},
	{
		Name:   "call foo#Client.List",
		Args:   []string{"foo", "List"},
		Expect: "[\n  \"a\",\n  \"b\",\n  \"c\"\n]\n",
	},
	{
		Name:   "call foo#Client.Get",
		Args:   []string{"foo", "Get"},
		Expect: "{\n  \"Name\": \"foo\"\n}\n",
	},
	{
		Name:   "call foo#Client.List",
		Args:   []string{"foo", "List", "help"},
		Expect: "See https://pkg.go.dev/github.com/aws/aws-sdk-go-v2/service/foo#Client.List\n",
	},
	{
		Name:   "call foo#Client.List -c",
		Args:   []string{"foo", "List", "-c"},
		Expect: `["a","b","c"]`,
	},
	{
		Name:   "call bar#Client.Get --compact",
		Args:   []string{"bar", "Get", "--compact"},
		Expect: `{"Name":"bar"}`,
	},
	{
		Name:   "call baz#Client.Echo",
		Args:   []string{"baz", "Echo", `{"Example": "value"}`},
		Expect: "{\n  \"Example\": \"value\"\n}\n",
	},
	{
		Name:   "call baz#Client.Echo Jsonnet",
		Args:   []string{"baz", "Echo", `{Example: std.extVar("value")}`, "--ext-str", "value=foo"},
		Expect: "{\n  \"Example\": \"foo\"\n}\n",
	},
	{
		Name:   "call baz#Client.Echo Jsonnet file",
		Args:   []string{"baz", "Echo", "tests/echo.jsonnet", "--ext-code", "a=1;b=2", "-c"},
		Expect: `{"Sum":3}`,
	},
	{
		Name:   "call baz#Client.Echo JMESPath",
		Args:   []string{"baz", "Echo", `{"Example": ["a","b","c"]}`, "--query", "Example[0]", "-c"},
		Expect: `"a"`,
	},
	{
		Name:   "call baz#Client.Echo raw string",
		Args:   []string{"baz", "Echo", `{"Example": "value"}`, "-q", "Example", "--raw-output"},
		Expect: "value\n",
	},
	{
		Name:   "call baz#Client.Echo raw object",
		Args:   []string{"baz", "Echo", `{"Example": "value"}`, "-r"},
		Expect: "{\n  \"Example\": \"value\"\n}\n",
	},
	{
		Name:   "call baz#Client.Paging",
		Args:   []string{"baz", "Paging", `{}`, "--follow-next", "Next=Start", "-c"},
		Expect: `{"Next":"1"}{"Next":"2"}{"Next":"3"}{}`,
	},
	{
		Name:   "call baz#Client.Echo with args",
		Args:   []string{"baz", "Echo", `{"Examples":[_(0), _(1), std.parseInt(_(2))]}`, "example1", "example2", "3", "-c"},
		Expect: `{"Examples":["example1","example2",3]}`,
	},
	{
		Name:   "call baz#Client.Echo with env",
		Args:   []string{"baz", "Echo", `{"Examples":[env('MYENV','default'),env('NOENV','default')]}`, "-c"},
		Expect: `{"Examples":["example1","default"]}`,
		Env:    map[string]string{"MYENV": "example1"}, // NOENV is not set
	},
	{
		Name:   "call baz#Client.Echo with must_env",
		Args:   []string{"baz", "Echo", `{"Example":must_env('MYENV')}`, "-c"},
		Expect: `{"Example":"example1"}`,
		Env:    map[string]string{"MYENV": "example1"},
	},
	{
		Name:    "call baz#Client.Echo with must_env is not set",
		Args:    []string{"baz", "Echo", `{"Example":must_env('MYENV')}`, "-c"},
		Env:     map[string]string{},
		IsError: true,
	},
	{
		Name:   "call baz#Client.Echo with dynamic flags",
		Args:   []string{"baz", "Echo", "--foo-Foo", "FOO", `{Baz:"baz"}`, "-c", "--bar=BAR"},
		Expect: `{"Bar":"BAR","Baz":"baz","FooFoo":"FOO"}`,
	},
	{
		Name:   "client options from config",
		Args:   []string{"baz", "Options", "{}", "-c"},
		Expect: `{"Region":"default-region","BaseEndpoint":"http://localhost:9000","UsePathStyle":true,"RetryMaxAttempts":0,"HTTPClient":null}`,
	},
	{
		Name:   "client options not defined in config",
		Args:   []string{"foo", "Options", "{}", "-c"},
		Expect: `{"Region":"default-region","BaseEndpoint":null,"UsePathStyle":false,"RetryMaxAttempts":0,"HTTPClient":null}`,
	},
	{
		Name:   "client options by flag (Jsonnet)",
		Args:   []string{"foo", "Options", "{}", "-c", "-C", `{UsePathStyle: true, RetryMaxAttempts: 1+2}`},
		Expect: `{"Region":"default-region","BaseEndpoint":null,"UsePathStyle":true,"RetryMaxAttempts":3,"HTTPClient":null}`,
	},
	{
		Name:   "client options by flag overrides config",
		Args:   []string{"baz", "Options", "{}", "-c", "--client-option", `{UsePathStyle: false, Region: "ap-northeast-1"}`},
		Expect: `{"Region":"ap-northeast-1","BaseEndpoint":"http://localhost:9000","UsePathStyle":false,"RetryMaxAttempts":0,"HTTPClient":null}`,
	},
	{
		Name:   "client options by env",
		Args:   []string{"foo", "Options", "{}", "-c"},
		Env:    map[string]string{"AWSLIM_CLIENT_OPTION": `{"UsePathStyle": true}`},
		Expect: `{"Region":"default-region","BaseEndpoint":null,"UsePathStyle":true,"RetryMaxAttempts":0,"HTTPClient":null}`,
	},
	{
		Name:    "client options unknown field (strict)",
		Args:    []string{"foo", "Options", "{}", "-c", "-C", `{Unknown: true}`},
		Expect:  "See https://pkg.go.dev/github.com/aws/aws-sdk-go-v2/service/foo#Options\n",
		IsError: true,
	},
	{
		Name:   "client options unknown field (no-strict)",
		Args:   []string{"foo", "Options", "{}", "-c", "--no-strict", "-C", `{Unknown: true, UsePathStyle: true}`},
		Expect: `{"Region":"default-region","BaseEndpoint":null,"UsePathStyle":true,"RetryMaxAttempts":0,"HTTPClient":null}`,
	},
	{
		Name:   "call baz#Client.Echo --no-strict is not a dynamic flag",
		Args:   []string{"baz", "Echo", `{Baz:"baz"}`, "-c", "--no-strict"},
		Expect: `{"Baz":"baz"}`,
	},
	{
		Name:   "flag by env (compact)",
		Args:   []string{"foo", "List"},
		Env:    map[string]string{"AWSLIM_COMPACT": "true"},
		Expect: `["a","b","c"]`,
	},
	{
		Name:   "dry-run shows client options",
		Args:   []string{"foo", "Options", `{}`, "-n", "-C", `{"UsePathStyle":true}`},
		Expect: "dry-run: foo#Client.Options will be called with:\n{}\nclient options:\n{\"UsePathStyle\":true}\n",
	},
}

func TestRun(t *testing.T) {
	t.Setenv("XDG_CONFIG_HOME", "testdata") // don't use real config
	for _, tc := range TestCases {
		ctx := context.Background()
		t.Run(tc.Name, func(t *testing.T) {
			for k, v := range tc.Env {
				t.Setenv(k, v)
			}
			buf := &bytes.Buffer{}
			c, err := newCLI(ctx, tc.Args, buf)
			if err != nil {
				t.Fatal(err)
			}
			err = c.Dispatch(ctx)
			if err != nil {
				t.Log(err)
			}
			if tc.IsError && err == nil {
				t.Fatal("expected error, but got nil")
			} else if !tc.IsError && err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got := buf.String(); got != tc.Expect {
				t.Errorf("unexpected output: got %q, expect %q", got, tc.Expect)
			}
		})
	}
}

func newCLI(ctx context.Context, args []string, out *bytes.Buffer) (*sdkclient.CLI, error) {
	c, err := sdkclient.NewCLI(ctx, args)
	if err != nil {
		return nil, err
	}
	c.SetWriter(out)
	return c, nil
}
