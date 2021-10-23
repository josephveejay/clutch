package ecs

import (
	"context"

	ecsv1 "github.com/lyft/clutch/backend/api/aws/ecs/v1"
	"github.com/lyft/clutch/backend/service/aws"
)

func newECSAPI(c aws.Client) ecsv1.EcsAPIServer {
	return &ecsAPI{
		client: c,
	}
}

type ecsAPI struct {
	client aws.Client
}

func (a *ecsAPI) GetECSClusters(ctx context.Context, req *ecsv1.GetEcsRequest) (*ecsv1.GetEcsResponse, error) {
	clusters, err := a.client.ListECSClusters(ctx, req.Region)
	if err != nil {
		return nil, err
	}
	return clusters, nil
}

func (a *ecsAPI) GetECSClusterServices(ctx context.Context, req *ecsv1.GetEcsSrvRequest) (*ecsv1.GetEcsSrvResponse, error) {
	services, err := a.client.ListECSServices(ctx, req.EcsCluster, req.Region)
	if err != nil {
		return nil, err
	}
	return services, nil
}

func (a *ecsAPI) GetECSClusterSrvTaskDef(ctx context.Context, req *ecsv1.GetEcsSrvTaskRequest) (*ecsv1.GetEcsSrvTaskResponse, error) {
	taskenvs, err := a.client.DescribeTaskDefs(ctx, req.EcsClusterservice, req.EcsCluster, req.Region)
	if err != nil {
		return nil, err
	}
	return taskenvs, nil
}
