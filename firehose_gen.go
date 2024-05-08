package sdkclient

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/firehose"
)

func firehose_CreateDeliveryStream(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := firehose.NewFromConfig(awsCfg)
	var in firehose.CreateDeliveryStreamInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.CreateDeliveryStream(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call CreateDeliveryStream: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func firehose_DeleteDeliveryStream(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := firehose.NewFromConfig(awsCfg)
	var in firehose.DeleteDeliveryStreamInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.DeleteDeliveryStream(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call DeleteDeliveryStream: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func firehose_DescribeDeliveryStream(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := firehose.NewFromConfig(awsCfg)
	var in firehose.DescribeDeliveryStreamInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.DescribeDeliveryStream(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call DescribeDeliveryStream: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func firehose_ListDeliveryStreams(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := firehose.NewFromConfig(awsCfg)
	var in firehose.ListDeliveryStreamsInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.ListDeliveryStreams(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call ListDeliveryStreams: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func firehose_ListTagsForDeliveryStream(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := firehose.NewFromConfig(awsCfg)
	var in firehose.ListTagsForDeliveryStreamInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.ListTagsForDeliveryStream(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call ListTagsForDeliveryStream: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func firehose_PutRecord(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := firehose.NewFromConfig(awsCfg)
	var in firehose.PutRecordInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.PutRecord(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call PutRecord: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func firehose_PutRecordBatch(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := firehose.NewFromConfig(awsCfg)
	var in firehose.PutRecordBatchInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.PutRecordBatch(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call PutRecordBatch: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func firehose_StartDeliveryStreamEncryption(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := firehose.NewFromConfig(awsCfg)
	var in firehose.StartDeliveryStreamEncryptionInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.StartDeliveryStreamEncryption(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call StartDeliveryStreamEncryption: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func firehose_StopDeliveryStreamEncryption(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := firehose.NewFromConfig(awsCfg)
	var in firehose.StopDeliveryStreamEncryptionInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.StopDeliveryStreamEncryption(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call StopDeliveryStreamEncryption: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func firehose_TagDeliveryStream(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := firehose.NewFromConfig(awsCfg)
	var in firehose.TagDeliveryStreamInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.TagDeliveryStream(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call TagDeliveryStream: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func firehose_UntagDeliveryStream(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := firehose.NewFromConfig(awsCfg)
	var in firehose.UntagDeliveryStreamInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.UntagDeliveryStream(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call UntagDeliveryStream: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func firehose_UpdateDestination(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := firehose.NewFromConfig(awsCfg)
	var in firehose.UpdateDestinationInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.UpdateDestination(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call UpdateDestination: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func init() {
	clientMethods["firehose_CreateDeliveryStream"] = firehose_CreateDeliveryStream
	clientMethods["firehose_DeleteDeliveryStream"] = firehose_DeleteDeliveryStream
	clientMethods["firehose_DescribeDeliveryStream"] = firehose_DescribeDeliveryStream
	clientMethods["firehose_ListDeliveryStreams"] = firehose_ListDeliveryStreams
	clientMethods["firehose_ListTagsForDeliveryStream"] = firehose_ListTagsForDeliveryStream
	clientMethods["firehose_PutRecord"] = firehose_PutRecord
	clientMethods["firehose_PutRecordBatch"] = firehose_PutRecordBatch
	clientMethods["firehose_StartDeliveryStreamEncryption"] = firehose_StartDeliveryStreamEncryption
	clientMethods["firehose_StopDeliveryStreamEncryption"] = firehose_StopDeliveryStreamEncryption
	clientMethods["firehose_TagDeliveryStream"] = firehose_TagDeliveryStream
	clientMethods["firehose_UntagDeliveryStream"] = firehose_UntagDeliveryStream
	clientMethods["firehose_UpdateDestination"] = firehose_UpdateDestination
}
