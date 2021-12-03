// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package dependenciesv2

import (
	context "context"
	commonv2 "github.com/scanoss/papi/api/commonv2"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// DependenciesClient is the client API for Dependencies service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DependenciesClient interface {
	// Standard echo
	Echo(ctx context.Context, in *commonv2.EchoRequest, opts ...grpc.CallOption) (*commonv2.EchoResponse, error)
	// Get dependency details
	GetDependencies(ctx context.Context, in *DependencyRequest, opts ...grpc.CallOption) (*DependencyResponse, error)
}

type dependenciesClient struct {
	cc grpc.ClientConnInterface
}

func NewDependenciesClient(cc grpc.ClientConnInterface) DependenciesClient {
	return &dependenciesClient{cc}
}

func (c *dependenciesClient) Echo(ctx context.Context, in *commonv2.EchoRequest, opts ...grpc.CallOption) (*commonv2.EchoResponse, error) {
	out := new(commonv2.EchoResponse)
	err := c.cc.Invoke(ctx, "/scanoss.api.dependencies.v2.Dependencies/Echo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dependenciesClient) GetDependencies(ctx context.Context, in *DependencyRequest, opts ...grpc.CallOption) (*DependencyResponse, error) {
	out := new(DependencyResponse)
	err := c.cc.Invoke(ctx, "/scanoss.api.dependencies.v2.Dependencies/GetDependencies", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DependenciesServer is the server API for Dependencies service.
// All implementations must embed UnimplementedDependenciesServer
// for forward compatibility
type DependenciesServer interface {
	// Standard echo
	Echo(context.Context, *commonv2.EchoRequest) (*commonv2.EchoResponse, error)
	// Get dependency details
	GetDependencies(context.Context, *DependencyRequest) (*DependencyResponse, error)
	mustEmbedUnimplementedDependenciesServer()
}

// UnimplementedDependenciesServer must be embedded to have forward compatible implementations.
type UnimplementedDependenciesServer struct {
}

func (UnimplementedDependenciesServer) Echo(context.Context, *commonv2.EchoRequest) (*commonv2.EchoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Echo not implemented")
}
func (UnimplementedDependenciesServer) GetDependencies(context.Context, *DependencyRequest) (*DependencyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDependencies not implemented")
}
func (UnimplementedDependenciesServer) mustEmbedUnimplementedDependenciesServer() {}

// UnsafeDependenciesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DependenciesServer will
// result in compilation errors.
type UnsafeDependenciesServer interface {
	mustEmbedUnimplementedDependenciesServer()
}

func RegisterDependenciesServer(s grpc.ServiceRegistrar, srv DependenciesServer) {
	s.RegisterService(&Dependencies_ServiceDesc, srv)
}

func _Dependencies_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(commonv2.EchoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DependenciesServer).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scanoss.api.dependencies.v2.Dependencies/Echo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DependenciesServer).Echo(ctx, req.(*commonv2.EchoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dependencies_GetDependencies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DependencyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DependenciesServer).GetDependencies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scanoss.api.dependencies.v2.Dependencies/GetDependencies",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DependenciesServer).GetDependencies(ctx, req.(*DependencyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Dependencies_ServiceDesc is the grpc.ServiceDesc for Dependencies service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Dependencies_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "scanoss.api.dependencies.v2.Dependencies",
	HandlerType: (*DependenciesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Echo",
			Handler:    _Dependencies_Echo_Handler,
		},
		{
			MethodName: "GetDependencies",
			Handler:    _Dependencies_GetDependencies_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protobuf/scanoss/api/dependencies/v2/scanoss-dependencies.proto",
}
