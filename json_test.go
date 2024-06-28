package sdkclient_test

import (
	"testing"

	sdkclient "github.com/fujiwara/awslim"
	"github.com/google/go-cmp/cmp"
)

var jsonTestCases = []struct {
	Name      string
	Src       any
	Expect    string
	CamelCase bool
}{
	{
		Name:      "no-camel",
		Src:       struct{ FooBar string }{FooBar: "baz"},
		Expect:    `{"FooBar":"baz"}`,
		CamelCase: false,
	},
	{
		Name:      "simple",
		Src:       struct{ FooBar string }{FooBar: "baz"},
		Expect:    `{"fooBar":"baz"}`,
		CamelCase: true,
	},
	{
		Name:      "nested",
		Src:       struct{ FooBar struct{ BazQux string } }{FooBar: struct{ BazQux string }{BazQux: "quux"}},
		Expect:    `{"fooBar":{"bazQux":"quux"}}`,
		CamelCase: true,
	},
	{
		Name:      "array",
		Src:       []any{struct{ FooBar string }{FooBar: "baz"}, 0, 1, true, nil},
		Expect:    `[{"fooBar":"baz"},0,1,true,null]`,
		CamelCase: true,
	},
	{
		Name:      "string",
		Src:       "FooBar",
		Expect:    `"FooBar"`,
		CamelCase: true,
	},
	{
		Name:      "number",
		Src:       42,
		Expect:    `42`,
		CamelCase: true,
	},
	{
		Name:      "true",
		Src:       true,
		Expect:    `true`,
		CamelCase: true,
	},
}

func TestMarshalJSON(t *testing.T) {
	for _, tc := range jsonTestCases {
		t.Run(tc.Name, func(t *testing.T) {
			b, err := sdkclient.MarshalJSON(tc.Src, tc.CamelCase)
			if err != nil {
				t.Fatalf("failed to unmarshal: %v", err)
			}
			if diff := cmp.Diff(tc.Expect, string(b)); diff != "" {
				t.Errorf("expect match (-got +want):\n%s", diff)
			}
		})
	}
}
