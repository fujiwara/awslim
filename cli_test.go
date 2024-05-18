package sdkclient_test

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"testing"

	"github.com/alecthomas/kong"
	"github.com/aws/aws-sdk-go-v2/aws"
	sdkclient "github.com/fujiwara/aws-sdk-client-go"
)

func init() {
	sdkclient.SetClientMethod("foo#Client.List", func(_ context.Context, _ aws.Config, _ json.RawMessage, _ io.Reader) (any, error) {
		return []string{"a", "b", "c"}, nil
	})
	sdkclient.SetClientMethod("foo#Client.Get", func(_ context.Context, _ aws.Config, _ json.RawMessage, _ io.Reader) (any, error) {
		return struct{ Name string }{Name: "foo"}, nil
	})
	sdkclient.SetClientMethod("bar#Client.List", func(_ context.Context, _ aws.Config, _ json.RawMessage, _ io.Reader) (any, error) {
		return []string{"x", "y", "z"}, nil
	})
	sdkclient.SetClientMethod("bar#Client.Get", func(_ context.Context, _ aws.Config, _ json.RawMessage, _ io.Reader) (any, error) {
		return struct{ Name string }{Name: "bar"}, nil
	})
	sdkclient.SetClientMethod("baz#Client.Echo", func(_ context.Context, _ aws.Config, b json.RawMessage, _ io.Reader) (any, error) {
		var v any
		err := json.Unmarshal(b, &v)
		return v, err
	})
}

type TestCase struct {
	Name    string
	Args    []string
	Expect  string
	IsError bool
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
		Expect: "Echo\n",
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
}

func TestRun(t *testing.T) {
	for _, tc := range TestCases {
		ctx := context.Background()
		t.Run(tc.Name, func(t *testing.T) {
			buf := &bytes.Buffer{}
			c, err := newCLI(tc.Args, buf)
			if err != nil {
				t.Fatal(err)
			}
			if err := c.Dispatch(ctx); err != nil {
				t.Fatal(err)
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
