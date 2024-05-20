package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"
	"text/template"
)

//go:generate go run ../aws-sdk-client-gen-gen/main.go
//go:generate go get ./...

const templateStr = `// Code generated by cmd/aws-sdk-client-gen/main.go; DO NOT EDIT.
package sdkclient

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/{{ .PkgName }}"
)

{{ range .Methods }}
func {{ $.PkgName }}_{{ .Name }}(ctx context.Context, p *clientMethodParam) (any, error) {
	svc := {{ $.PkgName }}.NewFromConfig(p.awsCfg)
	var in {{ .Input }}
	{{- if .InputReaderLengthField }}
	p.mustInject("{{ .InputReaderLengthField }}", p.InputReaderLength)
	{{- end }}
	if err := json.Unmarshal(p.InputBytes, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if err := p.validate("{{ $.PkgName }}.{{ .Name }}", "{{ .InputReaderField }}", "{{ .OutputReadCloserField }}"); err != nil {
		return nil, err
	}
	{{- if .InputReaderField }}
	if p.InputReader != nil {
		in.{{ .InputReaderField }} = p.InputReader
	}
	{{- end }}
	{{- if .OutputReadCloserField }}
	if out, err := svc.{{ .Name }}(ctx, &in); err != nil {
		return nil, err
	} else {
		if err := p.Output(out.{{ .OutputReadCloserField }}); err != nil {
			return nil, err
		}
		return out, nil
	}
	{{- else }}
	return svc.{{ .Name }}(ctx, &in)
	{{- end }}
}

{{ end }}

func init() {
{{- range .Methods }}
	clientMethods["{{ $.PkgName }}#Client.{{ .Name }}"] = {{ $.PkgName }}_{{ .Name }}
{{- end }}
}
`

func main() {
	generateAll()
}

func gen(pkgName string, clientType reflect.Type, genNames []string) error {
	log.Printf("generating %s_gen.go", pkgName)

	methods := make([]map[string]string, 0)
	for i := 0; i < clientType.NumMethod(); i++ {
		method := clientType.Method(i)
		if len(genNames) > 0 && !contains(genNames, method.Name) {
			continue
		}
		params := make([]string, 0)
		for j := 0; j < method.Type.NumIn(); j++ {
			params = append(params, method.Type.In(j).String())
		}
		if len(params) <= 1 {
			log.Printf("no params func %s", method.Name)
			continue
		}
		inputParam := method.Type.In(2)
		var inputReaderField, inputReaderLengthField string
		for j := 0; j < inputParam.Elem().NumField(); j++ {
			field := inputParam.Elem().Field(j)
			if t := field.Type.String(); t == "io.Reader" {
				log.Printf("found %s field in %s.%sInput %s %s", t, pkgName, method.Name, field.Name, t)
				if inputReaderField != "" {
					return fmt.Errorf("found multiple io.Reader fields in %s.%sInput", pkgName, method.Name)
				}
				inputReaderField = field.Name
			}
		}
		if inputReaderField != "" {
			for j := 0; j < inputParam.Elem().NumField(); j++ {
				field := inputParam.Elem().Field(j)
				if t := field.Name; strings.Contains(t, "Length") {
					log.Printf("found %s field in %s.%sInput %s %s", t, pkgName, method.Name, field.Name, t)
					if inputReaderLengthField != "" {
						return fmt.Errorf("found multiple Length fields in %s.%sInput", pkgName, method.Name)
					}
					inputReaderLengthField = field.Name
				}
			}
		}

		var outputReadCloserField string
		output := method.Type.Out(0).Elem()
		for j := 0; j < output.NumField(); j++ {
			field := output.Field(j)
			if t := field.Type.String(); t == "io.ReadCloser" {
				log.Printf("found %s field in %s.%sOutput %s %s", t, pkgName, method.Name, field.Name, t)
				if outputReadCloserField != "" {
					return fmt.Errorf("found multiple io.ReadCloser fields in %s.%sOutput", pkgName, method.Name)
				}
				outputReadCloserField = field.Name
			}
		}
		methods = append(methods, map[string]string{
			"Name":                   method.Name,
			"Input":                  strings.TrimPrefix(params[2], "*"),
			"InputReaderField":       inputReaderField,
			"InputReaderLengthField": inputReaderLengthField,
			"OutputReadCloserField":  outputReadCloserField,
		})
	}

	tmpl, err := template.New("clientGen").Parse(templateStr)
	if err != nil {
		return fmt.Errorf("failed to parse template: %w", err)
	}

	data := map[string]interface{}{
		"PkgName": pkgName,
		"Methods": methods,
	}

	buf := &bytes.Buffer{}
	if err := tmpl.Execute(buf, data); err != nil {
		return fmt.Errorf("failed to execute template: %w", err)
	}
	if err := os.WriteFile(pkgName+"_gen.go", buf.Bytes(), 0644); err != nil {
		return err
	}
	log.Printf("generated %s_gen.go", pkgName)
	return nil
}

func contains(ss []string, s string) bool {
	for _, v := range ss {
		if v == s {
			return true
		}
	}
	return false
}
