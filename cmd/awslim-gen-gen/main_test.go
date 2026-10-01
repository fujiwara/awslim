package main

import (
	"slices"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestResolveConfigFromYAML(t *testing.T) {
	src := `
services:
  sts:
  s3@v1.100.0:
    - GetObject
    - PutObject
  ecs@latest:
`
	cfg, err := configFromYAML(strings.NewReader(src))
	if err != nil {
		t.Fatal(err)
	}
	sc, err := resolveConfig(cfg)
	if err != nil {
		t.Fatal(err)
	}
	expectServices := map[string][]string{
		"sts": nil,
		"s3":  {"GetObject", "PutObject"},
		"ecs": nil,
	}
	if diff := cmp.Diff(expectServices, sc.Services); diff != "" {
		t.Errorf("unexpected services (-want +got):\n%s", diff)
	}
	expectVersions := map[string]string{
		"s3":  "v1.100.0",
		"ecs": "latest",
	}
	if diff := cmp.Diff(expectVersions, sc.Versions); diff != "" {
		t.Errorf("unexpected versions (-want +got):\n%s", diff)
	}
}

func TestResolveConfigFromEnv(t *testing.T) {
	sc, err := resolveConfig(configFromEnv("ecs,s3@v1.100.0,,sts"))
	if err != nil {
		t.Fatal(err)
	}
	expectServices := map[string][]string{"ecs": nil, "s3": nil, "sts": nil}
	if diff := cmp.Diff(expectServices, sc.Services); diff != "" {
		t.Errorf("unexpected services (-want +got):\n%s", diff)
	}
	expectVersions := map[string]string{"s3": "v1.100.0"}
	if diff := cmp.Diff(expectVersions, sc.Versions); diff != "" {
		t.Errorf("unexpected versions (-want +got):\n%s", diff)
	}
}

func TestResolveConfigError(t *testing.T) {
	for _, env := range []string{
		"s3,s3@v1.100.0",
		"s3@",
		"@v1.100.0",
	} {
		if _, err := resolveConfig(configFromEnv(env)); err == nil {
			t.Errorf("expected error for %q", env)
		}
	}
}

func TestGoGetArgs(t *testing.T) {
	if args := goGetArgs(nil); args != nil {
		t.Errorf("expected nil, got %v", args)
	}
	args := goGetArgs(map[string]string{"s3": "v1.100.0", "ecs": "latest"})
	slices.Sort(args[1:])
	expect := []string{
		"get",
		"github.com/aws/aws-sdk-go-v2/service/ecs@latest",
		"github.com/aws/aws-sdk-go-v2/service/s3@v1.100.0",
	}
	if diff := cmp.Diff(expect, args); diff != "" {
		t.Errorf("unexpected args (-want +got):\n%s", diff)
	}
}
