package sdkclient_test

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"testing"

	"github.com/alecthomas/kong"
	sdkclient "github.com/fujiwara/awslim"
)

type PagingOutput struct {
	Next string `json:"Next,omitempty"`
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
		Expect: "Get\nList\n",
	},
	{
		Name:   "list methods of bar",
		Args:   []string{"bar"},
		Expect: "Get\nList\n",
	},
	{
		Name:   "list methods of baz",
		Args:   []string{"baz"},
		Expect: "Echo\nPaging\n",
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
}

func TestRun(t *testing.T) {
	for _, tc := range TestCases {
		ctx := context.Background()
		t.Run(tc.Name, func(t *testing.T) {
			for k, v := range tc.Env {
				t.Setenv(k, v)
			}
			buf := &bytes.Buffer{}
			c, err := newCLI(tc.Args, buf)
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

func newCLI(args []string, w io.Writer) (*sdkclient.CLI, error) {
	c := &sdkclient.CLI{}
	c.SetWriter(w)
	p, err := kong.New(c)
	if err != nil {
		return nil, err
	}
	_, err = p.Parse(args)
	if err != nil {
		return nil, err
	}
	return c, nil
}
