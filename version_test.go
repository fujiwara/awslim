package sdkclient_test

import (
	"bytes"
	"context"
	"runtime/debug"
	"testing"

	sdkclient "github.com/fujiwara/awslim"
)

var testBuildInfo = &debug.BuildInfo{
	Deps: []*debug.Module{
		{Path: "github.com/alecthomas/kong", Version: "v1.16.1"},
		{Path: "github.com/aws/aws-sdk-go-v2", Version: "v1.45.1"},
		{Path: "github.com/aws/aws-sdk-go-v2/config", Version: "v1.33.1"},
		{Path: "github.com/aws/aws-sdk-go-v2/service/s3", Version: "v1.109.1"},
		{Path: "github.com/aws/aws-sdk-go-v2/service/sts", Version: "v1.0.0",
			Replace: &debug.Module{Path: "github.com/aws/aws-sdk-go-v2/service/sts", Version: "v1.47.1"}},
		{Path: "github.com/aws/smithy-go", Version: "v1.24.0"},
	},
}

func TestVersion(t *testing.T) {
	t.Setenv("XDG_CONFIG_HOME", "testdata")
	ctx := context.Background()

	cases := []struct {
		name   string
		args   []string
		info   *debug.BuildInfo
		expect string
		isErr  bool
	}{
		{
			name:   "version",
			args:   []string{"-v"},
			info:   testBuildInfo,
			expect: "awslim " + sdkclient.Version + "\naws-sdk-go-v2 v1.45.1\n",
		},
		{
			name:   "version with service",
			args:   []string{"-v", "s3"},
			info:   testBuildInfo,
			expect: "awslim " + sdkclient.Version + "\naws-sdk-go-v2 v1.45.1\naws-sdk-go-v2/service/s3 v1.109.1\n",
		},
		{
			name:   "version with replaced service",
			args:   []string{"-v", "sts"},
			info:   testBuildInfo,
			expect: "awslim " + sdkclient.Version + "\naws-sdk-go-v2 v1.45.1\naws-sdk-go-v2/service/sts v1.47.1\n",
		},
		{
			name:   "version with unknown service",
			args:   []string{"-v", "unknown"},
			info:   testBuildInfo,
			expect: "awslim " + sdkclient.Version + "\naws-sdk-go-v2 v1.45.1\n",
			isErr:  true,
		},
		{
			name:   "version without build info",
			args:   []string{"-v"},
			info:   nil,
			expect: "awslim " + sdkclient.Version + "\n",
		},
		{
			name: "sdk versions",
			args: []string{"--sdk-versions"},
			info: testBuildInfo,
			expect: "aws-sdk-go-v2 v1.45.1\n" +
				"aws-sdk-go-v2/config v1.33.1\n" +
				"aws-sdk-go-v2/service/s3 v1.109.1\n" +
				"aws-sdk-go-v2/service/sts v1.47.1\n",
		},
		{
			name:   "sdk versions without build info",
			args:   []string{"--sdk-versions"},
			info:   nil,
			expect: "",
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			sdkclient.SetBuildInfo(tc.info)
			t.Cleanup(func() { sdkclient.SetBuildInfo(nil) })
			buf := &bytes.Buffer{}
			c, err := newCLI(ctx, tc.args, buf)
			if err != nil {
				t.Fatal(err)
			}
			err = c.Dispatch(ctx)
			if tc.isErr && err == nil {
				t.Fatal("expected error, but got nil")
			} else if !tc.isErr && err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got := buf.String(); got != tc.expect {
				t.Errorf("unexpected output: got %q, expect %q", got, tc.expect)
			}
		})
	}
}
