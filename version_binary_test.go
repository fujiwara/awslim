package sdkclient_test

import (
	"os"
	"os/exec"
	"path/filepath"
	"slices"
	"strings"
	"testing"

	sdkclient "github.com/fujiwara/awslim"
)

// TestVersionBinary builds the real awslim binary and verifies that the SDK
// versions embedded in its build info are shown. Test binaries built by
// `go test` do not embed module dependency info on Go 1.26 and earlier, so
// this cannot be verified in-process.
func TestVersionBinary(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping in short mode")
	}
	goBin, err := exec.LookPath("go")
	if err != nil {
		t.Skip("go command not found")
	}
	bin := filepath.Join(t.TempDir(), "awslim")
	build := exec.Command(goBin, "build", "-o", bin, "./cmd/awslim")
	if out, err := build.CombinedOutput(); err != nil {
		t.Fatalf("failed to build: %v\n%s", err, out)
	}
	run := func(args ...string) (string, error) {
		cmd := exec.Command(bin, args...)
		cmd.Env = append(os.Environ(), "XDG_CONFIG_HOME=testdata")
		out, err := cmd.Output()
		return string(out), err
	}
	// expected versions from go.mod
	modVersion := func(path string) string {
		out, err := exec.Command(goBin, "list", "-m", "-f", "{{.Version}}", path).Output()
		if err != nil {
			t.Fatalf("failed to get version of %s: %v", path, err)
		}
		return strings.TrimSpace(string(out))
	}
	coreVersion := modVersion("github.com/aws/aws-sdk-go-v2")
	configVersion := modVersion("github.com/aws/aws-sdk-go-v2/config")

	t.Run("version", func(t *testing.T) {
		out, err := run("-v")
		if err != nil {
			t.Fatal(err)
		}
		expect := "awslim " + sdkclient.Version + "\naws-sdk-go-v2 " + coreVersion + "\n"
		if out != expect {
			t.Errorf("unexpected output: got %q, expect %q", out, expect)
		}
	})

	t.Run("sdk versions", func(t *testing.T) {
		out, err := run("--sdk-versions")
		if err != nil {
			t.Fatal(err)
		}
		lines := strings.Split(strings.TrimSpace(out), "\n")
		got := map[string]string{}
		for _, line := range lines {
			name, version, ok := strings.Cut(line, " ")
			if !ok {
				t.Fatalf("unexpected line: %q", line)
			}
			got[name] = version
		}
		if got["aws-sdk-go-v2"] != coreVersion {
			t.Errorf("unexpected core version: got %q, expect %q", got["aws-sdk-go-v2"], coreVersion)
		}
		if got["aws-sdk-go-v2/config"] != configVersion {
			t.Errorf("unexpected config version: got %q, expect %q", got["aws-sdk-go-v2/config"], configVersion)
		}
		// every module listed must match go.mod
		for name, version := range got {
			path := "github.com/aws/" + name
			if v := modVersion(path); v != version {
				t.Errorf("unexpected version of %s: got %q, expect %q", path, version, v)
			}
		}
		if !slices.IsSorted(lines) {
			t.Errorf("output is not sorted: %q", out)
		}
	})
}
