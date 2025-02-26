package sdkclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/jmespath/go-jmespath"
)

type clientMethodParam struct {
	InputBytes        json.RawMessage
	InputReader       io.Reader
	InputReaderLength *int64
	OutputWriter      io.Writer
	DryRun            bool
	Strict            bool
	NextToken         struct {
		OutputField string
		InputField  string
	}
	awsCfg  aws.Config
	cleanup []func() error
}

func (p *clientMethodParam) Cleanup() {
	for _, f := range p.cleanup {
		if err := f(); err != nil {
			slog.Warn("failed to cleanup", "error", err.Error())
		}
	}
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

func (p *clientMethodParam) InjectMap(in map[string]any) error {
	v := make(map[string]any)
	if err := json.Unmarshal(p.InputBytes, &v); err != nil {
		return fmt.Errorf("failed to unmarshal %s: %w", p.InputBytes, err)
	}
	for field, value := range in {
		v[field] = value
	}
	if b, err := json.Marshal(v); err != nil {
		return fmt.Errorf("failed to marshal %v: %w", v, err)
	} else {
		p.InputBytes = b
	}
	return nil
}

func (p *clientMethodParam) Inject(field string, value any) error {
	v := make(map[string]any)
	if err := json.Unmarshal(p.InputBytes, &v); err != nil {
		return fmt.Errorf("failed to unmarshal %s: %w", p.InputBytes, err)
	}
	v[field] = value
	if b, err := json.Marshal(v); err != nil {
		return fmt.Errorf("failed to marshal %v: %w", v, err)
	} else {
		p.InputBytes = b
	}
	return nil
}

func (p *clientMethodParam) MustInject(m map[string]any) {
	for k, v := range m {
		if err := p.Inject(k, v); err != nil {
			panic(err)
		}
	}
}

func (p *clientMethodParam) SetNextToken(s string) {
	nt := strings.SplitN(s, "=", 2)
	switch len(nt) {
	case 0:
		return
	case 1:
		p.NextToken.OutputField = nt[0]
		p.NextToken.InputField = nt[0]
	case 2:
		p.NextToken.OutputField = nt[0]
		p.NextToken.InputField = nt[1]
	}
}

func (p *clientMethodParam) FollowNext(out any) (bool, error) {
	if p.NextToken.OutputField == "" {
		return false, nil
	}

	token, err := jmespath.Search(p.NextToken.OutputField, out)
	if err != nil {
		return false, fmt.Errorf("failed to extract %s: %w", p.NextToken.OutputField, err)
	}
	var tokenValue string
	switch t := token.(type) {
	case string:
		tokenValue = t
	case *string:
		tokenValue = aws.ToString(t)
	default:
	}
	if tokenValue == "" {
		return false, nil
	}
	if err := p.Inject(p.NextToken.InputField, tokenValue); err != nil {
		return false, fmt.Errorf("failed to inject %s=%v: %w", p.NextToken.InputField, tokenValue, err)
	}
	return true, nil
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
