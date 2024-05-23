package sdkclient_test

import (
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	sdkclient "github.com/fujiwara/aws-sdk-client-go"
)

type ParamTestCase struct {
	Name         string
	Param        *sdkclient.ClientMethodParam
	Inject       map[string]any
	ExpectedJSON string
}

var ParamTestCases = []ParamTestCase{
	{
		Name: "inject nil",
		Param: &sdkclient.ClientMethodParam{
			InputBytes: []byte(`{"foo":"bar"}`),
		},
		Inject:       map[string]any{},
		ExpectedJSON: `{"foo":"bar"}`,
	},
	{
		Name: "inject int64",
		Param: &sdkclient.ClientMethodParam{
			InputBytes: []byte(`{"foo": "bar"}`),
		},
		Inject:       map[string]any{"int64": int64(42), "int64p": aws.Int64(42)},
		ExpectedJSON: `{"foo":"bar","int64":42,"int64p":42}`,
	},
	{
		Name: "inject string",
		Param: &sdkclient.ClientMethodParam{
			InputBytes: []byte(`{"foo": "bar"}`),
		},
		Inject:       map[string]any{"string": "baz", "stringp": aws.String("bazzz")},
		ExpectedJSON: `{"foo":"bar","string":"baz","stringp":"bazzz"}`,
	},
	{
		Name: "inject bool",
		Param: &sdkclient.ClientMethodParam{
			InputBytes: []byte(`{"foo": "bar"}`),
		},
		Inject:       map[string]any{"bool": true, "boolp": aws.Bool(false)},
		ExpectedJSON: `{"bool":true,"boolp":false,"foo":"bar"}`,
	},
}

func TestClientMethodParams(t *testing.T) {
	for _, c := range ParamTestCases {
		t.Run(c.Name, func(t *testing.T) {
			c.Param.MustInject(c.Inject)
			if string(c.Param.InputBytes) != c.ExpectedJSON {
				t.Errorf("unexpected InputBytes: got %s, want %s", c.Param.InputBytes, c.ExpectedJSON)
			}
		})
	}
}
