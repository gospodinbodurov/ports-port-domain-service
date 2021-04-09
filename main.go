package main

import (
	"flag"
	"log"
	"math"
	"net"

	"github.com/palantir/stacktrace"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	api "github.com/gospodinbodurov/ports-apis/port-domain-service/api"
	"github.com/gospodinbodurov/ports-port-domain-service/service"
	"github.com/gospodinbodurov/ports-port-domain-service/storage"
)

func startGRPCServer(address string) error {
	lis, err := net.Listen("tcp", address)

	if err != nil {
		return stacktrace.Propagate(err, "Failed to bind to %s", address)
	}

	storage := &storage.MapStorage{}

	err = storage.Init()

	if err != nil {
		return stacktrace.Propagate(err, "Failed to init storage")
	}

	s := &service.DomainPortService{
		Storage: storage,
	}

	ServerMaxReceiveMessageSize := math.MaxInt32

	opts := []grpc.ServerOption{grpc.MaxRecvMsgSize(ServerMaxReceiveMessageSize)}
	grpcServer := grpc.NewServer(opts...)

	api.RegisterDomainPortServiceServer(grpcServer, s)

	log.Printf("starting HTTP/2 gRPC server on %s", address)
	reflection.Register(grpcServer)
	if err := grpcServer.Serve(lis); err != nil {
		return stacktrace.Propagate(err, "Failed to serve")
	}

	return nil
}

func main() {
	flag.String("grpcAddress", "localhost:6666", "Provide an url for starting the grpc service")

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)

	grpcAddress := viper.GetString("grpcAddress")

	log.Println("Starting grpc server")
	err := startGRPCServer(grpcAddress)

	if err != nil {
		log.Println(err.Error())
	}
}
