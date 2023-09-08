package api

import (
	"fmt"
	hand "gc/database/app/grpc/handler"
	"gc/database/app/grpc/proto"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func (api *Api) startGRPCServer() error {

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", api.IP, api.GRPC.Port))
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}

	creds, err := credentials.NewServerTLSFromFile(api.CertFile, api.KeyFile)
	if err != nil {
		return fmt.Errorf("could not load TLS keys: %s", err)
	}

	opts := []grpc.ServerOption{grpc.Creds(creds)}

	grpcServer := grpc.NewServer(opts...)

	proto.RegisterServiceDBServer(grpcServer, &hand.Server{})

	log.Printf("Start HTTP/2 gRPC server on: %s", fmt.Sprintf("%s:%d", api.IP, api.GRPC.Port))
	if err := grpcServer.Serve(lis); err != nil {
		return fmt.Errorf("failed to serve: %s", err)
	}

	return nil
}
