package sdkclient_test

import (
	"bytes"
	"context"
	"regexp"
	"strings"
	"testing"
)

func TestVersion(t *testing.T) {
	t.Setenv("XDG_CONFIG_HOME", "testdata")
	ctx := context.Background()

	t.Run("version", func(t *testing.T) {
		buf := &bytes.Buffer{}
		c, err := newCLI(ctx, []string{"-v"}, buf)
		if err != nil {
			t.Fatal(err)
		}
		if err := c.Dispatch(ctx); err != nil {
			t.Fatal(err)
		}
		lines := strings.Split(strings.TrimSpace(buf.String()), "\n")
		if len(lines) != 2 {
			t.Fatalf("unexpected output: %q", buf.String())
		}
		if lines[0] != "awslim HEAD" {
			t.Errorf("unexpected awslim version line: %q", lines[0])
		}
		if !regexp.MustCompile(`^aws-sdk-go-v2 v\d+\.\d+\.\d+`).MatchString(lines[1]) {
			t.Errorf("unexpected sdk version line: %q", lines[1])
		}
	})

	t.Run("version with unknown service", func(t *testing.T) {
		buf := &bytes.Buffer{}
		c, err := newCLI(ctx, []string{"-v", "unknown"}, buf)
		if err != nil {
			t.Fatal(err)
		}
		if err := c.Dispatch(ctx); err == nil {
			t.Fatal("expected error, but got nil")
		}
	})

	t.Run("sdk versions", func(t *testing.T) {
		buf := &bytes.Buffer{}
		c, err := newCLI(ctx, []string{"--sdk-versions"}, buf)
		if err != nil {
			t.Fatal(err)
		}
		if err := c.Dispatch(ctx); err != nil {
			t.Fatal(err)
		}
		re := regexp.MustCompile(`^aws-sdk-go-v2(/\S+)? v\d+\.\d+\.\d+`)
		lines := strings.Split(strings.TrimSpace(buf.String()), "\n")
		if len(lines) < 2 { // at least core and config modules
			t.Fatalf("unexpected output: %q", buf.String())
		}
		for _, line := range lines {
			if !re.MatchString(line) {
				t.Errorf("unexpected line: %q", line)
			}
		}
		if !strings.HasPrefix(lines[0], "aws-sdk-go-v2 v") {
			t.Errorf("core module must come first: %q", lines[0])
		}
	})
}
