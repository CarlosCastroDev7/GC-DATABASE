package handler

import (
	"context"
	"gc/database/app/grpc/proto"
)

// Server :
type Server struct {
	proto.UnimplementedServiceDBServer
}

func (*Server) Connection(ctx context.Context, r *proto.Request) (*proto.Response, error) {

	var response = &proto.Response{
		Json: "hola",
	}

	return response, nil
}
