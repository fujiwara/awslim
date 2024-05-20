package sdkclient

import (
	"encoding/json"
	"io"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
)

type clientMethodParam struct {
	awsCfg      aws.Config
	b           json.RawMessage
	inputReader io.Reader
	cleanup     []func() error
}

func (p *clientMethodParam) Cleanup() {
	for _, f := range p.cleanup {
		if err := f(); err != nil {
			log.Printf("[warn] failed to cleanup: %v", err)
		}
	}
}
