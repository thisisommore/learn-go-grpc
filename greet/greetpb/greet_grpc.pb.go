// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package greetpb

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

// GreetServiceClient is the client API for GreetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GreetServiceClient interface {
	GreetManyTimes(ctx context.Context, in *GreetManyTimesRequest, opts ...grpc.CallOption) (GreetService_GreetManyTimesClient, error)
}

type greetServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGreetServiceClient(cc grpc.ClientConnInterface) GreetServiceClient {
	return &greetServiceClient{cc}
}

func (c *greetServiceClient) GreetManyTimes(ctx context.Context, in *GreetManyTimesRequest, opts ...grpc.CallOption) (GreetService_GreetManyTimesClient, error) {
	stream, err := c.cc.NewStream(ctx, &GreetService_ServiceDesc.Streams[0], "/greet.GreetService/GreetManyTimes", opts...)
	if err != nil {
		return nil, err
	}
	x := &greetServiceGreetManyTimesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type GreetService_GreetManyTimesClient interface {
	Recv() (*GreetManyTimesResponse, error)
	grpc.ClientStream
}

type greetServiceGreetManyTimesClient struct {
	grpc.ClientStream
}

func (x *greetServiceGreetManyTimesClient) Recv() (*GreetManyTimesResponse, error) {
	m := new(GreetManyTimesResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GreetServiceServer is the server API for GreetService service.
// All implementations must embed UnimplementedGreetServiceServer
// for forward compatibility
type GreetServiceServer interface {
	GreetManyTimes(*GreetManyTimesRequest, GreetService_GreetManyTimesServer) error
	mustEmbedUnimplementedGreetServiceServer()
}

// UnimplementedGreetServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGreetServiceServer struct {
}

func (UnimplementedGreetServiceServer) GreetManyTimes(*GreetManyTimesRequest, GreetService_GreetManyTimesServer) error {
	return status.Errorf(codes.Unimplemented, "method GreetManyTimes not implemented")
}
func (UnimplementedGreetServiceServer) mustEmbedUnimplementedGreetServiceServer() {}

// UnsafeGreetServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GreetServiceServer will
// result in compilation errors.
type UnsafeGreetServiceServer interface {
	mustEmbedUnimplementedGreetServiceServer()
}

func RegisterGreetServiceServer(s grpc.ServiceRegistrar, srv GreetServiceServer) {
	s.RegisterService(&GreetService_ServiceDesc, srv)
}

func _GreetService_GreetManyTimes_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GreetManyTimesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GreetServiceServer).GreetManyTimes(m, &greetServiceGreetManyTimesServer{stream})
}

type GreetService_GreetManyTimesServer interface {
	Send(*GreetManyTimesResponse) error
	grpc.ServerStream
}

type greetServiceGreetManyTimesServer struct {
	grpc.ServerStream
}

func (x *greetServiceGreetManyTimesServer) Send(m *GreetManyTimesResponse) error {
	return x.ServerStream.SendMsg(m)
}

// GreetService_ServiceDesc is the grpc.ServiceDesc for GreetService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GreetService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "greet.GreetService",
	HandlerType: (*GreetServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GreetManyTimes",
			Handler:       _GreetService_GreetManyTimes_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "greetpb/greet.proto",
}
