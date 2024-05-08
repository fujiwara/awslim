package sdkclient

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/kinesis"
)

func kinesis_AddTagsToStream(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := kinesis.NewFromConfig(awsCfg)
	var in kinesis.AddTagsToStreamInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.AddTagsToStream(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call AddTagsToStream: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func kinesis_CreateStream(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := kinesis.NewFromConfig(awsCfg)
	var in kinesis.CreateStreamInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.CreateStream(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call CreateStream: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func kinesis_DecreaseStreamRetentionPeriod(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := kinesis.NewFromConfig(awsCfg)
	var in kinesis.DecreaseStreamRetentionPeriodInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.DecreaseStreamRetentionPeriod(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call DecreaseStreamRetentionPeriod: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func kinesis_DeleteResourcePolicy(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := kinesis.NewFromConfig(awsCfg)
	var in kinesis.DeleteResourcePolicyInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.DeleteResourcePolicy(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call DeleteResourcePolicy: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func kinesis_DeleteStream(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := kinesis.NewFromConfig(awsCfg)
	var in kinesis.DeleteStreamInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.DeleteStream(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call DeleteStream: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func kinesis_DeregisterStreamConsumer(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := kinesis.NewFromConfig(awsCfg)
	var in kinesis.DeregisterStreamConsumerInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.DeregisterStreamConsumer(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call DeregisterStreamConsumer: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func kinesis_DescribeLimits(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := kinesis.NewFromConfig(awsCfg)
	var in kinesis.DescribeLimitsInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.DescribeLimits(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call DescribeLimits: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func kinesis_DescribeStream(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := kinesis.NewFromConfig(awsCfg)
	var in kinesis.DescribeStreamInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.DescribeStream(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call DescribeStream: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func kinesis_DescribeStreamConsumer(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := kinesis.NewFromConfig(awsCfg)
	var in kinesis.DescribeStreamConsumerInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.DescribeStreamConsumer(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call DescribeStreamConsumer: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func kinesis_DescribeStreamSummary(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := kinesis.NewFromConfig(awsCfg)
	var in kinesis.DescribeStreamSummaryInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.DescribeStreamSummary(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call DescribeStreamSummary: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func kinesis_DisableEnhancedMonitoring(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := kinesis.NewFromConfig(awsCfg)
	var in kinesis.DisableEnhancedMonitoringInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.DisableEnhancedMonitoring(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call DisableEnhancedMonitoring: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func kinesis_EnableEnhancedMonitoring(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := kinesis.NewFromConfig(awsCfg)
	var in kinesis.EnableEnhancedMonitoringInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.EnableEnhancedMonitoring(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call EnableEnhancedMonitoring: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func kinesis_GetRecords(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := kinesis.NewFromConfig(awsCfg)
	var in kinesis.GetRecordsInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.GetRecords(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call GetRecords: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func kinesis_GetResourcePolicy(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := kinesis.NewFromConfig(awsCfg)
	var in kinesis.GetResourcePolicyInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.GetResourcePolicy(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call GetResourcePolicy: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func kinesis_GetShardIterator(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := kinesis.NewFromConfig(awsCfg)
	var in kinesis.GetShardIteratorInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.GetShardIterator(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call GetShardIterator: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func kinesis_IncreaseStreamRetentionPeriod(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := kinesis.NewFromConfig(awsCfg)
	var in kinesis.IncreaseStreamRetentionPeriodInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.IncreaseStreamRetentionPeriod(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call IncreaseStreamRetentionPeriod: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func kinesis_ListShards(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := kinesis.NewFromConfig(awsCfg)
	var in kinesis.ListShardsInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.ListShards(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call ListShards: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func kinesis_ListStreamConsumers(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := kinesis.NewFromConfig(awsCfg)
	var in kinesis.ListStreamConsumersInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.ListStreamConsumers(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call ListStreamConsumers: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func kinesis_ListStreams(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := kinesis.NewFromConfig(awsCfg)
	var in kinesis.ListStreamsInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.ListStreams(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call ListStreams: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func kinesis_ListTagsForStream(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := kinesis.NewFromConfig(awsCfg)
	var in kinesis.ListTagsForStreamInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.ListTagsForStream(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call ListTagsForStream: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func kinesis_MergeShards(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := kinesis.NewFromConfig(awsCfg)
	var in kinesis.MergeShardsInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.MergeShards(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call MergeShards: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func kinesis_PutRecord(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := kinesis.NewFromConfig(awsCfg)
	var in kinesis.PutRecordInput
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

func kinesis_PutRecords(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := kinesis.NewFromConfig(awsCfg)
	var in kinesis.PutRecordsInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.PutRecords(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call PutRecords: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func kinesis_PutResourcePolicy(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := kinesis.NewFromConfig(awsCfg)
	var in kinesis.PutResourcePolicyInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.PutResourcePolicy(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call PutResourcePolicy: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func kinesis_RegisterStreamConsumer(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := kinesis.NewFromConfig(awsCfg)
	var in kinesis.RegisterStreamConsumerInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.RegisterStreamConsumer(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call RegisterStreamConsumer: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func kinesis_RemoveTagsFromStream(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := kinesis.NewFromConfig(awsCfg)
	var in kinesis.RemoveTagsFromStreamInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.RemoveTagsFromStream(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call RemoveTagsFromStream: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func kinesis_SplitShard(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := kinesis.NewFromConfig(awsCfg)
	var in kinesis.SplitShardInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.SplitShard(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call SplitShard: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func kinesis_StartStreamEncryption(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := kinesis.NewFromConfig(awsCfg)
	var in kinesis.StartStreamEncryptionInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.StartStreamEncryption(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call StartStreamEncryption: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func kinesis_StopStreamEncryption(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := kinesis.NewFromConfig(awsCfg)
	var in kinesis.StopStreamEncryptionInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.StopStreamEncryption(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call StopStreamEncryption: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func kinesis_SubscribeToShard(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := kinesis.NewFromConfig(awsCfg)
	var in kinesis.SubscribeToShardInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.SubscribeToShard(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call SubscribeToShard: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func kinesis_UpdateShardCount(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := kinesis.NewFromConfig(awsCfg)
	var in kinesis.UpdateShardCountInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.UpdateShardCount(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call UpdateShardCount: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func kinesis_UpdateStreamMode(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := kinesis.NewFromConfig(awsCfg)
	var in kinesis.UpdateStreamModeInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.UpdateStreamMode(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call UpdateStreamMode: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func init() {
	clientMethods["kinesis_AddTagsToStream"] = kinesis_AddTagsToStream
	clientMethods["kinesis_CreateStream"] = kinesis_CreateStream
	clientMethods["kinesis_DecreaseStreamRetentionPeriod"] = kinesis_DecreaseStreamRetentionPeriod
	clientMethods["kinesis_DeleteResourcePolicy"] = kinesis_DeleteResourcePolicy
	clientMethods["kinesis_DeleteStream"] = kinesis_DeleteStream
	clientMethods["kinesis_DeregisterStreamConsumer"] = kinesis_DeregisterStreamConsumer
	clientMethods["kinesis_DescribeLimits"] = kinesis_DescribeLimits
	clientMethods["kinesis_DescribeStream"] = kinesis_DescribeStream
	clientMethods["kinesis_DescribeStreamConsumer"] = kinesis_DescribeStreamConsumer
	clientMethods["kinesis_DescribeStreamSummary"] = kinesis_DescribeStreamSummary
	clientMethods["kinesis_DisableEnhancedMonitoring"] = kinesis_DisableEnhancedMonitoring
	clientMethods["kinesis_EnableEnhancedMonitoring"] = kinesis_EnableEnhancedMonitoring
	clientMethods["kinesis_GetRecords"] = kinesis_GetRecords
	clientMethods["kinesis_GetResourcePolicy"] = kinesis_GetResourcePolicy
	clientMethods["kinesis_GetShardIterator"] = kinesis_GetShardIterator
	clientMethods["kinesis_IncreaseStreamRetentionPeriod"] = kinesis_IncreaseStreamRetentionPeriod
	clientMethods["kinesis_ListShards"] = kinesis_ListShards
	clientMethods["kinesis_ListStreamConsumers"] = kinesis_ListStreamConsumers
	clientMethods["kinesis_ListStreams"] = kinesis_ListStreams
	clientMethods["kinesis_ListTagsForStream"] = kinesis_ListTagsForStream
	clientMethods["kinesis_MergeShards"] = kinesis_MergeShards
	clientMethods["kinesis_PutRecord"] = kinesis_PutRecord
	clientMethods["kinesis_PutRecords"] = kinesis_PutRecords
	clientMethods["kinesis_PutResourcePolicy"] = kinesis_PutResourcePolicy
	clientMethods["kinesis_RegisterStreamConsumer"] = kinesis_RegisterStreamConsumer
	clientMethods["kinesis_RemoveTagsFromStream"] = kinesis_RemoveTagsFromStream
	clientMethods["kinesis_SplitShard"] = kinesis_SplitShard
	clientMethods["kinesis_StartStreamEncryption"] = kinesis_StartStreamEncryption
	clientMethods["kinesis_StopStreamEncryption"] = kinesis_StopStreamEncryption
	clientMethods["kinesis_SubscribeToShard"] = kinesis_SubscribeToShard
	clientMethods["kinesis_UpdateShardCount"] = kinesis_UpdateShardCount
	clientMethods["kinesis_UpdateStreamMode"] = kinesis_UpdateStreamMode
}
