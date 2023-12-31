// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.2
// source: proto/db.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	ServiceDB_Connection_FullMethodName = "/proto.ServiceDB/Connection"
)

// ServiceDBClient is the client API for ServiceDB service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServiceDBClient interface {
	Connection(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type serviceDBClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceDBClient(cc grpc.ClientConnInterface) ServiceDBClient {
	return &serviceDBClient{cc}
}

func (c *serviceDBClient) Connection(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, ServiceDB_Connection_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceDBServer is the server API for ServiceDB service.
// All implementations must embed UnimplementedServiceDBServer
// for forward compatibility
type ServiceDBServer interface {
	Connection(context.Context, *Request) (*Response, error)
	mustEmbedUnimplementedServiceDBServer()
}

// UnimplementedServiceDBServer must be embedded to have forward compatible implementations.
type UnimplementedServiceDBServer struct {
}

func (UnimplementedServiceDBServer) Connection(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Connection not implemented")
}
func (UnimplementedServiceDBServer) mustEmbedUnimplementedServiceDBServer() {}

// UnsafeServiceDBServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServiceDBServer will
// result in compilation errors.
type UnsafeServiceDBServer interface {
	mustEmbedUnimplementedServiceDBServer()
}

func RegisterServiceDBServer(s grpc.ServiceRegistrar, srv ServiceDBServer) {
	s.RegisterService(&ServiceDB_ServiceDesc, srv)
}

func _ServiceDB_Connection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceDBServer).Connection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ServiceDB_Connection_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceDBServer).Connection(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

// ServiceDB_ServiceDesc is the grpc.ServiceDesc for ServiceDB service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ServiceDB_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.ServiceDB",
	HandlerType: (*ServiceDBServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Connection",
			Handler:    _ServiceDB_Connection_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/db.proto",
}
