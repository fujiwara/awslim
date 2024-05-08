package sdkclient

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
)

func ssm_AddTagsToResource(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ssm.NewFromConfig(awsCfg)
	var in ssm.AddTagsToResourceInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.AddTagsToResource(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call AddTagsToResource: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ssm_AssociateOpsItemRelatedItem(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ssm.NewFromConfig(awsCfg)
	var in ssm.AssociateOpsItemRelatedItemInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.AssociateOpsItemRelatedItem(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call AssociateOpsItemRelatedItem: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ssm_CancelCommand(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ssm.NewFromConfig(awsCfg)
	var in ssm.CancelCommandInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.CancelCommand(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call CancelCommand: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ssm_CancelMaintenanceWindowExecution(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ssm.NewFromConfig(awsCfg)
	var in ssm.CancelMaintenanceWindowExecutionInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.CancelMaintenanceWindowExecution(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call CancelMaintenanceWindowExecution: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ssm_CreateActivation(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ssm.NewFromConfig(awsCfg)
	var in ssm.CreateActivationInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.CreateActivation(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call CreateActivation: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ssm_CreateAssociation(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ssm.NewFromConfig(awsCfg)
	var in ssm.CreateAssociationInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.CreateAssociation(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call CreateAssociation: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ssm_CreateAssociationBatch(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ssm.NewFromConfig(awsCfg)
	var in ssm.CreateAssociationBatchInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.CreateAssociationBatch(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call CreateAssociationBatch: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ssm_CreateDocument(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ssm.NewFromConfig(awsCfg)
	var in ssm.CreateDocumentInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.CreateDocument(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call CreateDocument: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ssm_CreateMaintenanceWindow(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ssm.NewFromConfig(awsCfg)
	var in ssm.CreateMaintenanceWindowInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.CreateMaintenanceWindow(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call CreateMaintenanceWindow: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ssm_CreateOpsItem(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ssm.NewFromConfig(awsCfg)
	var in ssm.CreateOpsItemInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.CreateOpsItem(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call CreateOpsItem: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ssm_CreateOpsMetadata(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ssm.NewFromConfig(awsCfg)
	var in ssm.CreateOpsMetadataInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.CreateOpsMetadata(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call CreateOpsMetadata: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ssm_CreatePatchBaseline(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ssm.NewFromConfig(awsCfg)
	var in ssm.CreatePatchBaselineInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.CreatePatchBaseline(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call CreatePatchBaseline: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ssm_CreateResourceDataSync(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ssm.NewFromConfig(awsCfg)
	var in ssm.CreateResourceDataSyncInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.CreateResourceDataSync(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call CreateResourceDataSync: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ssm_DeleteActivation(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ssm.NewFromConfig(awsCfg)
	var in ssm.DeleteActivationInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.DeleteActivation(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call DeleteActivation: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ssm_DeleteAssociation(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ssm.NewFromConfig(awsCfg)
	var in ssm.DeleteAssociationInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.DeleteAssociation(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call DeleteAssociation: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ssm_DeleteDocument(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ssm.NewFromConfig(awsCfg)
	var in ssm.DeleteDocumentInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.DeleteDocument(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call DeleteDocument: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ssm_DeleteInventory(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ssm.NewFromConfig(awsCfg)
	var in ssm.DeleteInventoryInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.DeleteInventory(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call DeleteInventory: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ssm_DeleteMaintenanceWindow(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ssm.NewFromConfig(awsCfg)
	var in ssm.DeleteMaintenanceWindowInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.DeleteMaintenanceWindow(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call DeleteMaintenanceWindow: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ssm_DeleteOpsItem(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ssm.NewFromConfig(awsCfg)
	var in ssm.DeleteOpsItemInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.DeleteOpsItem(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call DeleteOpsItem: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ssm_DeleteOpsMetadata(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ssm.NewFromConfig(awsCfg)
	var in ssm.DeleteOpsMetadataInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.DeleteOpsMetadata(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call DeleteOpsMetadata: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ssm_DeleteParameter(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ssm.NewFromConfig(awsCfg)
	var in ssm.DeleteParameterInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.DeleteParameter(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call DeleteParameter: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ssm_DeleteParameters(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ssm.NewFromConfig(awsCfg)
	var in ssm.DeleteParametersInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.DeleteParameters(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call DeleteParameters: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ssm_DeletePatchBaseline(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ssm.NewFromConfig(awsCfg)
	var in ssm.DeletePatchBaselineInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.DeletePatchBaseline(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call DeletePatchBaseline: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ssm_DeleteResourceDataSync(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ssm.NewFromConfig(awsCfg)
	var in ssm.DeleteResourceDataSyncInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.DeleteResourceDataSync(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call DeleteResourceDataSync: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ssm_DeleteResourcePolicy(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ssm.NewFromConfig(awsCfg)
	var in ssm.DeleteResourcePolicyInput
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

func ssm_DeregisterManagedInstance(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ssm.NewFromConfig(awsCfg)
	var in ssm.DeregisterManagedInstanceInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.DeregisterManagedInstance(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call DeregisterManagedInstance: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ssm_DeregisterPatchBaselineForPatchGroup(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ssm.NewFromConfig(awsCfg)
	var in ssm.DeregisterPatchBaselineForPatchGroupInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.DeregisterPatchBaselineForPatchGroup(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call DeregisterPatchBaselineForPatchGroup: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ssm_DeregisterTargetFromMaintenanceWindow(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ssm.NewFromConfig(awsCfg)
	var in ssm.DeregisterTargetFromMaintenanceWindowInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.DeregisterTargetFromMaintenanceWindow(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call DeregisterTargetFromMaintenanceWindow: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ssm_DeregisterTaskFromMaintenanceWindow(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ssm.NewFromConfig(awsCfg)
	var in ssm.DeregisterTaskFromMaintenanceWindowInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.DeregisterTaskFromMaintenanceWindow(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call DeregisterTaskFromMaintenanceWindow: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ssm_DescribeActivations(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ssm.NewFromConfig(awsCfg)
	var in ssm.DescribeActivationsInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.DescribeActivations(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call DescribeActivations: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ssm_DescribeAssociation(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ssm.NewFromConfig(awsCfg)
	var in ssm.DescribeAssociationInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.DescribeAssociation(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call DescribeAssociation: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ssm_DescribeAssociationExecutionTargets(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ssm.NewFromConfig(awsCfg)
	var in ssm.DescribeAssociationExecutionTargetsInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.DescribeAssociationExecutionTargets(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call DescribeAssociationExecutionTargets: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ssm_DescribeAssociationExecutions(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ssm.NewFromConfig(awsCfg)
	var in ssm.DescribeAssociationExecutionsInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.DescribeAssociationExecutions(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call DescribeAssociationExecutions: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ssm_DescribeAutomationExecutions(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ssm.NewFromConfig(awsCfg)
	var in ssm.DescribeAutomationExecutionsInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.DescribeAutomationExecutions(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call DescribeAutomationExecutions: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ssm_DescribeAutomationStepExecutions(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ssm.NewFromConfig(awsCfg)
	var in ssm.DescribeAutomationStepExecutionsInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.DescribeAutomationStepExecutions(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call DescribeAutomationStepExecutions: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ssm_DescribeAvailablePatches(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ssm.NewFromConfig(awsCfg)
	var in ssm.DescribeAvailablePatchesInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.DescribeAvailablePatches(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call DescribeAvailablePatches: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ssm_DescribeDocument(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ssm.NewFromConfig(awsCfg)
	var in ssm.DescribeDocumentInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.DescribeDocument(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call DescribeDocument: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ssm_DescribeDocumentPermission(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ssm.NewFromConfig(awsCfg)
	var in ssm.DescribeDocumentPermissionInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.DescribeDocumentPermission(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call DescribeDocumentPermission: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ssm_DescribeEffectiveInstanceAssociations(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ssm.NewFromConfig(awsCfg)
	var in ssm.DescribeEffectiveInstanceAssociationsInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.DescribeEffectiveInstanceAssociations(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call DescribeEffectiveInstanceAssociations: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ssm_DescribeEffectivePatchesForPatchBaseline(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ssm.NewFromConfig(awsCfg)
	var in ssm.DescribeEffectivePatchesForPatchBaselineInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.DescribeEffectivePatchesForPatchBaseline(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call DescribeEffectivePatchesForPatchBaseline: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ssm_DescribeInstanceAssociationsStatus(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ssm.NewFromConfig(awsCfg)
	var in ssm.DescribeInstanceAssociationsStatusInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.DescribeInstanceAssociationsStatus(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call DescribeInstanceAssociationsStatus: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ssm_DescribeInstanceInformation(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ssm.NewFromConfig(awsCfg)
	var in ssm.DescribeInstanceInformationInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.DescribeInstanceInformation(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call DescribeInstanceInformation: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ssm_DescribeInstancePatchStates(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ssm.NewFromConfig(awsCfg)
	var in ssm.DescribeInstancePatchStatesInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.DescribeInstancePatchStates(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call DescribeInstancePatchStates: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ssm_DescribeInstancePatchStatesForPatchGroup(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ssm.NewFromConfig(awsCfg)
	var in ssm.DescribeInstancePatchStatesForPatchGroupInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.DescribeInstancePatchStatesForPatchGroup(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call DescribeInstancePatchStatesForPatchGroup: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ssm_DescribeInstancePatches(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ssm.NewFromConfig(awsCfg)
	var in ssm.DescribeInstancePatchesInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.DescribeInstancePatches(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call DescribeInstancePatches: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ssm_DescribeInstanceProperties(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ssm.NewFromConfig(awsCfg)
	var in ssm.DescribeInstancePropertiesInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.DescribeInstanceProperties(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call DescribeInstanceProperties: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ssm_DescribeInventoryDeletions(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ssm.NewFromConfig(awsCfg)
	var in ssm.DescribeInventoryDeletionsInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.DescribeInventoryDeletions(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call DescribeInventoryDeletions: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ssm_DescribeMaintenanceWindowExecutionTaskInvocations(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ssm.NewFromConfig(awsCfg)
	var in ssm.DescribeMaintenanceWindowExecutionTaskInvocationsInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.DescribeMaintenanceWindowExecutionTaskInvocations(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call DescribeMaintenanceWindowExecutionTaskInvocations: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ssm_DescribeMaintenanceWindowExecutionTasks(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ssm.NewFromConfig(awsCfg)
	var in ssm.DescribeMaintenanceWindowExecutionTasksInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.DescribeMaintenanceWindowExecutionTasks(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call DescribeMaintenanceWindowExecutionTasks: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ssm_DescribeMaintenanceWindowExecutions(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ssm.NewFromConfig(awsCfg)
	var in ssm.DescribeMaintenanceWindowExecutionsInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.DescribeMaintenanceWindowExecutions(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call DescribeMaintenanceWindowExecutions: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ssm_DescribeMaintenanceWindowSchedule(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ssm.NewFromConfig(awsCfg)
	var in ssm.DescribeMaintenanceWindowScheduleInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.DescribeMaintenanceWindowSchedule(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call DescribeMaintenanceWindowSchedule: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ssm_DescribeMaintenanceWindowTargets(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ssm.NewFromConfig(awsCfg)
	var in ssm.DescribeMaintenanceWindowTargetsInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.DescribeMaintenanceWindowTargets(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call DescribeMaintenanceWindowTargets: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ssm_DescribeMaintenanceWindowTasks(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ssm.NewFromConfig(awsCfg)
	var in ssm.DescribeMaintenanceWindowTasksInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.DescribeMaintenanceWindowTasks(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call DescribeMaintenanceWindowTasks: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ssm_DescribeMaintenanceWindows(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ssm.NewFromConfig(awsCfg)
	var in ssm.DescribeMaintenanceWindowsInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.DescribeMaintenanceWindows(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call DescribeMaintenanceWindows: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ssm_DescribeMaintenanceWindowsForTarget(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ssm.NewFromConfig(awsCfg)
	var in ssm.DescribeMaintenanceWindowsForTargetInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.DescribeMaintenanceWindowsForTarget(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call DescribeMaintenanceWindowsForTarget: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ssm_DescribeOpsItems(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ssm.NewFromConfig(awsCfg)
	var in ssm.DescribeOpsItemsInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.DescribeOpsItems(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call DescribeOpsItems: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ssm_DescribeParameters(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ssm.NewFromConfig(awsCfg)
	var in ssm.DescribeParametersInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.DescribeParameters(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call DescribeParameters: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ssm_DescribePatchBaselines(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ssm.NewFromConfig(awsCfg)
	var in ssm.DescribePatchBaselinesInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.DescribePatchBaselines(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call DescribePatchBaselines: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ssm_DescribePatchGroupState(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ssm.NewFromConfig(awsCfg)
	var in ssm.DescribePatchGroupStateInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.DescribePatchGroupState(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call DescribePatchGroupState: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ssm_DescribePatchGroups(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ssm.NewFromConfig(awsCfg)
	var in ssm.DescribePatchGroupsInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.DescribePatchGroups(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call DescribePatchGroups: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ssm_DescribePatchProperties(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ssm.NewFromConfig(awsCfg)
	var in ssm.DescribePatchPropertiesInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.DescribePatchProperties(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call DescribePatchProperties: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ssm_DescribeSessions(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ssm.NewFromConfig(awsCfg)
	var in ssm.DescribeSessionsInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.DescribeSessions(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call DescribeSessions: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ssm_DisassociateOpsItemRelatedItem(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ssm.NewFromConfig(awsCfg)
	var in ssm.DisassociateOpsItemRelatedItemInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.DisassociateOpsItemRelatedItem(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call DisassociateOpsItemRelatedItem: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ssm_GetAutomationExecution(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ssm.NewFromConfig(awsCfg)
	var in ssm.GetAutomationExecutionInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.GetAutomationExecution(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call GetAutomationExecution: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ssm_GetCalendarState(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ssm.NewFromConfig(awsCfg)
	var in ssm.GetCalendarStateInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.GetCalendarState(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call GetCalendarState: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ssm_GetCommandInvocation(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ssm.NewFromConfig(awsCfg)
	var in ssm.GetCommandInvocationInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.GetCommandInvocation(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call GetCommandInvocation: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ssm_GetConnectionStatus(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ssm.NewFromConfig(awsCfg)
	var in ssm.GetConnectionStatusInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.GetConnectionStatus(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call GetConnectionStatus: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ssm_GetDefaultPatchBaseline(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ssm.NewFromConfig(awsCfg)
	var in ssm.GetDefaultPatchBaselineInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.GetDefaultPatchBaseline(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call GetDefaultPatchBaseline: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ssm_GetDeployablePatchSnapshotForInstance(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ssm.NewFromConfig(awsCfg)
	var in ssm.GetDeployablePatchSnapshotForInstanceInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.GetDeployablePatchSnapshotForInstance(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call GetDeployablePatchSnapshotForInstance: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ssm_GetDocument(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ssm.NewFromConfig(awsCfg)
	var in ssm.GetDocumentInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.GetDocument(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call GetDocument: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ssm_GetInventory(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ssm.NewFromConfig(awsCfg)
	var in ssm.GetInventoryInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.GetInventory(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call GetInventory: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ssm_GetInventorySchema(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ssm.NewFromConfig(awsCfg)
	var in ssm.GetInventorySchemaInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.GetInventorySchema(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call GetInventorySchema: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ssm_GetMaintenanceWindow(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ssm.NewFromConfig(awsCfg)
	var in ssm.GetMaintenanceWindowInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.GetMaintenanceWindow(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call GetMaintenanceWindow: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ssm_GetMaintenanceWindowExecution(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ssm.NewFromConfig(awsCfg)
	var in ssm.GetMaintenanceWindowExecutionInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.GetMaintenanceWindowExecution(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call GetMaintenanceWindowExecution: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ssm_GetMaintenanceWindowExecutionTask(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ssm.NewFromConfig(awsCfg)
	var in ssm.GetMaintenanceWindowExecutionTaskInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.GetMaintenanceWindowExecutionTask(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call GetMaintenanceWindowExecutionTask: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ssm_GetMaintenanceWindowExecutionTaskInvocation(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ssm.NewFromConfig(awsCfg)
	var in ssm.GetMaintenanceWindowExecutionTaskInvocationInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.GetMaintenanceWindowExecutionTaskInvocation(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call GetMaintenanceWindowExecutionTaskInvocation: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ssm_GetMaintenanceWindowTask(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ssm.NewFromConfig(awsCfg)
	var in ssm.GetMaintenanceWindowTaskInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.GetMaintenanceWindowTask(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call GetMaintenanceWindowTask: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ssm_GetOpsItem(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ssm.NewFromConfig(awsCfg)
	var in ssm.GetOpsItemInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.GetOpsItem(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call GetOpsItem: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ssm_GetOpsMetadata(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ssm.NewFromConfig(awsCfg)
	var in ssm.GetOpsMetadataInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.GetOpsMetadata(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call GetOpsMetadata: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ssm_GetOpsSummary(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ssm.NewFromConfig(awsCfg)
	var in ssm.GetOpsSummaryInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.GetOpsSummary(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call GetOpsSummary: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ssm_GetParameter(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ssm.NewFromConfig(awsCfg)
	var in ssm.GetParameterInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.GetParameter(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call GetParameter: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ssm_GetParameterHistory(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ssm.NewFromConfig(awsCfg)
	var in ssm.GetParameterHistoryInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.GetParameterHistory(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call GetParameterHistory: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ssm_GetParameters(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ssm.NewFromConfig(awsCfg)
	var in ssm.GetParametersInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.GetParameters(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call GetParameters: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ssm_GetParametersByPath(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ssm.NewFromConfig(awsCfg)
	var in ssm.GetParametersByPathInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.GetParametersByPath(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call GetParametersByPath: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ssm_GetPatchBaseline(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ssm.NewFromConfig(awsCfg)
	var in ssm.GetPatchBaselineInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.GetPatchBaseline(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call GetPatchBaseline: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ssm_GetPatchBaselineForPatchGroup(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ssm.NewFromConfig(awsCfg)
	var in ssm.GetPatchBaselineForPatchGroupInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.GetPatchBaselineForPatchGroup(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call GetPatchBaselineForPatchGroup: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ssm_GetResourcePolicies(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ssm.NewFromConfig(awsCfg)
	var in ssm.GetResourcePoliciesInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.GetResourcePolicies(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call GetResourcePolicies: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ssm_GetServiceSetting(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ssm.NewFromConfig(awsCfg)
	var in ssm.GetServiceSettingInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.GetServiceSetting(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call GetServiceSetting: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ssm_LabelParameterVersion(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ssm.NewFromConfig(awsCfg)
	var in ssm.LabelParameterVersionInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.LabelParameterVersion(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call LabelParameterVersion: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ssm_ListAssociationVersions(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ssm.NewFromConfig(awsCfg)
	var in ssm.ListAssociationVersionsInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.ListAssociationVersions(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call ListAssociationVersions: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ssm_ListAssociations(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ssm.NewFromConfig(awsCfg)
	var in ssm.ListAssociationsInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.ListAssociations(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call ListAssociations: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ssm_ListCommandInvocations(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ssm.NewFromConfig(awsCfg)
	var in ssm.ListCommandInvocationsInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.ListCommandInvocations(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call ListCommandInvocations: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ssm_ListCommands(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ssm.NewFromConfig(awsCfg)
	var in ssm.ListCommandsInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.ListCommands(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call ListCommands: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ssm_ListComplianceItems(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ssm.NewFromConfig(awsCfg)
	var in ssm.ListComplianceItemsInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.ListComplianceItems(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call ListComplianceItems: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ssm_ListComplianceSummaries(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ssm.NewFromConfig(awsCfg)
	var in ssm.ListComplianceSummariesInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.ListComplianceSummaries(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call ListComplianceSummaries: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ssm_ListDocumentMetadataHistory(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ssm.NewFromConfig(awsCfg)
	var in ssm.ListDocumentMetadataHistoryInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.ListDocumentMetadataHistory(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call ListDocumentMetadataHistory: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ssm_ListDocumentVersions(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ssm.NewFromConfig(awsCfg)
	var in ssm.ListDocumentVersionsInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.ListDocumentVersions(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call ListDocumentVersions: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ssm_ListDocuments(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ssm.NewFromConfig(awsCfg)
	var in ssm.ListDocumentsInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.ListDocuments(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call ListDocuments: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ssm_ListInventoryEntries(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ssm.NewFromConfig(awsCfg)
	var in ssm.ListInventoryEntriesInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.ListInventoryEntries(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call ListInventoryEntries: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ssm_ListOpsItemEvents(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ssm.NewFromConfig(awsCfg)
	var in ssm.ListOpsItemEventsInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.ListOpsItemEvents(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call ListOpsItemEvents: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ssm_ListOpsItemRelatedItems(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ssm.NewFromConfig(awsCfg)
	var in ssm.ListOpsItemRelatedItemsInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.ListOpsItemRelatedItems(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call ListOpsItemRelatedItems: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ssm_ListOpsMetadata(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ssm.NewFromConfig(awsCfg)
	var in ssm.ListOpsMetadataInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.ListOpsMetadata(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call ListOpsMetadata: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ssm_ListResourceComplianceSummaries(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ssm.NewFromConfig(awsCfg)
	var in ssm.ListResourceComplianceSummariesInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.ListResourceComplianceSummaries(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call ListResourceComplianceSummaries: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ssm_ListResourceDataSync(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ssm.NewFromConfig(awsCfg)
	var in ssm.ListResourceDataSyncInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.ListResourceDataSync(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call ListResourceDataSync: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ssm_ListTagsForResource(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ssm.NewFromConfig(awsCfg)
	var in ssm.ListTagsForResourceInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.ListTagsForResource(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call ListTagsForResource: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ssm_ModifyDocumentPermission(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ssm.NewFromConfig(awsCfg)
	var in ssm.ModifyDocumentPermissionInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.ModifyDocumentPermission(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call ModifyDocumentPermission: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ssm_PutComplianceItems(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ssm.NewFromConfig(awsCfg)
	var in ssm.PutComplianceItemsInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.PutComplianceItems(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call PutComplianceItems: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ssm_PutInventory(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ssm.NewFromConfig(awsCfg)
	var in ssm.PutInventoryInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.PutInventory(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call PutInventory: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ssm_PutParameter(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ssm.NewFromConfig(awsCfg)
	var in ssm.PutParameterInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.PutParameter(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call PutParameter: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ssm_PutResourcePolicy(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ssm.NewFromConfig(awsCfg)
	var in ssm.PutResourcePolicyInput
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

func ssm_RegisterDefaultPatchBaseline(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ssm.NewFromConfig(awsCfg)
	var in ssm.RegisterDefaultPatchBaselineInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.RegisterDefaultPatchBaseline(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call RegisterDefaultPatchBaseline: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ssm_RegisterPatchBaselineForPatchGroup(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ssm.NewFromConfig(awsCfg)
	var in ssm.RegisterPatchBaselineForPatchGroupInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.RegisterPatchBaselineForPatchGroup(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call RegisterPatchBaselineForPatchGroup: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ssm_RegisterTargetWithMaintenanceWindow(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ssm.NewFromConfig(awsCfg)
	var in ssm.RegisterTargetWithMaintenanceWindowInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.RegisterTargetWithMaintenanceWindow(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call RegisterTargetWithMaintenanceWindow: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ssm_RegisterTaskWithMaintenanceWindow(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ssm.NewFromConfig(awsCfg)
	var in ssm.RegisterTaskWithMaintenanceWindowInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.RegisterTaskWithMaintenanceWindow(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call RegisterTaskWithMaintenanceWindow: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ssm_RemoveTagsFromResource(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ssm.NewFromConfig(awsCfg)
	var in ssm.RemoveTagsFromResourceInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.RemoveTagsFromResource(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call RemoveTagsFromResource: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ssm_ResetServiceSetting(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ssm.NewFromConfig(awsCfg)
	var in ssm.ResetServiceSettingInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.ResetServiceSetting(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call ResetServiceSetting: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ssm_ResumeSession(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ssm.NewFromConfig(awsCfg)
	var in ssm.ResumeSessionInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.ResumeSession(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call ResumeSession: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ssm_SendAutomationSignal(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ssm.NewFromConfig(awsCfg)
	var in ssm.SendAutomationSignalInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.SendAutomationSignal(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call SendAutomationSignal: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ssm_SendCommand(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ssm.NewFromConfig(awsCfg)
	var in ssm.SendCommandInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.SendCommand(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call SendCommand: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ssm_StartAssociationsOnce(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ssm.NewFromConfig(awsCfg)
	var in ssm.StartAssociationsOnceInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.StartAssociationsOnce(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call StartAssociationsOnce: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ssm_StartAutomationExecution(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ssm.NewFromConfig(awsCfg)
	var in ssm.StartAutomationExecutionInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.StartAutomationExecution(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call StartAutomationExecution: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ssm_StartChangeRequestExecution(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ssm.NewFromConfig(awsCfg)
	var in ssm.StartChangeRequestExecutionInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.StartChangeRequestExecution(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call StartChangeRequestExecution: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ssm_StartSession(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ssm.NewFromConfig(awsCfg)
	var in ssm.StartSessionInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.StartSession(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call StartSession: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ssm_StopAutomationExecution(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ssm.NewFromConfig(awsCfg)
	var in ssm.StopAutomationExecutionInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.StopAutomationExecution(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call StopAutomationExecution: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ssm_TerminateSession(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ssm.NewFromConfig(awsCfg)
	var in ssm.TerminateSessionInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.TerminateSession(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call TerminateSession: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ssm_UnlabelParameterVersion(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ssm.NewFromConfig(awsCfg)
	var in ssm.UnlabelParameterVersionInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.UnlabelParameterVersion(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call UnlabelParameterVersion: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ssm_UpdateAssociation(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ssm.NewFromConfig(awsCfg)
	var in ssm.UpdateAssociationInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.UpdateAssociation(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call UpdateAssociation: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ssm_UpdateAssociationStatus(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ssm.NewFromConfig(awsCfg)
	var in ssm.UpdateAssociationStatusInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.UpdateAssociationStatus(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call UpdateAssociationStatus: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ssm_UpdateDocument(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ssm.NewFromConfig(awsCfg)
	var in ssm.UpdateDocumentInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.UpdateDocument(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call UpdateDocument: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ssm_UpdateDocumentDefaultVersion(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ssm.NewFromConfig(awsCfg)
	var in ssm.UpdateDocumentDefaultVersionInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.UpdateDocumentDefaultVersion(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call UpdateDocumentDefaultVersion: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ssm_UpdateDocumentMetadata(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ssm.NewFromConfig(awsCfg)
	var in ssm.UpdateDocumentMetadataInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.UpdateDocumentMetadata(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call UpdateDocumentMetadata: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ssm_UpdateMaintenanceWindow(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ssm.NewFromConfig(awsCfg)
	var in ssm.UpdateMaintenanceWindowInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.UpdateMaintenanceWindow(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call UpdateMaintenanceWindow: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ssm_UpdateMaintenanceWindowTarget(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ssm.NewFromConfig(awsCfg)
	var in ssm.UpdateMaintenanceWindowTargetInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.UpdateMaintenanceWindowTarget(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call UpdateMaintenanceWindowTarget: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ssm_UpdateMaintenanceWindowTask(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ssm.NewFromConfig(awsCfg)
	var in ssm.UpdateMaintenanceWindowTaskInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.UpdateMaintenanceWindowTask(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call UpdateMaintenanceWindowTask: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ssm_UpdateManagedInstanceRole(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ssm.NewFromConfig(awsCfg)
	var in ssm.UpdateManagedInstanceRoleInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.UpdateManagedInstanceRole(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call UpdateManagedInstanceRole: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ssm_UpdateOpsItem(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ssm.NewFromConfig(awsCfg)
	var in ssm.UpdateOpsItemInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.UpdateOpsItem(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call UpdateOpsItem: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ssm_UpdateOpsMetadata(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ssm.NewFromConfig(awsCfg)
	var in ssm.UpdateOpsMetadataInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.UpdateOpsMetadata(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call UpdateOpsMetadata: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ssm_UpdatePatchBaseline(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ssm.NewFromConfig(awsCfg)
	var in ssm.UpdatePatchBaselineInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.UpdatePatchBaseline(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call UpdatePatchBaseline: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ssm_UpdateResourceDataSync(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ssm.NewFromConfig(awsCfg)
	var in ssm.UpdateResourceDataSyncInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.UpdateResourceDataSync(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call UpdateResourceDataSync: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ssm_UpdateServiceSetting(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ssm.NewFromConfig(awsCfg)
	var in ssm.UpdateServiceSettingInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.UpdateServiceSetting(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call UpdateServiceSetting: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func init() {
	clientMethods["ssm_AddTagsToResource"] = ssm_AddTagsToResource
	clientMethods["ssm_AssociateOpsItemRelatedItem"] = ssm_AssociateOpsItemRelatedItem
	clientMethods["ssm_CancelCommand"] = ssm_CancelCommand
	clientMethods["ssm_CancelMaintenanceWindowExecution"] = ssm_CancelMaintenanceWindowExecution
	clientMethods["ssm_CreateActivation"] = ssm_CreateActivation
	clientMethods["ssm_CreateAssociation"] = ssm_CreateAssociation
	clientMethods["ssm_CreateAssociationBatch"] = ssm_CreateAssociationBatch
	clientMethods["ssm_CreateDocument"] = ssm_CreateDocument
	clientMethods["ssm_CreateMaintenanceWindow"] = ssm_CreateMaintenanceWindow
	clientMethods["ssm_CreateOpsItem"] = ssm_CreateOpsItem
	clientMethods["ssm_CreateOpsMetadata"] = ssm_CreateOpsMetadata
	clientMethods["ssm_CreatePatchBaseline"] = ssm_CreatePatchBaseline
	clientMethods["ssm_CreateResourceDataSync"] = ssm_CreateResourceDataSync
	clientMethods["ssm_DeleteActivation"] = ssm_DeleteActivation
	clientMethods["ssm_DeleteAssociation"] = ssm_DeleteAssociation
	clientMethods["ssm_DeleteDocument"] = ssm_DeleteDocument
	clientMethods["ssm_DeleteInventory"] = ssm_DeleteInventory
	clientMethods["ssm_DeleteMaintenanceWindow"] = ssm_DeleteMaintenanceWindow
	clientMethods["ssm_DeleteOpsItem"] = ssm_DeleteOpsItem
	clientMethods["ssm_DeleteOpsMetadata"] = ssm_DeleteOpsMetadata
	clientMethods["ssm_DeleteParameter"] = ssm_DeleteParameter
	clientMethods["ssm_DeleteParameters"] = ssm_DeleteParameters
	clientMethods["ssm_DeletePatchBaseline"] = ssm_DeletePatchBaseline
	clientMethods["ssm_DeleteResourceDataSync"] = ssm_DeleteResourceDataSync
	clientMethods["ssm_DeleteResourcePolicy"] = ssm_DeleteResourcePolicy
	clientMethods["ssm_DeregisterManagedInstance"] = ssm_DeregisterManagedInstance
	clientMethods["ssm_DeregisterPatchBaselineForPatchGroup"] = ssm_DeregisterPatchBaselineForPatchGroup
	clientMethods["ssm_DeregisterTargetFromMaintenanceWindow"] = ssm_DeregisterTargetFromMaintenanceWindow
	clientMethods["ssm_DeregisterTaskFromMaintenanceWindow"] = ssm_DeregisterTaskFromMaintenanceWindow
	clientMethods["ssm_DescribeActivations"] = ssm_DescribeActivations
	clientMethods["ssm_DescribeAssociation"] = ssm_DescribeAssociation
	clientMethods["ssm_DescribeAssociationExecutionTargets"] = ssm_DescribeAssociationExecutionTargets
	clientMethods["ssm_DescribeAssociationExecutions"] = ssm_DescribeAssociationExecutions
	clientMethods["ssm_DescribeAutomationExecutions"] = ssm_DescribeAutomationExecutions
	clientMethods["ssm_DescribeAutomationStepExecutions"] = ssm_DescribeAutomationStepExecutions
	clientMethods["ssm_DescribeAvailablePatches"] = ssm_DescribeAvailablePatches
	clientMethods["ssm_DescribeDocument"] = ssm_DescribeDocument
	clientMethods["ssm_DescribeDocumentPermission"] = ssm_DescribeDocumentPermission
	clientMethods["ssm_DescribeEffectiveInstanceAssociations"] = ssm_DescribeEffectiveInstanceAssociations
	clientMethods["ssm_DescribeEffectivePatchesForPatchBaseline"] = ssm_DescribeEffectivePatchesForPatchBaseline
	clientMethods["ssm_DescribeInstanceAssociationsStatus"] = ssm_DescribeInstanceAssociationsStatus
	clientMethods["ssm_DescribeInstanceInformation"] = ssm_DescribeInstanceInformation
	clientMethods["ssm_DescribeInstancePatchStates"] = ssm_DescribeInstancePatchStates
	clientMethods["ssm_DescribeInstancePatchStatesForPatchGroup"] = ssm_DescribeInstancePatchStatesForPatchGroup
	clientMethods["ssm_DescribeInstancePatches"] = ssm_DescribeInstancePatches
	clientMethods["ssm_DescribeInstanceProperties"] = ssm_DescribeInstanceProperties
	clientMethods["ssm_DescribeInventoryDeletions"] = ssm_DescribeInventoryDeletions
	clientMethods["ssm_DescribeMaintenanceWindowExecutionTaskInvocations"] = ssm_DescribeMaintenanceWindowExecutionTaskInvocations
	clientMethods["ssm_DescribeMaintenanceWindowExecutionTasks"] = ssm_DescribeMaintenanceWindowExecutionTasks
	clientMethods["ssm_DescribeMaintenanceWindowExecutions"] = ssm_DescribeMaintenanceWindowExecutions
	clientMethods["ssm_DescribeMaintenanceWindowSchedule"] = ssm_DescribeMaintenanceWindowSchedule
	clientMethods["ssm_DescribeMaintenanceWindowTargets"] = ssm_DescribeMaintenanceWindowTargets
	clientMethods["ssm_DescribeMaintenanceWindowTasks"] = ssm_DescribeMaintenanceWindowTasks
	clientMethods["ssm_DescribeMaintenanceWindows"] = ssm_DescribeMaintenanceWindows
	clientMethods["ssm_DescribeMaintenanceWindowsForTarget"] = ssm_DescribeMaintenanceWindowsForTarget
	clientMethods["ssm_DescribeOpsItems"] = ssm_DescribeOpsItems
	clientMethods["ssm_DescribeParameters"] = ssm_DescribeParameters
	clientMethods["ssm_DescribePatchBaselines"] = ssm_DescribePatchBaselines
	clientMethods["ssm_DescribePatchGroupState"] = ssm_DescribePatchGroupState
	clientMethods["ssm_DescribePatchGroups"] = ssm_DescribePatchGroups
	clientMethods["ssm_DescribePatchProperties"] = ssm_DescribePatchProperties
	clientMethods["ssm_DescribeSessions"] = ssm_DescribeSessions
	clientMethods["ssm_DisassociateOpsItemRelatedItem"] = ssm_DisassociateOpsItemRelatedItem
	clientMethods["ssm_GetAutomationExecution"] = ssm_GetAutomationExecution
	clientMethods["ssm_GetCalendarState"] = ssm_GetCalendarState
	clientMethods["ssm_GetCommandInvocation"] = ssm_GetCommandInvocation
	clientMethods["ssm_GetConnectionStatus"] = ssm_GetConnectionStatus
	clientMethods["ssm_GetDefaultPatchBaseline"] = ssm_GetDefaultPatchBaseline
	clientMethods["ssm_GetDeployablePatchSnapshotForInstance"] = ssm_GetDeployablePatchSnapshotForInstance
	clientMethods["ssm_GetDocument"] = ssm_GetDocument
	clientMethods["ssm_GetInventory"] = ssm_GetInventory
	clientMethods["ssm_GetInventorySchema"] = ssm_GetInventorySchema
	clientMethods["ssm_GetMaintenanceWindow"] = ssm_GetMaintenanceWindow
	clientMethods["ssm_GetMaintenanceWindowExecution"] = ssm_GetMaintenanceWindowExecution
	clientMethods["ssm_GetMaintenanceWindowExecutionTask"] = ssm_GetMaintenanceWindowExecutionTask
	clientMethods["ssm_GetMaintenanceWindowExecutionTaskInvocation"] = ssm_GetMaintenanceWindowExecutionTaskInvocation
	clientMethods["ssm_GetMaintenanceWindowTask"] = ssm_GetMaintenanceWindowTask
	clientMethods["ssm_GetOpsItem"] = ssm_GetOpsItem
	clientMethods["ssm_GetOpsMetadata"] = ssm_GetOpsMetadata
	clientMethods["ssm_GetOpsSummary"] = ssm_GetOpsSummary
	clientMethods["ssm_GetParameter"] = ssm_GetParameter
	clientMethods["ssm_GetParameterHistory"] = ssm_GetParameterHistory
	clientMethods["ssm_GetParameters"] = ssm_GetParameters
	clientMethods["ssm_GetParametersByPath"] = ssm_GetParametersByPath
	clientMethods["ssm_GetPatchBaseline"] = ssm_GetPatchBaseline
	clientMethods["ssm_GetPatchBaselineForPatchGroup"] = ssm_GetPatchBaselineForPatchGroup
	clientMethods["ssm_GetResourcePolicies"] = ssm_GetResourcePolicies
	clientMethods["ssm_GetServiceSetting"] = ssm_GetServiceSetting
	clientMethods["ssm_LabelParameterVersion"] = ssm_LabelParameterVersion
	clientMethods["ssm_ListAssociationVersions"] = ssm_ListAssociationVersions
	clientMethods["ssm_ListAssociations"] = ssm_ListAssociations
	clientMethods["ssm_ListCommandInvocations"] = ssm_ListCommandInvocations
	clientMethods["ssm_ListCommands"] = ssm_ListCommands
	clientMethods["ssm_ListComplianceItems"] = ssm_ListComplianceItems
	clientMethods["ssm_ListComplianceSummaries"] = ssm_ListComplianceSummaries
	clientMethods["ssm_ListDocumentMetadataHistory"] = ssm_ListDocumentMetadataHistory
	clientMethods["ssm_ListDocumentVersions"] = ssm_ListDocumentVersions
	clientMethods["ssm_ListDocuments"] = ssm_ListDocuments
	clientMethods["ssm_ListInventoryEntries"] = ssm_ListInventoryEntries
	clientMethods["ssm_ListOpsItemEvents"] = ssm_ListOpsItemEvents
	clientMethods["ssm_ListOpsItemRelatedItems"] = ssm_ListOpsItemRelatedItems
	clientMethods["ssm_ListOpsMetadata"] = ssm_ListOpsMetadata
	clientMethods["ssm_ListResourceComplianceSummaries"] = ssm_ListResourceComplianceSummaries
	clientMethods["ssm_ListResourceDataSync"] = ssm_ListResourceDataSync
	clientMethods["ssm_ListTagsForResource"] = ssm_ListTagsForResource
	clientMethods["ssm_ModifyDocumentPermission"] = ssm_ModifyDocumentPermission
	clientMethods["ssm_PutComplianceItems"] = ssm_PutComplianceItems
	clientMethods["ssm_PutInventory"] = ssm_PutInventory
	clientMethods["ssm_PutParameter"] = ssm_PutParameter
	clientMethods["ssm_PutResourcePolicy"] = ssm_PutResourcePolicy
	clientMethods["ssm_RegisterDefaultPatchBaseline"] = ssm_RegisterDefaultPatchBaseline
	clientMethods["ssm_RegisterPatchBaselineForPatchGroup"] = ssm_RegisterPatchBaselineForPatchGroup
	clientMethods["ssm_RegisterTargetWithMaintenanceWindow"] = ssm_RegisterTargetWithMaintenanceWindow
	clientMethods["ssm_RegisterTaskWithMaintenanceWindow"] = ssm_RegisterTaskWithMaintenanceWindow
	clientMethods["ssm_RemoveTagsFromResource"] = ssm_RemoveTagsFromResource
	clientMethods["ssm_ResetServiceSetting"] = ssm_ResetServiceSetting
	clientMethods["ssm_ResumeSession"] = ssm_ResumeSession
	clientMethods["ssm_SendAutomationSignal"] = ssm_SendAutomationSignal
	clientMethods["ssm_SendCommand"] = ssm_SendCommand
	clientMethods["ssm_StartAssociationsOnce"] = ssm_StartAssociationsOnce
	clientMethods["ssm_StartAutomationExecution"] = ssm_StartAutomationExecution
	clientMethods["ssm_StartChangeRequestExecution"] = ssm_StartChangeRequestExecution
	clientMethods["ssm_StartSession"] = ssm_StartSession
	clientMethods["ssm_StopAutomationExecution"] = ssm_StopAutomationExecution
	clientMethods["ssm_TerminateSession"] = ssm_TerminateSession
	clientMethods["ssm_UnlabelParameterVersion"] = ssm_UnlabelParameterVersion
	clientMethods["ssm_UpdateAssociation"] = ssm_UpdateAssociation
	clientMethods["ssm_UpdateAssociationStatus"] = ssm_UpdateAssociationStatus
	clientMethods["ssm_UpdateDocument"] = ssm_UpdateDocument
	clientMethods["ssm_UpdateDocumentDefaultVersion"] = ssm_UpdateDocumentDefaultVersion
	clientMethods["ssm_UpdateDocumentMetadata"] = ssm_UpdateDocumentMetadata
	clientMethods["ssm_UpdateMaintenanceWindow"] = ssm_UpdateMaintenanceWindow
	clientMethods["ssm_UpdateMaintenanceWindowTarget"] = ssm_UpdateMaintenanceWindowTarget
	clientMethods["ssm_UpdateMaintenanceWindowTask"] = ssm_UpdateMaintenanceWindowTask
	clientMethods["ssm_UpdateManagedInstanceRole"] = ssm_UpdateManagedInstanceRole
	clientMethods["ssm_UpdateOpsItem"] = ssm_UpdateOpsItem
	clientMethods["ssm_UpdateOpsMetadata"] = ssm_UpdateOpsMetadata
	clientMethods["ssm_UpdatePatchBaseline"] = ssm_UpdatePatchBaseline
	clientMethods["ssm_UpdateResourceDataSync"] = ssm_UpdateResourceDataSync
	clientMethods["ssm_UpdateServiceSetting"] = ssm_UpdateServiceSetting
}
