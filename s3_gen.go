package sdkclient

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func s3_AbortMultipartUpload(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := s3.NewFromConfig(awsCfg)
	var in s3.AbortMultipartUploadInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.AbortMultipartUpload(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call AbortMultipartUpload: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func s3_CompleteMultipartUpload(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := s3.NewFromConfig(awsCfg)
	var in s3.CompleteMultipartUploadInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.CompleteMultipartUpload(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call CompleteMultipartUpload: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func s3_CopyObject(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := s3.NewFromConfig(awsCfg)
	var in s3.CopyObjectInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.CopyObject(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call CopyObject: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func s3_CreateBucket(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := s3.NewFromConfig(awsCfg)
	var in s3.CreateBucketInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.CreateBucket(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call CreateBucket: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func s3_CreateMultipartUpload(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := s3.NewFromConfig(awsCfg)
	var in s3.CreateMultipartUploadInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.CreateMultipartUpload(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call CreateMultipartUpload: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func s3_CreateSession(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := s3.NewFromConfig(awsCfg)
	var in s3.CreateSessionInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.CreateSession(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call CreateSession: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func s3_DeleteBucket(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := s3.NewFromConfig(awsCfg)
	var in s3.DeleteBucketInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.DeleteBucket(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call DeleteBucket: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func s3_DeleteBucketAnalyticsConfiguration(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := s3.NewFromConfig(awsCfg)
	var in s3.DeleteBucketAnalyticsConfigurationInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.DeleteBucketAnalyticsConfiguration(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call DeleteBucketAnalyticsConfiguration: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func s3_DeleteBucketCors(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := s3.NewFromConfig(awsCfg)
	var in s3.DeleteBucketCorsInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.DeleteBucketCors(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call DeleteBucketCors: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func s3_DeleteBucketEncryption(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := s3.NewFromConfig(awsCfg)
	var in s3.DeleteBucketEncryptionInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.DeleteBucketEncryption(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call DeleteBucketEncryption: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func s3_DeleteBucketIntelligentTieringConfiguration(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := s3.NewFromConfig(awsCfg)
	var in s3.DeleteBucketIntelligentTieringConfigurationInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.DeleteBucketIntelligentTieringConfiguration(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call DeleteBucketIntelligentTieringConfiguration: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func s3_DeleteBucketInventoryConfiguration(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := s3.NewFromConfig(awsCfg)
	var in s3.DeleteBucketInventoryConfigurationInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.DeleteBucketInventoryConfiguration(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call DeleteBucketInventoryConfiguration: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func s3_DeleteBucketLifecycle(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := s3.NewFromConfig(awsCfg)
	var in s3.DeleteBucketLifecycleInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.DeleteBucketLifecycle(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call DeleteBucketLifecycle: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func s3_DeleteBucketMetricsConfiguration(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := s3.NewFromConfig(awsCfg)
	var in s3.DeleteBucketMetricsConfigurationInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.DeleteBucketMetricsConfiguration(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call DeleteBucketMetricsConfiguration: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func s3_DeleteBucketOwnershipControls(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := s3.NewFromConfig(awsCfg)
	var in s3.DeleteBucketOwnershipControlsInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.DeleteBucketOwnershipControls(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call DeleteBucketOwnershipControls: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func s3_DeleteBucketPolicy(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := s3.NewFromConfig(awsCfg)
	var in s3.DeleteBucketPolicyInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.DeleteBucketPolicy(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call DeleteBucketPolicy: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func s3_DeleteBucketReplication(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := s3.NewFromConfig(awsCfg)
	var in s3.DeleteBucketReplicationInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.DeleteBucketReplication(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call DeleteBucketReplication: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func s3_DeleteBucketTagging(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := s3.NewFromConfig(awsCfg)
	var in s3.DeleteBucketTaggingInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.DeleteBucketTagging(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call DeleteBucketTagging: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func s3_DeleteBucketWebsite(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := s3.NewFromConfig(awsCfg)
	var in s3.DeleteBucketWebsiteInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.DeleteBucketWebsite(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call DeleteBucketWebsite: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func s3_DeleteObject(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := s3.NewFromConfig(awsCfg)
	var in s3.DeleteObjectInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.DeleteObject(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call DeleteObject: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func s3_DeleteObjectTagging(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := s3.NewFromConfig(awsCfg)
	var in s3.DeleteObjectTaggingInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.DeleteObjectTagging(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call DeleteObjectTagging: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func s3_DeleteObjects(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := s3.NewFromConfig(awsCfg)
	var in s3.DeleteObjectsInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.DeleteObjects(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call DeleteObjects: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func s3_DeletePublicAccessBlock(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := s3.NewFromConfig(awsCfg)
	var in s3.DeletePublicAccessBlockInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.DeletePublicAccessBlock(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call DeletePublicAccessBlock: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func s3_GetBucketAccelerateConfiguration(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := s3.NewFromConfig(awsCfg)
	var in s3.GetBucketAccelerateConfigurationInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.GetBucketAccelerateConfiguration(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call GetBucketAccelerateConfiguration: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func s3_GetBucketAcl(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := s3.NewFromConfig(awsCfg)
	var in s3.GetBucketAclInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.GetBucketAcl(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call GetBucketAcl: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func s3_GetBucketAnalyticsConfiguration(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := s3.NewFromConfig(awsCfg)
	var in s3.GetBucketAnalyticsConfigurationInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.GetBucketAnalyticsConfiguration(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call GetBucketAnalyticsConfiguration: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func s3_GetBucketCors(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := s3.NewFromConfig(awsCfg)
	var in s3.GetBucketCorsInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.GetBucketCors(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call GetBucketCors: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func s3_GetBucketEncryption(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := s3.NewFromConfig(awsCfg)
	var in s3.GetBucketEncryptionInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.GetBucketEncryption(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call GetBucketEncryption: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func s3_GetBucketIntelligentTieringConfiguration(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := s3.NewFromConfig(awsCfg)
	var in s3.GetBucketIntelligentTieringConfigurationInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.GetBucketIntelligentTieringConfiguration(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call GetBucketIntelligentTieringConfiguration: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func s3_GetBucketInventoryConfiguration(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := s3.NewFromConfig(awsCfg)
	var in s3.GetBucketInventoryConfigurationInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.GetBucketInventoryConfiguration(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call GetBucketInventoryConfiguration: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func s3_GetBucketLifecycleConfiguration(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := s3.NewFromConfig(awsCfg)
	var in s3.GetBucketLifecycleConfigurationInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.GetBucketLifecycleConfiguration(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call GetBucketLifecycleConfiguration: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func s3_GetBucketLocation(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := s3.NewFromConfig(awsCfg)
	var in s3.GetBucketLocationInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.GetBucketLocation(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call GetBucketLocation: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func s3_GetBucketLogging(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := s3.NewFromConfig(awsCfg)
	var in s3.GetBucketLoggingInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.GetBucketLogging(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call GetBucketLogging: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func s3_GetBucketMetricsConfiguration(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := s3.NewFromConfig(awsCfg)
	var in s3.GetBucketMetricsConfigurationInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.GetBucketMetricsConfiguration(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call GetBucketMetricsConfiguration: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func s3_GetBucketNotificationConfiguration(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := s3.NewFromConfig(awsCfg)
	var in s3.GetBucketNotificationConfigurationInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.GetBucketNotificationConfiguration(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call GetBucketNotificationConfiguration: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func s3_GetBucketOwnershipControls(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := s3.NewFromConfig(awsCfg)
	var in s3.GetBucketOwnershipControlsInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.GetBucketOwnershipControls(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call GetBucketOwnershipControls: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func s3_GetBucketPolicy(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := s3.NewFromConfig(awsCfg)
	var in s3.GetBucketPolicyInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.GetBucketPolicy(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call GetBucketPolicy: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func s3_GetBucketPolicyStatus(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := s3.NewFromConfig(awsCfg)
	var in s3.GetBucketPolicyStatusInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.GetBucketPolicyStatus(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call GetBucketPolicyStatus: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func s3_GetBucketReplication(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := s3.NewFromConfig(awsCfg)
	var in s3.GetBucketReplicationInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.GetBucketReplication(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call GetBucketReplication: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func s3_GetBucketRequestPayment(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := s3.NewFromConfig(awsCfg)
	var in s3.GetBucketRequestPaymentInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.GetBucketRequestPayment(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call GetBucketRequestPayment: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func s3_GetBucketTagging(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := s3.NewFromConfig(awsCfg)
	var in s3.GetBucketTaggingInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.GetBucketTagging(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call GetBucketTagging: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func s3_GetBucketVersioning(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := s3.NewFromConfig(awsCfg)
	var in s3.GetBucketVersioningInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.GetBucketVersioning(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call GetBucketVersioning: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func s3_GetBucketWebsite(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := s3.NewFromConfig(awsCfg)
	var in s3.GetBucketWebsiteInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.GetBucketWebsite(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call GetBucketWebsite: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func s3_GetObject(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := s3.NewFromConfig(awsCfg)
	var in s3.GetObjectInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.GetObject(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call GetObject: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func s3_GetObjectAcl(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := s3.NewFromConfig(awsCfg)
	var in s3.GetObjectAclInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.GetObjectAcl(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call GetObjectAcl: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func s3_GetObjectAttributes(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := s3.NewFromConfig(awsCfg)
	var in s3.GetObjectAttributesInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.GetObjectAttributes(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call GetObjectAttributes: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func s3_GetObjectLegalHold(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := s3.NewFromConfig(awsCfg)
	var in s3.GetObjectLegalHoldInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.GetObjectLegalHold(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call GetObjectLegalHold: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func s3_GetObjectLockConfiguration(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := s3.NewFromConfig(awsCfg)
	var in s3.GetObjectLockConfigurationInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.GetObjectLockConfiguration(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call GetObjectLockConfiguration: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func s3_GetObjectRetention(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := s3.NewFromConfig(awsCfg)
	var in s3.GetObjectRetentionInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.GetObjectRetention(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call GetObjectRetention: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func s3_GetObjectTagging(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := s3.NewFromConfig(awsCfg)
	var in s3.GetObjectTaggingInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.GetObjectTagging(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call GetObjectTagging: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func s3_GetObjectTorrent(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := s3.NewFromConfig(awsCfg)
	var in s3.GetObjectTorrentInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.GetObjectTorrent(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call GetObjectTorrent: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func s3_GetPublicAccessBlock(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := s3.NewFromConfig(awsCfg)
	var in s3.GetPublicAccessBlockInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.GetPublicAccessBlock(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call GetPublicAccessBlock: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func s3_HeadBucket(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := s3.NewFromConfig(awsCfg)
	var in s3.HeadBucketInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.HeadBucket(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call HeadBucket: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func s3_HeadObject(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := s3.NewFromConfig(awsCfg)
	var in s3.HeadObjectInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.HeadObject(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call HeadObject: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func s3_ListBucketAnalyticsConfigurations(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := s3.NewFromConfig(awsCfg)
	var in s3.ListBucketAnalyticsConfigurationsInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.ListBucketAnalyticsConfigurations(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call ListBucketAnalyticsConfigurations: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func s3_ListBucketIntelligentTieringConfigurations(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := s3.NewFromConfig(awsCfg)
	var in s3.ListBucketIntelligentTieringConfigurationsInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.ListBucketIntelligentTieringConfigurations(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call ListBucketIntelligentTieringConfigurations: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func s3_ListBucketInventoryConfigurations(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := s3.NewFromConfig(awsCfg)
	var in s3.ListBucketInventoryConfigurationsInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.ListBucketInventoryConfigurations(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call ListBucketInventoryConfigurations: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func s3_ListBucketMetricsConfigurations(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := s3.NewFromConfig(awsCfg)
	var in s3.ListBucketMetricsConfigurationsInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.ListBucketMetricsConfigurations(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call ListBucketMetricsConfigurations: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func s3_ListBuckets(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := s3.NewFromConfig(awsCfg)
	var in s3.ListBucketsInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.ListBuckets(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call ListBuckets: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func s3_ListDirectoryBuckets(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := s3.NewFromConfig(awsCfg)
	var in s3.ListDirectoryBucketsInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.ListDirectoryBuckets(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call ListDirectoryBuckets: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func s3_ListMultipartUploads(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := s3.NewFromConfig(awsCfg)
	var in s3.ListMultipartUploadsInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.ListMultipartUploads(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call ListMultipartUploads: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func s3_ListObjectVersions(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := s3.NewFromConfig(awsCfg)
	var in s3.ListObjectVersionsInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.ListObjectVersions(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call ListObjectVersions: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func s3_ListObjects(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := s3.NewFromConfig(awsCfg)
	var in s3.ListObjectsInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.ListObjects(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call ListObjects: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func s3_ListObjectsV2(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := s3.NewFromConfig(awsCfg)
	var in s3.ListObjectsV2Input
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.ListObjectsV2(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call ListObjectsV2: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func s3_ListParts(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := s3.NewFromConfig(awsCfg)
	var in s3.ListPartsInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.ListParts(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call ListParts: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func s3_PutBucketAccelerateConfiguration(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := s3.NewFromConfig(awsCfg)
	var in s3.PutBucketAccelerateConfigurationInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.PutBucketAccelerateConfiguration(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call PutBucketAccelerateConfiguration: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func s3_PutBucketAcl(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := s3.NewFromConfig(awsCfg)
	var in s3.PutBucketAclInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.PutBucketAcl(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call PutBucketAcl: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func s3_PutBucketAnalyticsConfiguration(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := s3.NewFromConfig(awsCfg)
	var in s3.PutBucketAnalyticsConfigurationInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.PutBucketAnalyticsConfiguration(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call PutBucketAnalyticsConfiguration: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func s3_PutBucketCors(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := s3.NewFromConfig(awsCfg)
	var in s3.PutBucketCorsInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.PutBucketCors(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call PutBucketCors: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func s3_PutBucketEncryption(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := s3.NewFromConfig(awsCfg)
	var in s3.PutBucketEncryptionInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.PutBucketEncryption(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call PutBucketEncryption: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func s3_PutBucketIntelligentTieringConfiguration(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := s3.NewFromConfig(awsCfg)
	var in s3.PutBucketIntelligentTieringConfigurationInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.PutBucketIntelligentTieringConfiguration(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call PutBucketIntelligentTieringConfiguration: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func s3_PutBucketInventoryConfiguration(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := s3.NewFromConfig(awsCfg)
	var in s3.PutBucketInventoryConfigurationInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.PutBucketInventoryConfiguration(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call PutBucketInventoryConfiguration: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func s3_PutBucketLifecycleConfiguration(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := s3.NewFromConfig(awsCfg)
	var in s3.PutBucketLifecycleConfigurationInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.PutBucketLifecycleConfiguration(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call PutBucketLifecycleConfiguration: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func s3_PutBucketLogging(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := s3.NewFromConfig(awsCfg)
	var in s3.PutBucketLoggingInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.PutBucketLogging(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call PutBucketLogging: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func s3_PutBucketMetricsConfiguration(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := s3.NewFromConfig(awsCfg)
	var in s3.PutBucketMetricsConfigurationInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.PutBucketMetricsConfiguration(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call PutBucketMetricsConfiguration: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func s3_PutBucketNotificationConfiguration(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := s3.NewFromConfig(awsCfg)
	var in s3.PutBucketNotificationConfigurationInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.PutBucketNotificationConfiguration(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call PutBucketNotificationConfiguration: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func s3_PutBucketOwnershipControls(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := s3.NewFromConfig(awsCfg)
	var in s3.PutBucketOwnershipControlsInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.PutBucketOwnershipControls(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call PutBucketOwnershipControls: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func s3_PutBucketPolicy(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := s3.NewFromConfig(awsCfg)
	var in s3.PutBucketPolicyInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.PutBucketPolicy(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call PutBucketPolicy: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func s3_PutBucketReplication(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := s3.NewFromConfig(awsCfg)
	var in s3.PutBucketReplicationInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.PutBucketReplication(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call PutBucketReplication: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func s3_PutBucketRequestPayment(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := s3.NewFromConfig(awsCfg)
	var in s3.PutBucketRequestPaymentInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.PutBucketRequestPayment(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call PutBucketRequestPayment: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func s3_PutBucketTagging(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := s3.NewFromConfig(awsCfg)
	var in s3.PutBucketTaggingInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.PutBucketTagging(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call PutBucketTagging: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func s3_PutBucketVersioning(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := s3.NewFromConfig(awsCfg)
	var in s3.PutBucketVersioningInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.PutBucketVersioning(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call PutBucketVersioning: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func s3_PutBucketWebsite(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := s3.NewFromConfig(awsCfg)
	var in s3.PutBucketWebsiteInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.PutBucketWebsite(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call PutBucketWebsite: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func s3_PutObject(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := s3.NewFromConfig(awsCfg)
	var in s3.PutObjectInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.PutObject(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call PutObject: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func s3_PutObjectAcl(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := s3.NewFromConfig(awsCfg)
	var in s3.PutObjectAclInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.PutObjectAcl(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call PutObjectAcl: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func s3_PutObjectLegalHold(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := s3.NewFromConfig(awsCfg)
	var in s3.PutObjectLegalHoldInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.PutObjectLegalHold(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call PutObjectLegalHold: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func s3_PutObjectLockConfiguration(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := s3.NewFromConfig(awsCfg)
	var in s3.PutObjectLockConfigurationInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.PutObjectLockConfiguration(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call PutObjectLockConfiguration: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func s3_PutObjectRetention(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := s3.NewFromConfig(awsCfg)
	var in s3.PutObjectRetentionInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.PutObjectRetention(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call PutObjectRetention: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func s3_PutObjectTagging(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := s3.NewFromConfig(awsCfg)
	var in s3.PutObjectTaggingInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.PutObjectTagging(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call PutObjectTagging: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func s3_PutPublicAccessBlock(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := s3.NewFromConfig(awsCfg)
	var in s3.PutPublicAccessBlockInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.PutPublicAccessBlock(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call PutPublicAccessBlock: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func s3_RestoreObject(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := s3.NewFromConfig(awsCfg)
	var in s3.RestoreObjectInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.RestoreObject(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call RestoreObject: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func s3_SelectObjectContent(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := s3.NewFromConfig(awsCfg)
	var in s3.SelectObjectContentInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.SelectObjectContent(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call SelectObjectContent: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func s3_UploadPart(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := s3.NewFromConfig(awsCfg)
	var in s3.UploadPartInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.UploadPart(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call UploadPart: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func s3_UploadPartCopy(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := s3.NewFromConfig(awsCfg)
	var in s3.UploadPartCopyInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.UploadPartCopy(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call UploadPartCopy: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func s3_WriteGetObjectResponse(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := s3.NewFromConfig(awsCfg)
	var in s3.WriteGetObjectResponseInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.WriteGetObjectResponse(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call WriteGetObjectResponse: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func init() {
	clientMethods["s3_AbortMultipartUpload"] = s3_AbortMultipartUpload
	clientMethods["s3_CompleteMultipartUpload"] = s3_CompleteMultipartUpload
	clientMethods["s3_CopyObject"] = s3_CopyObject
	clientMethods["s3_CreateBucket"] = s3_CreateBucket
	clientMethods["s3_CreateMultipartUpload"] = s3_CreateMultipartUpload
	clientMethods["s3_CreateSession"] = s3_CreateSession
	clientMethods["s3_DeleteBucket"] = s3_DeleteBucket
	clientMethods["s3_DeleteBucketAnalyticsConfiguration"] = s3_DeleteBucketAnalyticsConfiguration
	clientMethods["s3_DeleteBucketCors"] = s3_DeleteBucketCors
	clientMethods["s3_DeleteBucketEncryption"] = s3_DeleteBucketEncryption
	clientMethods["s3_DeleteBucketIntelligentTieringConfiguration"] = s3_DeleteBucketIntelligentTieringConfiguration
	clientMethods["s3_DeleteBucketInventoryConfiguration"] = s3_DeleteBucketInventoryConfiguration
	clientMethods["s3_DeleteBucketLifecycle"] = s3_DeleteBucketLifecycle
	clientMethods["s3_DeleteBucketMetricsConfiguration"] = s3_DeleteBucketMetricsConfiguration
	clientMethods["s3_DeleteBucketOwnershipControls"] = s3_DeleteBucketOwnershipControls
	clientMethods["s3_DeleteBucketPolicy"] = s3_DeleteBucketPolicy
	clientMethods["s3_DeleteBucketReplication"] = s3_DeleteBucketReplication
	clientMethods["s3_DeleteBucketTagging"] = s3_DeleteBucketTagging
	clientMethods["s3_DeleteBucketWebsite"] = s3_DeleteBucketWebsite
	clientMethods["s3_DeleteObject"] = s3_DeleteObject
	clientMethods["s3_DeleteObjectTagging"] = s3_DeleteObjectTagging
	clientMethods["s3_DeleteObjects"] = s3_DeleteObjects
	clientMethods["s3_DeletePublicAccessBlock"] = s3_DeletePublicAccessBlock
	clientMethods["s3_GetBucketAccelerateConfiguration"] = s3_GetBucketAccelerateConfiguration
	clientMethods["s3_GetBucketAcl"] = s3_GetBucketAcl
	clientMethods["s3_GetBucketAnalyticsConfiguration"] = s3_GetBucketAnalyticsConfiguration
	clientMethods["s3_GetBucketCors"] = s3_GetBucketCors
	clientMethods["s3_GetBucketEncryption"] = s3_GetBucketEncryption
	clientMethods["s3_GetBucketIntelligentTieringConfiguration"] = s3_GetBucketIntelligentTieringConfiguration
	clientMethods["s3_GetBucketInventoryConfiguration"] = s3_GetBucketInventoryConfiguration
	clientMethods["s3_GetBucketLifecycleConfiguration"] = s3_GetBucketLifecycleConfiguration
	clientMethods["s3_GetBucketLocation"] = s3_GetBucketLocation
	clientMethods["s3_GetBucketLogging"] = s3_GetBucketLogging
	clientMethods["s3_GetBucketMetricsConfiguration"] = s3_GetBucketMetricsConfiguration
	clientMethods["s3_GetBucketNotificationConfiguration"] = s3_GetBucketNotificationConfiguration
	clientMethods["s3_GetBucketOwnershipControls"] = s3_GetBucketOwnershipControls
	clientMethods["s3_GetBucketPolicy"] = s3_GetBucketPolicy
	clientMethods["s3_GetBucketPolicyStatus"] = s3_GetBucketPolicyStatus
	clientMethods["s3_GetBucketReplication"] = s3_GetBucketReplication
	clientMethods["s3_GetBucketRequestPayment"] = s3_GetBucketRequestPayment
	clientMethods["s3_GetBucketTagging"] = s3_GetBucketTagging
	clientMethods["s3_GetBucketVersioning"] = s3_GetBucketVersioning
	clientMethods["s3_GetBucketWebsite"] = s3_GetBucketWebsite
	clientMethods["s3_GetObject"] = s3_GetObject
	clientMethods["s3_GetObjectAcl"] = s3_GetObjectAcl
	clientMethods["s3_GetObjectAttributes"] = s3_GetObjectAttributes
	clientMethods["s3_GetObjectLegalHold"] = s3_GetObjectLegalHold
	clientMethods["s3_GetObjectLockConfiguration"] = s3_GetObjectLockConfiguration
	clientMethods["s3_GetObjectRetention"] = s3_GetObjectRetention
	clientMethods["s3_GetObjectTagging"] = s3_GetObjectTagging
	clientMethods["s3_GetObjectTorrent"] = s3_GetObjectTorrent
	clientMethods["s3_GetPublicAccessBlock"] = s3_GetPublicAccessBlock
	clientMethods["s3_HeadBucket"] = s3_HeadBucket
	clientMethods["s3_HeadObject"] = s3_HeadObject
	clientMethods["s3_ListBucketAnalyticsConfigurations"] = s3_ListBucketAnalyticsConfigurations
	clientMethods["s3_ListBucketIntelligentTieringConfigurations"] = s3_ListBucketIntelligentTieringConfigurations
	clientMethods["s3_ListBucketInventoryConfigurations"] = s3_ListBucketInventoryConfigurations
	clientMethods["s3_ListBucketMetricsConfigurations"] = s3_ListBucketMetricsConfigurations
	clientMethods["s3_ListBuckets"] = s3_ListBuckets
	clientMethods["s3_ListDirectoryBuckets"] = s3_ListDirectoryBuckets
	clientMethods["s3_ListMultipartUploads"] = s3_ListMultipartUploads
	clientMethods["s3_ListObjectVersions"] = s3_ListObjectVersions
	clientMethods["s3_ListObjects"] = s3_ListObjects
	clientMethods["s3_ListObjectsV2"] = s3_ListObjectsV2
	clientMethods["s3_ListParts"] = s3_ListParts
	clientMethods["s3_PutBucketAccelerateConfiguration"] = s3_PutBucketAccelerateConfiguration
	clientMethods["s3_PutBucketAcl"] = s3_PutBucketAcl
	clientMethods["s3_PutBucketAnalyticsConfiguration"] = s3_PutBucketAnalyticsConfiguration
	clientMethods["s3_PutBucketCors"] = s3_PutBucketCors
	clientMethods["s3_PutBucketEncryption"] = s3_PutBucketEncryption
	clientMethods["s3_PutBucketIntelligentTieringConfiguration"] = s3_PutBucketIntelligentTieringConfiguration
	clientMethods["s3_PutBucketInventoryConfiguration"] = s3_PutBucketInventoryConfiguration
	clientMethods["s3_PutBucketLifecycleConfiguration"] = s3_PutBucketLifecycleConfiguration
	clientMethods["s3_PutBucketLogging"] = s3_PutBucketLogging
	clientMethods["s3_PutBucketMetricsConfiguration"] = s3_PutBucketMetricsConfiguration
	clientMethods["s3_PutBucketNotificationConfiguration"] = s3_PutBucketNotificationConfiguration
	clientMethods["s3_PutBucketOwnershipControls"] = s3_PutBucketOwnershipControls
	clientMethods["s3_PutBucketPolicy"] = s3_PutBucketPolicy
	clientMethods["s3_PutBucketReplication"] = s3_PutBucketReplication
	clientMethods["s3_PutBucketRequestPayment"] = s3_PutBucketRequestPayment
	clientMethods["s3_PutBucketTagging"] = s3_PutBucketTagging
	clientMethods["s3_PutBucketVersioning"] = s3_PutBucketVersioning
	clientMethods["s3_PutBucketWebsite"] = s3_PutBucketWebsite
	clientMethods["s3_PutObject"] = s3_PutObject
	clientMethods["s3_PutObjectAcl"] = s3_PutObjectAcl
	clientMethods["s3_PutObjectLegalHold"] = s3_PutObjectLegalHold
	clientMethods["s3_PutObjectLockConfiguration"] = s3_PutObjectLockConfiguration
	clientMethods["s3_PutObjectRetention"] = s3_PutObjectRetention
	clientMethods["s3_PutObjectTagging"] = s3_PutObjectTagging
	clientMethods["s3_PutPublicAccessBlock"] = s3_PutPublicAccessBlock
	clientMethods["s3_RestoreObject"] = s3_RestoreObject
	clientMethods["s3_SelectObjectContent"] = s3_SelectObjectContent
	clientMethods["s3_UploadPart"] = s3_UploadPart
	clientMethods["s3_UploadPartCopy"] = s3_UploadPartCopy
	clientMethods["s3_WriteGetObjectResponse"] = s3_WriteGetObjectResponse
}
