package sdkclient

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
)

func ecs_CreateCapacityProvider(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ecs.NewFromConfig(awsCfg)
	var in ecs.CreateCapacityProviderInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.CreateCapacityProvider(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call CreateCapacityProvider: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ecs_CreateCluster(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ecs.NewFromConfig(awsCfg)
	var in ecs.CreateClusterInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.CreateCluster(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call CreateCluster: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ecs_CreateService(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ecs.NewFromConfig(awsCfg)
	var in ecs.CreateServiceInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.CreateService(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call CreateService: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ecs_CreateTaskSet(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ecs.NewFromConfig(awsCfg)
	var in ecs.CreateTaskSetInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.CreateTaskSet(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call CreateTaskSet: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ecs_DeleteAccountSetting(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ecs.NewFromConfig(awsCfg)
	var in ecs.DeleteAccountSettingInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.DeleteAccountSetting(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call DeleteAccountSetting: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ecs_DeleteAttributes(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ecs.NewFromConfig(awsCfg)
	var in ecs.DeleteAttributesInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.DeleteAttributes(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call DeleteAttributes: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ecs_DeleteCapacityProvider(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ecs.NewFromConfig(awsCfg)
	var in ecs.DeleteCapacityProviderInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.DeleteCapacityProvider(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call DeleteCapacityProvider: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ecs_DeleteCluster(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ecs.NewFromConfig(awsCfg)
	var in ecs.DeleteClusterInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.DeleteCluster(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call DeleteCluster: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ecs_DeleteService(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ecs.NewFromConfig(awsCfg)
	var in ecs.DeleteServiceInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.DeleteService(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call DeleteService: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ecs_DeleteTaskDefinitions(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ecs.NewFromConfig(awsCfg)
	var in ecs.DeleteTaskDefinitionsInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.DeleteTaskDefinitions(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call DeleteTaskDefinitions: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ecs_DeleteTaskSet(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ecs.NewFromConfig(awsCfg)
	var in ecs.DeleteTaskSetInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.DeleteTaskSet(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call DeleteTaskSet: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ecs_DeregisterContainerInstance(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ecs.NewFromConfig(awsCfg)
	var in ecs.DeregisterContainerInstanceInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.DeregisterContainerInstance(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call DeregisterContainerInstance: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ecs_DeregisterTaskDefinition(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ecs.NewFromConfig(awsCfg)
	var in ecs.DeregisterTaskDefinitionInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.DeregisterTaskDefinition(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call DeregisterTaskDefinition: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ecs_DescribeCapacityProviders(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ecs.NewFromConfig(awsCfg)
	var in ecs.DescribeCapacityProvidersInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.DescribeCapacityProviders(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call DescribeCapacityProviders: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ecs_DescribeClusters(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ecs.NewFromConfig(awsCfg)
	var in ecs.DescribeClustersInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.DescribeClusters(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call DescribeClusters: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ecs_DescribeContainerInstances(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ecs.NewFromConfig(awsCfg)
	var in ecs.DescribeContainerInstancesInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.DescribeContainerInstances(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call DescribeContainerInstances: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ecs_DescribeServices(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ecs.NewFromConfig(awsCfg)
	var in ecs.DescribeServicesInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.DescribeServices(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call DescribeServices: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ecs_DescribeTaskDefinition(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ecs.NewFromConfig(awsCfg)
	var in ecs.DescribeTaskDefinitionInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.DescribeTaskDefinition(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call DescribeTaskDefinition: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ecs_DescribeTaskSets(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ecs.NewFromConfig(awsCfg)
	var in ecs.DescribeTaskSetsInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.DescribeTaskSets(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call DescribeTaskSets: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ecs_DescribeTasks(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ecs.NewFromConfig(awsCfg)
	var in ecs.DescribeTasksInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.DescribeTasks(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call DescribeTasks: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ecs_DiscoverPollEndpoint(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ecs.NewFromConfig(awsCfg)
	var in ecs.DiscoverPollEndpointInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.DiscoverPollEndpoint(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call DiscoverPollEndpoint: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ecs_ExecuteCommand(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ecs.NewFromConfig(awsCfg)
	var in ecs.ExecuteCommandInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.ExecuteCommand(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call ExecuteCommand: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ecs_GetTaskProtection(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ecs.NewFromConfig(awsCfg)
	var in ecs.GetTaskProtectionInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.GetTaskProtection(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call GetTaskProtection: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ecs_ListAccountSettings(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ecs.NewFromConfig(awsCfg)
	var in ecs.ListAccountSettingsInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.ListAccountSettings(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call ListAccountSettings: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ecs_ListAttributes(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ecs.NewFromConfig(awsCfg)
	var in ecs.ListAttributesInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.ListAttributes(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call ListAttributes: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ecs_ListClusters(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ecs.NewFromConfig(awsCfg)
	var in ecs.ListClustersInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.ListClusters(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call ListClusters: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ecs_ListContainerInstances(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ecs.NewFromConfig(awsCfg)
	var in ecs.ListContainerInstancesInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.ListContainerInstances(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call ListContainerInstances: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ecs_ListServices(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ecs.NewFromConfig(awsCfg)
	var in ecs.ListServicesInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.ListServices(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call ListServices: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ecs_ListServicesByNamespace(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ecs.NewFromConfig(awsCfg)
	var in ecs.ListServicesByNamespaceInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.ListServicesByNamespace(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call ListServicesByNamespace: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ecs_ListTagsForResource(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ecs.NewFromConfig(awsCfg)
	var in ecs.ListTagsForResourceInput
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

func ecs_ListTaskDefinitionFamilies(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ecs.NewFromConfig(awsCfg)
	var in ecs.ListTaskDefinitionFamiliesInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.ListTaskDefinitionFamilies(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call ListTaskDefinitionFamilies: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ecs_ListTaskDefinitions(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ecs.NewFromConfig(awsCfg)
	var in ecs.ListTaskDefinitionsInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.ListTaskDefinitions(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call ListTaskDefinitions: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ecs_ListTasks(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ecs.NewFromConfig(awsCfg)
	var in ecs.ListTasksInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.ListTasks(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call ListTasks: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ecs_PutAccountSetting(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ecs.NewFromConfig(awsCfg)
	var in ecs.PutAccountSettingInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.PutAccountSetting(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call PutAccountSetting: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ecs_PutAccountSettingDefault(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ecs.NewFromConfig(awsCfg)
	var in ecs.PutAccountSettingDefaultInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.PutAccountSettingDefault(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call PutAccountSettingDefault: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ecs_PutAttributes(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ecs.NewFromConfig(awsCfg)
	var in ecs.PutAttributesInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.PutAttributes(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call PutAttributes: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ecs_PutClusterCapacityProviders(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ecs.NewFromConfig(awsCfg)
	var in ecs.PutClusterCapacityProvidersInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.PutClusterCapacityProviders(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call PutClusterCapacityProviders: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ecs_RegisterContainerInstance(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ecs.NewFromConfig(awsCfg)
	var in ecs.RegisterContainerInstanceInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.RegisterContainerInstance(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call RegisterContainerInstance: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ecs_RegisterTaskDefinition(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ecs.NewFromConfig(awsCfg)
	var in ecs.RegisterTaskDefinitionInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.RegisterTaskDefinition(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call RegisterTaskDefinition: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ecs_RunTask(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ecs.NewFromConfig(awsCfg)
	var in ecs.RunTaskInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.RunTask(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call RunTask: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ecs_StartTask(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ecs.NewFromConfig(awsCfg)
	var in ecs.StartTaskInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.StartTask(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call StartTask: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ecs_StopTask(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ecs.NewFromConfig(awsCfg)
	var in ecs.StopTaskInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.StopTask(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call StopTask: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ecs_SubmitAttachmentStateChanges(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ecs.NewFromConfig(awsCfg)
	var in ecs.SubmitAttachmentStateChangesInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.SubmitAttachmentStateChanges(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call SubmitAttachmentStateChanges: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ecs_SubmitContainerStateChange(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ecs.NewFromConfig(awsCfg)
	var in ecs.SubmitContainerStateChangeInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.SubmitContainerStateChange(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call SubmitContainerStateChange: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ecs_SubmitTaskStateChange(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ecs.NewFromConfig(awsCfg)
	var in ecs.SubmitTaskStateChangeInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.SubmitTaskStateChange(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call SubmitTaskStateChange: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ecs_TagResource(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ecs.NewFromConfig(awsCfg)
	var in ecs.TagResourceInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.TagResource(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call TagResource: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ecs_UntagResource(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ecs.NewFromConfig(awsCfg)
	var in ecs.UntagResourceInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.UntagResource(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call UntagResource: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ecs_UpdateCapacityProvider(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ecs.NewFromConfig(awsCfg)
	var in ecs.UpdateCapacityProviderInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.UpdateCapacityProvider(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call UpdateCapacityProvider: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ecs_UpdateCluster(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ecs.NewFromConfig(awsCfg)
	var in ecs.UpdateClusterInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.UpdateCluster(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call UpdateCluster: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ecs_UpdateClusterSettings(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ecs.NewFromConfig(awsCfg)
	var in ecs.UpdateClusterSettingsInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.UpdateClusterSettings(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call UpdateClusterSettings: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ecs_UpdateContainerAgent(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ecs.NewFromConfig(awsCfg)
	var in ecs.UpdateContainerAgentInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.UpdateContainerAgent(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call UpdateContainerAgent: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ecs_UpdateContainerInstancesState(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ecs.NewFromConfig(awsCfg)
	var in ecs.UpdateContainerInstancesStateInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.UpdateContainerInstancesState(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call UpdateContainerInstancesState: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ecs_UpdateService(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ecs.NewFromConfig(awsCfg)
	var in ecs.UpdateServiceInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.UpdateService(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call UpdateService: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ecs_UpdateServicePrimaryTaskSet(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ecs.NewFromConfig(awsCfg)
	var in ecs.UpdateServicePrimaryTaskSetInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.UpdateServicePrimaryTaskSet(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call UpdateServicePrimaryTaskSet: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ecs_UpdateTaskProtection(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ecs.NewFromConfig(awsCfg)
	var in ecs.UpdateTaskProtectionInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.UpdateTaskProtection(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call UpdateTaskProtection: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func ecs_UpdateTaskSet(ctx context.Context, b json.RawMessage) (json.RawMessage, error) {
	svc := ecs.NewFromConfig(awsCfg)
	var in ecs.UpdateTaskSetInput
	if err := json.Unmarshal(b, &in); err != nil {
		return nil, fmt.Errorf("failed to unmarshal request: %w", err)
	}
	if out, err := svc.UpdateTaskSet(ctx, &in); err != nil {
		return nil, fmt.Errorf("failed to call UpdateTaskSet: %w", err)
	} else {
		o, err := json.Marshal(out)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}
		return o, nil
	}
}

func init() {
	clientMethods["ecs_CreateCapacityProvider"] = ecs_CreateCapacityProvider
	clientMethods["ecs_CreateCluster"] = ecs_CreateCluster
	clientMethods["ecs_CreateService"] = ecs_CreateService
	clientMethods["ecs_CreateTaskSet"] = ecs_CreateTaskSet
	clientMethods["ecs_DeleteAccountSetting"] = ecs_DeleteAccountSetting
	clientMethods["ecs_DeleteAttributes"] = ecs_DeleteAttributes
	clientMethods["ecs_DeleteCapacityProvider"] = ecs_DeleteCapacityProvider
	clientMethods["ecs_DeleteCluster"] = ecs_DeleteCluster
	clientMethods["ecs_DeleteService"] = ecs_DeleteService
	clientMethods["ecs_DeleteTaskDefinitions"] = ecs_DeleteTaskDefinitions
	clientMethods["ecs_DeleteTaskSet"] = ecs_DeleteTaskSet
	clientMethods["ecs_DeregisterContainerInstance"] = ecs_DeregisterContainerInstance
	clientMethods["ecs_DeregisterTaskDefinition"] = ecs_DeregisterTaskDefinition
	clientMethods["ecs_DescribeCapacityProviders"] = ecs_DescribeCapacityProviders
	clientMethods["ecs_DescribeClusters"] = ecs_DescribeClusters
	clientMethods["ecs_DescribeContainerInstances"] = ecs_DescribeContainerInstances
	clientMethods["ecs_DescribeServices"] = ecs_DescribeServices
	clientMethods["ecs_DescribeTaskDefinition"] = ecs_DescribeTaskDefinition
	clientMethods["ecs_DescribeTaskSets"] = ecs_DescribeTaskSets
	clientMethods["ecs_DescribeTasks"] = ecs_DescribeTasks
	clientMethods["ecs_DiscoverPollEndpoint"] = ecs_DiscoverPollEndpoint
	clientMethods["ecs_ExecuteCommand"] = ecs_ExecuteCommand
	clientMethods["ecs_GetTaskProtection"] = ecs_GetTaskProtection
	clientMethods["ecs_ListAccountSettings"] = ecs_ListAccountSettings
	clientMethods["ecs_ListAttributes"] = ecs_ListAttributes
	clientMethods["ecs_ListClusters"] = ecs_ListClusters
	clientMethods["ecs_ListContainerInstances"] = ecs_ListContainerInstances
	clientMethods["ecs_ListServices"] = ecs_ListServices
	clientMethods["ecs_ListServicesByNamespace"] = ecs_ListServicesByNamespace
	clientMethods["ecs_ListTagsForResource"] = ecs_ListTagsForResource
	clientMethods["ecs_ListTaskDefinitionFamilies"] = ecs_ListTaskDefinitionFamilies
	clientMethods["ecs_ListTaskDefinitions"] = ecs_ListTaskDefinitions
	clientMethods["ecs_ListTasks"] = ecs_ListTasks
	clientMethods["ecs_PutAccountSetting"] = ecs_PutAccountSetting
	clientMethods["ecs_PutAccountSettingDefault"] = ecs_PutAccountSettingDefault
	clientMethods["ecs_PutAttributes"] = ecs_PutAttributes
	clientMethods["ecs_PutClusterCapacityProviders"] = ecs_PutClusterCapacityProviders
	clientMethods["ecs_RegisterContainerInstance"] = ecs_RegisterContainerInstance
	clientMethods["ecs_RegisterTaskDefinition"] = ecs_RegisterTaskDefinition
	clientMethods["ecs_RunTask"] = ecs_RunTask
	clientMethods["ecs_StartTask"] = ecs_StartTask
	clientMethods["ecs_StopTask"] = ecs_StopTask
	clientMethods["ecs_SubmitAttachmentStateChanges"] = ecs_SubmitAttachmentStateChanges
	clientMethods["ecs_SubmitContainerStateChange"] = ecs_SubmitContainerStateChange
	clientMethods["ecs_SubmitTaskStateChange"] = ecs_SubmitTaskStateChange
	clientMethods["ecs_TagResource"] = ecs_TagResource
	clientMethods["ecs_UntagResource"] = ecs_UntagResource
	clientMethods["ecs_UpdateCapacityProvider"] = ecs_UpdateCapacityProvider
	clientMethods["ecs_UpdateCluster"] = ecs_UpdateCluster
	clientMethods["ecs_UpdateClusterSettings"] = ecs_UpdateClusterSettings
	clientMethods["ecs_UpdateContainerAgent"] = ecs_UpdateContainerAgent
	clientMethods["ecs_UpdateContainerInstancesState"] = ecs_UpdateContainerInstancesState
	clientMethods["ecs_UpdateService"] = ecs_UpdateService
	clientMethods["ecs_UpdateServicePrimaryTaskSet"] = ecs_UpdateServicePrimaryTaskSet
	clientMethods["ecs_UpdateTaskProtection"] = ecs_UpdateTaskProtection
	clientMethods["ecs_UpdateTaskSet"] = ecs_UpdateTaskSet
}
