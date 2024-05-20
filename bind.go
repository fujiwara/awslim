package sdkclient

import (
	"encoding/json"
	"io"

	"github.com/aws/aws-sdk-go-v2/aws"
)

type clientMethodParam struct {
	awsCfg     aws.Config
	b          json.RawMessage
	bindKey    string
	bindReader io.Reader
}
