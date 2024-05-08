package main

import (
	"fmt"
	"log"
	"os"

	"github.com/goccy/go-yaml"
)

type GenerateConfig struct {
	Services map[string][]string `json:"services"`
}

func main() {
	log.Println("generating gen.go")
	o, err := os.Create("gen.go")
	if err != nil {
		log.Fatalf("failed to create gen.go: %v", err)
	}
	f, err := os.Open("../../gen.yaml")
	if err != nil {
		log.Fatalf("failed to open gen.yaml: %v", err)
	}
	defer f.Close()
	cfg := GenerateConfig{}
	if err := yaml.NewDecoder(f).Decode(&cfg); err != nil {
		log.Fatalf("failed to decode gen.yaml: %v", err)
	}

	fmt.Fprintln(o, "// Code generated by cmd/aws-sdk-client-gen-gen/main.go; DO NOT EDIT.")
	fmt.Fprintln(o, "package main")
	fmt.Fprintln(o, "import \"reflect\"")
	for k := range cfg.Services {
		fmt.Fprintf(o, "import \"github.com/aws/aws-sdk-go-v2/service/%s\"\n", k)
	}
	fmt.Fprintln(o, "func generateAll() {")
	for k, v := range cfg.Services {
		if len(v) == 0 {
			fmt.Fprintf(o, "	gen(\"%s\", reflect.TypeOf(%s.New(%s.Options{})), nil)\n", k, k, k)
		} else {
			fmt.Fprintf(o, "	gen(\"%s\", reflect.TypeOf(%s.New(%s.Options{})), %#v)\n", k, k, k, v)
		}
	}
	fmt.Fprintln(o, "}")
	log.Println("generated gen.go")
}
