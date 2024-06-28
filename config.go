package sdkclient

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/goccy/go-yaml"
	"github.com/google/go-jsonnet"
	"github.com/mattn/go-shellwords"
)

const (
	configSubdir   = "awslim"
	configBasename = "config"
)

type RuntimeConfig struct {
	Open    string           `json:"open" yaml:"open"`
	Aliases map[string]Alias `json:"aliases" yaml:"aliases"`
}

type Alias string

func (a Alias) Parse() ([]string, error) {
	return shellwords.Parse(string(a))
}

func RuntimeConfigDir() string {
	if h := os.Getenv("XDG_CONFIG_HOME"); h != "" {
		return filepath.Join(h, configSubdir)
	} else {
		d, err := os.UserHomeDir()
		if err != nil {
			d = os.Getenv("HOME")
		}
		return filepath.Join(d, ".config", configSubdir)
	}
}

func loadRuntimeConfig() (*RuntimeConfig, error) {
	dir := RuntimeConfigDir()
	if _, err := os.Stat(dir); err != nil {
		if os.IsNotExist(err) {
			return &RuntimeConfig{}, nil
		}
		return nil, err
	}
	var out RuntimeConfig
	var found bool
	for _, ext := range []string{"jsonnet", "yaml", "yml", "json"} {
		f := filepath.Join(dir, configBasename+"."+ext)
		if _, err := os.Stat(f); err != nil {
			if os.IsNotExist(err) {
				continue
			}
			return nil, err
		}
		slog.Debug(fmt.Sprintf("loading config file: %s", f))
		switch ext {
		case "jsonnet", "json":
			vm := jsonnet.MakeVM()
			b, err := vm.EvaluateFile(f)
			if err != nil {
				return nil, err
			}
			if found {
				slog.Warn(fmt.Sprintf("found multiple config files, ignoring %s", f))
			} else {
				if err := json.Unmarshal([]byte(b), &out); err != nil {
					return nil, err
				}
				found = true
			}
		case "yaml", "yml":
			b, err := os.ReadFile(f)
			if err != nil {
				return nil, err
			}
			if found {
				slog.Warn(fmt.Sprintf("found multiple config files, ignoring %s", f))
			} else {
				if err := yaml.Unmarshal(b, &out); err != nil {
					return nil, err
				}
				found = true
			}
		}
	}
	return &out, nil
}

func (c *CLI) resolveAliases(args []string) ([]string, error) {
	if len(args) == 0 {
		return args, nil
	}
	first := args[0]
	rest := args[1:]
	if alias, ok := c.rc.Aliases[first]; ok {
		slog.Debug(fmt.Sprintf("alias found %s -> %s", first, alias))
		parsed, err := alias.Parse()
		if err != nil {
			return nil, err
		}
		switch len(args) {
		case 0:
			return nil, fmt.Errorf("alias %s is empty", c.Service)
		}
		return append(parsed, rest...), nil
	}
	return args, nil
}
