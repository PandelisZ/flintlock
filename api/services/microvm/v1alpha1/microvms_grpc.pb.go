// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: services/microvm/v1alpha1/microvms.proto

package v1alpha1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	MicroVM_CreateMicroVM_FullMethodName      = "/microvm.services.api.v1alpha1.MicroVM/CreateMicroVM"
	MicroVM_DeleteMicroVM_FullMethodName      = "/microvm.services.api.v1alpha1.MicroVM/DeleteMicroVM"
	MicroVM_GetMicroVM_FullMethodName         = "/microvm.services.api.v1alpha1.MicroVM/GetMicroVM"
	MicroVM_ListMicroVMs_FullMethodName       = "/microvm.services.api.v1alpha1.MicroVM/ListMicroVMs"
	MicroVM_ListMicroVMsStream_FullMethodName = "/microvm.services.api.v1alpha1.MicroVM/ListMicroVMsStream"
)

// MicroVMClient is the client API for MicroVM service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// MicroVM providers a service to create and manage the lifecycle of microvms.
type MicroVMClient interface {
	CreateMicroVM(ctx context.Context, in *CreateMicroVMRequest, opts ...grpc.CallOption) (*CreateMicroVMResponse, error)
	DeleteMicroVM(ctx context.Context, in *DeleteMicroVMRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetMicroVM(ctx context.Context, in *GetMicroVMRequest, opts ...grpc.CallOption) (*GetMicroVMResponse, error)
	ListMicroVMs(ctx context.Context, in *ListMicroVMsRequest, opts ...grpc.CallOption) (*ListMicroVMsResponse, error)
	ListMicroVMsStream(ctx context.Context, in *ListMicroVMsRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[ListMessage], error)
}

type microVMClient struct {
	cc grpc.ClientConnInterface
}

func NewMicroVMClient(cc grpc.ClientConnInterface) MicroVMClient {
	return &microVMClient{cc}
}

func (c *microVMClient) CreateMicroVM(ctx context.Context, in *CreateMicroVMRequest, opts ...grpc.CallOption) (*CreateMicroVMResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateMicroVMResponse)
	err := c.cc.Invoke(ctx, MicroVM_CreateMicroVM_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *microVMClient) DeleteMicroVM(ctx context.Context, in *DeleteMicroVMRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, MicroVM_DeleteMicroVM_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *microVMClient) GetMicroVM(ctx context.Context, in *GetMicroVMRequest, opts ...grpc.CallOption) (*GetMicroVMResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetMicroVMResponse)
	err := c.cc.Invoke(ctx, MicroVM_GetMicroVM_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *microVMClient) ListMicroVMs(ctx context.Context, in *ListMicroVMsRequest, opts ...grpc.CallOption) (*ListMicroVMsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListMicroVMsResponse)
	err := c.cc.Invoke(ctx, MicroVM_ListMicroVMs_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *microVMClient) ListMicroVMsStream(ctx context.Context, in *ListMicroVMsRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[ListMessage], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &MicroVM_ServiceDesc.Streams[0], MicroVM_ListMicroVMsStream_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[ListMicroVMsRequest, ListMessage]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type MicroVM_ListMicroVMsStreamClient = grpc.ServerStreamingClient[ListMessage]

// MicroVMServer is the server API for MicroVM service.
// All implementations should embed UnimplementedMicroVMServer
// for forward compatibility.
//
// MicroVM providers a service to create and manage the lifecycle of microvms.
type MicroVMServer interface {
	CreateMicroVM(context.Context, *CreateMicroVMRequest) (*CreateMicroVMResponse, error)
	DeleteMicroVM(context.Context, *DeleteMicroVMRequest) (*emptypb.Empty, error)
	GetMicroVM(context.Context, *GetMicroVMRequest) (*GetMicroVMResponse, error)
	ListMicroVMs(context.Context, *ListMicroVMsRequest) (*ListMicroVMsResponse, error)
	ListMicroVMsStream(*ListMicroVMsRequest, grpc.ServerStreamingServer[ListMessage]) error
}

// UnimplementedMicroVMServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedMicroVMServer struct{}

func (UnimplementedMicroVMServer) CreateMicroVM(context.Context, *CreateMicroVMRequest) (*CreateMicroVMResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMicroVM not implemented")
}
func (UnimplementedMicroVMServer) DeleteMicroVM(context.Context, *DeleteMicroVMRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteMicroVM not implemented")
}
func (UnimplementedMicroVMServer) GetMicroVM(context.Context, *GetMicroVMRequest) (*GetMicroVMResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMicroVM not implemented")
}
func (UnimplementedMicroVMServer) ListMicroVMs(context.Context, *ListMicroVMsRequest) (*ListMicroVMsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListMicroVMs not implemented")
}
func (UnimplementedMicroVMServer) ListMicroVMsStream(*ListMicroVMsRequest, grpc.ServerStreamingServer[ListMessage]) error {
	return status.Errorf(codes.Unimplemented, "method ListMicroVMsStream not implemented")
}
func (UnimplementedMicroVMServer) testEmbeddedByValue() {}

// UnsafeMicroVMServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MicroVMServer will
// result in compilation errors.
type UnsafeMicroVMServer interface {
	mustEmbedUnimplementedMicroVMServer()
}

func RegisterMicroVMServer(s grpc.ServiceRegistrar, srv MicroVMServer) {
	// If the following call pancis, it indicates UnimplementedMicroVMServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&MicroVM_ServiceDesc, srv)
}

func _MicroVM_CreateMicroVM_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateMicroVMRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MicroVMServer).CreateMicroVM(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MicroVM_CreateMicroVM_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MicroVMServer).CreateMicroVM(ctx, req.(*CreateMicroVMRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MicroVM_DeleteMicroVM_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteMicroVMRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MicroVMServer).DeleteMicroVM(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MicroVM_DeleteMicroVM_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MicroVMServer).DeleteMicroVM(ctx, req.(*DeleteMicroVMRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MicroVM_GetMicroVM_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMicroVMRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MicroVMServer).GetMicroVM(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MicroVM_GetMicroVM_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MicroVMServer).GetMicroVM(ctx, req.(*GetMicroVMRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MicroVM_ListMicroVMs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListMicroVMsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MicroVMServer).ListMicroVMs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MicroVM_ListMicroVMs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MicroVMServer).ListMicroVMs(ctx, req.(*ListMicroVMsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MicroVM_ListMicroVMsStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListMicroVMsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MicroVMServer).ListMicroVMsStream(m, &grpc.GenericServerStream[ListMicroVMsRequest, ListMessage]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type MicroVM_ListMicroVMsStreamServer = grpc.ServerStreamingServer[ListMessage]

// MicroVM_ServiceDesc is the grpc.ServiceDesc for MicroVM service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MicroVM_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "microvm.services.api.v1alpha1.MicroVM",
	HandlerType: (*MicroVMServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateMicroVM",
			Handler:    _MicroVM_CreateMicroVM_Handler,
		},
		{
			MethodName: "DeleteMicroVM",
			Handler:    _MicroVM_DeleteMicroVM_Handler,
		},
		{
			MethodName: "GetMicroVM",
			Handler:    _MicroVM_GetMicroVM_Handler,
		},
		{
			MethodName: "ListMicroVMs",
			Handler:    _MicroVM_ListMicroVMs_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListMicroVMsStream",
			Handler:       _MicroVM_ListMicroVMsStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "services/microvm/v1alpha1/microvms.proto",
}
