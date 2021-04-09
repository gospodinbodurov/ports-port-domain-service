package client

import (
	"context"
	"time"

	"github.com/gospodinbodurov/ports-apis/port-domain-service/api"
	"github.com/palantir/stacktrace"
	"google.golang.org/grpc"
)

type DomainPortClient struct {
	conn *grpc.ClientConn
}

func (dpc *DomainPortClient) Init(address string) error {
	conn, err := grpc.Dial(address, grpc.WithInsecure())

	if err != nil {
		return stacktrace.Propagate(err, "Failed to dial %s", address)
	}

	dpc.conn = conn

	return nil
}

func (dpc *DomainPortClient) GetPort(key string) (*api.Port, error) {
	client := api.NewDomainPortServiceClient(dpc.conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	in := &api.GetPortRequest{
		PortKey: key,
	}

	response, err := client.GetPort(ctx, in)

	if err != nil {
		return nil, stacktrace.Propagate(err, "Failed to get port %s", key)
	}

	return response.Port, nil
}

func (dpc *DomainPortClient) PutPort(port *api.Port) error {
	client := api.NewDomainPortServiceClient(dpc.conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	in := &api.PutPortRequest{
		Port: port,
	}

	_, err := client.PutPort(ctx, in)

	if err != nil {
		return stacktrace.Propagate(err, "Failed to put port %s", port.PortKey)
	}

	return nil
}

func (dpc *DomainPortClient) Close() {
	if dpc.conn != nil {
		dpc.conn.Close()
		dpc.conn = nil
	}
}
