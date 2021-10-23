package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
	ecsv1 "github.com/lyft/clutch/backend/api/aws/ecs/v1"
)

func (c *client) ListECSClusters(ctx context.Context, region string) (*ecsv1.GetEcsResponse, error) {
	rc, err := c.getRegionalClient(region)
	if err != nil {
		fmt.Println(err)
	}
	reqParams := &ecs.ListClustersInput{}
	resp, err := rc.ecs.ListClusters(ctx, reqParams)
	if err != nil {
		fmt.Printf("ecs.ListClusters: %s", err.Error())
	}
	fmt.Println(resp.ClusterArns)
	//var ecscluster ecsv1.GetEcsResponse
	//ecscluster.Cluster = resp.ClusterArns
	return &ecsv1.GetEcsResponse{Cluster: resp.ClusterArns}, nil
}

func (c *client) ListECSServices(ctx context.Context, ecs_cluster string, region string) (*ecsv1.GetEcsSrvResponse, error) {
	rc, err := c.getRegionalClient(region)
	if err != nil {
		fmt.Println(err)
	}
	lsrvinput := &ecs.ListServicesInput{
		Cluster: aws.String(ecs_cluster),
	}
	resp, err := rc.ecs.ListServices(ctx, lsrvinput)
	if err != nil {
		fmt.Printf("ecs.DescribeClusters: %s", err.Error())
	}
	fmt.Println(resp.ServiceArns)
	return &ecsv1.GetEcsSrvResponse{Service: resp.ServiceArns}, nil
}

func (c *client) DescribeTaskDefs(ctx context.Context, ecs_clusterservice string, ecs_cluster string, region string) (*ecsv1.GetEcsSrvTaskResponse, error) {
	rc, err := c.getRegionalClient(region)
	if err != nil {
		fmt.Println(err)
	}
	dsrvinput := &ecs.DescribeServicesInput{
		Services: []string{*aws.String(ecs_clusterservice)},
		Cluster:  aws.String(ecs_cluster),
	}
	dsrvresp, err := rc.ecs.DescribeServices(ctx, dsrvinput)
	if err != nil {
		fmt.Printf("ecs.DescribeServices: %s", err.Error())
	}
	fmt.Println(*dsrvresp.Services[0].TaskDefinition)
	dtskinput := &ecs.DescribeTaskDefinitionInput{
		TaskDefinition: aws.String(*dsrvresp.Services[0].TaskDefinition),
	}
	dtskresp, err := rc.ecs.DescribeTaskDefinition(ctx, dtskinput)
	if err != nil {
		fmt.Printf("ecs.DescribeTasks: %s", err.Error())
	}
	fmt.Println(dtskresp.TaskDefinition.ContainerDefinitions[0].Environment[0])
	//taskenvslist := make([]map[string]string, len(dtskresp.TaskDefinition.ContainerDefinitions[0].Environment))
	taskenvs := make(map[string]string, len(dtskresp.TaskDefinition.ContainerDefinitions[0].Environment))

	for _, v := range dtskresp.TaskDefinition.ContainerDefinitions[0].Environment {
		//taskenvslist = append(taskenvslist, map[string]string{*v.Name: *v.Value})
		taskenvs[*v.Name] = *v.Value
	}
	fmt.Println(taskenvs)
	return &ecsv1.GetEcsSrvTaskResponse{Env: taskenvs}, nil
}
