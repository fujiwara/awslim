package sdkclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/itchyny/gojq"
)

type clientMethodParam struct {
	InputBytes        json.RawMessage
	InputReader       io.Reader
	InputReaderLength *int64
	OutputWriter      io.Writer
	DryRun            bool
	Strict            bool

	awsCfg  aws.Config
	cleanup []func() error
}

func (p *clientMethodParam) Cleanup() {
	for _, f := range p.cleanup {
		if err := f(); err != nil {
			log.Printf("[warn] failed to cleanup: %v", err)
		}
	}
}

func UnmarshalJSON[T any](b []byte, v T, strict bool) error {
	dec := json.NewDecoder(bytes.NewReader(b))
	if strict {
		dec.DisallowUnknownFields()
	}
	if err := dec.Decode(v); err != nil {
		return fmt.Errorf("failed to unmarshal to %T: %w", v, err)
	}
	return nil
}

func (p *clientMethodParam) Output(src io.ReadCloser) error {
	if p.OutputWriter == nil {
		return nil
	}
	defer src.Close()
	_, err := io.Copy(p.OutputWriter, src)
	return err
}

func (p *clientMethodParam) Validate(name, inputReaderField, outputReadCloserField string) error {
	if p.InputReader != nil && inputReaderField == "" {
		return fmt.Errorf("%sInput has not io.Reader field", name)
	}
	if p.OutputWriter != nil && outputReadCloserField == "" {
		return fmt.Errorf("%sOutput has not io.ReadCloser field", name)
	}
	return nil
}

func (p *clientMethodParam) MustInject(in map[string]any) {
	v := make(map[string]any)
	if err := json.Unmarshal(p.InputBytes, &v); err != nil {
		panic(fmt.Sprintf("failed to marshal %s:", err))
	}
	var qs []string
	for field, value := range in {
		if value == nil {
			qs = append(qs, fmt.Sprintf("del(.%s)", field))
		} else {
			jv, err := json.Marshal(value)
			if err != nil {
				panic(fmt.Sprintf("failed to marshal %s:", err))
			}
			qs = append(qs, fmt.Sprintf(".%s = %s", field, string(jv)))
		}
	}
	q := strings.Join(qs, " | ")
	query, err := gojq.Parse(q)
	if err != nil {
		panic(fmt.Sprintf("failed to parse query %s: %v", q, err))
	}
	iter := query.Run(v)
	for {
		if v, ok := iter.Next(); ok {
			if !ok {
				break
			}
			if err, ok := v.(error); ok {
				if err, ok := err.(*gojq.HaltError); ok && err.Value() == nil {
					break
				}
				panic(err)
			}
			p.InputBytes, _ = json.Marshal(v)
			return
		}
	}
}
