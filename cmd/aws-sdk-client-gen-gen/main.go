package main

import (
	"bytes"
	"log"
	"os"
	"text/template"

	"github.com/goccy/go-yaml"
)

const templateStr = `// Code generated by cmd/aws-sdk-client-gen-gen/main.go; DO NOT EDIT.
package main

import (
	"reflect"
{{ range $key, $value := .Services }}
	"github.com/aws/aws-sdk-go-v2/service/{{ $key }}"
{{- end }}
)

func generateAll() {
	var err error
{{ range $key, $value := .Services }}
{{- if eq (len $value) 0 }}
	err = gen("{{ $key }}", reflect.TypeOf({{ $key }}.New({{ $key }}.Options{})), nil)
{{- else }}
	err = gen("{{ $key }}", reflect.TypeOf({{ $key }}.New({{ $key }}.Options{})), {{ printf "%#v" $value }})
{{- end }}
	if err != nil {
		panic("failed to generate {{ $key }}" + err.Error())
	}
{{ end }}
}
`

type GenerateConfig struct {
	Services map[string][]string `json:"services"`
}

func main() {
	log.Println("generating gen.go")

	f, err := os.Open("../../gen.yaml")
	if err != nil {
		log.Fatalf("failed to open gen.yaml: %v", err)
	}
	defer f.Close()

	cfg := GenerateConfig{}
	if err := yaml.NewDecoder(f).Decode(&cfg); err != nil {
		log.Fatalf("failed to decode gen.yaml: %v", err)
	}
	tmpl, err := template.New("gen").Parse(templateStr)
	if err != nil {
		log.Fatalf("failed to parse template: %v", err)
	}
	buf := &bytes.Buffer{}
	if err := tmpl.Execute(buf, cfg); err != nil {
		log.Fatalf("failed to execute template: %v", err)
	}

	if err := os.WriteFile("gen.go", buf.Bytes(), 0644); err != nil {
		log.Fatalf("failed to write gen.go: %v", err)
	}
	log.Println("generated gen.go")
}
