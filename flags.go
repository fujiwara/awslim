package sdkclient

import (
	"fmt"
	"strings"

	"github.com/alecthomas/kong"
)

func flagExists(k *kong.Kong, flagName string) bool {
	for _, flagSet := range k.Model.AllFlags(false) {
		for _, f := range flagSet {
			if f.Name == flagName {
				return true
			}
		}
	}
	return false
}

func parseDynamicFlags(k *kong.Kong, args []string, convert func(string) string) ([]string, map[string]any, error) {
	if convert == nil {
		convert = func(s string) string { return s }
	}
	dynamicFlags := make(map[string]any)
	rArgs := make([]string, 0, len(args))
	for i := 0; i < len(args); i++ {
		arg := args[i]
		if arg == "--" {
			break
		}
		if !strings.HasPrefix(arg, "--") {
			rArgs = append(rArgs, arg)
			continue
		}
		parts := strings.SplitN(arg, "=", 2)
		name := strings.TrimPrefix(parts[0], "--")
		if flagExists(k, name) {
			rArgs = append(rArgs, arg)
			continue
		}
		if len(parts) == 2 {
			// --flag=value
			dynamicFlags[convert(name)] = parts[1]
		} else {
			// --flag value
			// read next arg as value
			if i+1 < len(args) {
				dynamicFlags[convert(name)] = args[i+1]
				i++
			} else {
				return nil, nil, fmt.Errorf("flag %s requires a value", arg)
			}
		}
	}
	return rArgs, dynamicFlags, nil
}
