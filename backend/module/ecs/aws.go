package ecs

import (
	"errors"

	"github.com/golang/protobuf/ptypes/any"
	"github.com/uber-go/tally"
	"go.uber.org/zap"

	ecsv1 "github.com/lyft/clutch/backend/api/aws/ecs/v1"
	"github.com/lyft/clutch/backend/module"
	"github.com/lyft/clutch/backend/service"
	"github.com/lyft/clutch/backend/service/aws"
)

const (
	Name = "clutch.module.ecs"
)

func New(*any.Any, *zap.Logger, tally.Scope) (module.Module, error) {
	awsClient, ok := service.Registry["clutch.service.aws"]
	if !ok {
		return nil, errors.New("could not find service")
	}

	c, ok := awsClient.(aws.Client)
	if !ok {
		return nil, errors.New("ecs: service was not the correct type")
	}

	mod := &mod{
		ecs: newECSAPI(c),
	}

	return mod, nil
}

type mod struct {
	ecs ecsv1.EcsAPIServer
}

func (m *mod) Register(r module.Registrar) error {
	ecsv1.RegisterEcsAPIServer(r.GRPCServer(), m.ecs)
	return r.RegisterJSONGateway(ecsv1.RegisterEcsAPIHandler)
}
