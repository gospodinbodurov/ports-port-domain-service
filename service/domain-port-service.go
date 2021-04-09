package service

import (
	"context"

	api "github.com/gospodinbodurov/ports-apis/port-domain-service/api"
	"github.com/gospodinbodurov/ports-port-domain-service/storage"

	"github.com/palantir/stacktrace"
)

type DomainPortService struct {
	Storage storage.Storage
	api.UnimplementedDomainPortServiceServer
}

func (dps *DomainPortService) PutPort(ctx context.Context, in *api.PutPortRequest) (*api.PutPortResponse, error) {
	err := dps.Storage.PutPort(in.Port)

	if err != nil {
		return &api.PutPortResponse{}, stacktrace.Propagate(err, "Failed to put %s", in.Port.PortKey)
	}

	return &api.PutPortResponse{}, nil
}

func (dps *DomainPortService) GetPort(ctx context.Context, in *api.GetPortRequest) (*api.GetPortResponse, error) {
	port, err := dps.Storage.GetPort(in.PortKey)

	if err != nil {
		return &api.GetPortResponse{}, stacktrace.Propagate(err, "Failed to get %s", in.PortKey)
	}

	return &api.GetPortResponse{
		Port: port,
	}, nil
}
