package sdkclient

import (
	"context"
	"fmt"
	"maps"
	"runtime/debug"
	"slices"
	"strings"
)

// Version is the version of awslim. Updated by tagpr, and overridden by -ldflags at build time.
var Version = "v0.7.0"

// readBuildInfo is replaced in tests.
var readBuildInfo = debug.ReadBuildInfo

const (
	sdkModulePrefix        = "github.com/aws/aws-sdk-go-v2"
	sdkServiceModulePrefix = sdkModulePrefix + "/service/"
)

// sdkModuleVersions returns versions of aws-sdk-go-v2 modules embedded in the binary.
// The keys are module paths without the "github.com/aws/aws-sdk-go-v2" prefix
// (e.g. "" for the core module, "/config", "/service/s3").
func sdkModuleVersions() map[string]string {
	versions := map[string]string{}
	info, ok := readBuildInfo()
	if !ok {
		return versions
	}
	for _, dep := range info.Deps {
		if dep.Replace != nil {
			dep = dep.Replace
		}
		if dep.Path == sdkModulePrefix || strings.HasPrefix(dep.Path, sdkModulePrefix+"/") {
			versions[strings.TrimPrefix(dep.Path, sdkModulePrefix)] = dep.Version
		}
	}
	return versions
}

// ShowVersion prints the version of awslim and the core AWS SDK.
// If a service is specified, the version of the service module is also printed.
func (c *CLI) ShowVersion(_ context.Context) error {
	fmt.Fprintf(c.w, "awslim %s\n", Version)
	versions := sdkModuleVersions()
	if v, ok := versions[""]; ok {
		fmt.Fprintf(c.w, "aws-sdk-go-v2 %s\n", v)
	}
	if c.Service != "" {
		v, ok := versions["/service/"+c.Service]
		if !ok {
			return fmt.Errorf("unknown service: %s", c.Service)
		}
		fmt.Fprintf(c.w, "aws-sdk-go-v2/service/%s %s\n", c.Service, v)
	}
	return nil
}

// ShowSDKVersions prints versions of all aws-sdk-go-v2 modules embedded in the binary.
func (c *CLI) ShowSDKVersions(_ context.Context) error {
	versions := sdkModuleVersions()
	for _, path := range slices.Sorted(maps.Keys(versions)) {
		fmt.Fprintf(c.w, "aws-sdk-go-v2%s %s\n", path, versions[path])
	}
	return nil
}
