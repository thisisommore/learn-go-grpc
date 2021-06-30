// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package calculatorpb

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

// CalulatorServiceClient is the client API for CalulatorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CalulatorServiceClient interface {
	Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddResponse, error)
	Sum(ctx context.Context, opts ...grpc.CallOption) (CalulatorService_SumClient, error)
}

type calulatorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCalulatorServiceClient(cc grpc.ClientConnInterface) CalulatorServiceClient {
	return &calulatorServiceClient{cc}
}

func (c *calulatorServiceClient) Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddResponse, error) {
	out := new(AddResponse)
	err := c.cc.Invoke(ctx, "/calculator.CalulatorService/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calulatorServiceClient) Sum(ctx context.Context, opts ...grpc.CallOption) (CalulatorService_SumClient, error) {
	stream, err := c.cc.NewStream(ctx, &CalulatorService_ServiceDesc.Streams[0], "/calculator.CalulatorService/Sum", opts...)
	if err != nil {
		return nil, err
	}
	x := &calulatorServiceSumClient{stream}
	return x, nil
}

type CalulatorService_SumClient interface {
	Send(*SumRequest) error
	CloseAndRecv() (*SumResponse, error)
	grpc.ClientStream
}

type calulatorServiceSumClient struct {
	grpc.ClientStream
}

func (x *calulatorServiceSumClient) Send(m *SumRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *calulatorServiceSumClient) CloseAndRecv() (*SumResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(SumResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CalulatorServiceServer is the server API for CalulatorService service.
// All implementations must embed UnimplementedCalulatorServiceServer
// for forward compatibility
type CalulatorServiceServer interface {
	Add(context.Context, *AddRequest) (*AddResponse, error)
	Sum(CalulatorService_SumServer) error
	mustEmbedUnimplementedCalulatorServiceServer()
}

// UnimplementedCalulatorServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCalulatorServiceServer struct {
}

func (UnimplementedCalulatorServiceServer) Add(context.Context, *AddRequest) (*AddResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (UnimplementedCalulatorServiceServer) Sum(CalulatorService_SumServer) error {
	return status.Errorf(codes.Unimplemented, "method Sum not implemented")
}
func (UnimplementedCalulatorServiceServer) mustEmbedUnimplementedCalulatorServiceServer() {}

// UnsafeCalulatorServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CalulatorServiceServer will
// result in compilation errors.
type UnsafeCalulatorServiceServer interface {
	mustEmbedUnimplementedCalulatorServiceServer()
}

func RegisterCalulatorServiceServer(s grpc.ServiceRegistrar, srv CalulatorServiceServer) {
	s.RegisterService(&CalulatorService_ServiceDesc, srv)
}

func _CalulatorService_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalulatorServiceServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calculator.CalulatorService/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalulatorServiceServer).Add(ctx, req.(*AddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalulatorService_Sum_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CalulatorServiceServer).Sum(&calulatorServiceSumServer{stream})
}

type CalulatorService_SumServer interface {
	SendAndClose(*SumResponse) error
	Recv() (*SumRequest, error)
	grpc.ServerStream
}

type calulatorServiceSumServer struct {
	grpc.ServerStream
}

func (x *calulatorServiceSumServer) SendAndClose(m *SumResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *calulatorServiceSumServer) Recv() (*SumRequest, error) {
	m := new(SumRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CalulatorService_ServiceDesc is the grpc.ServiceDesc for CalulatorService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CalulatorService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "calculator.CalulatorService",
	HandlerType: (*CalulatorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _CalulatorService_Add_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Sum",
			Handler:       _CalulatorService_Sum_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "calculator/calculatorpb/calculator.proto",
}
