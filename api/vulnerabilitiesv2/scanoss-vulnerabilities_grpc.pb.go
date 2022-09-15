// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package vulnerabilitiesv2

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

// VulnerabilitiesClient is the client API for Vulnerabilities service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VulnerabilitiesClient interface {
	// Standard echo
	Echo(ctx context.Context, in *commonv2.EchoRequest, opts ...grpc.CallOption) (*commonv2.EchoResponse, error)
	// Get CPEs associated with a PURL
	GetCpes(ctx context.Context, in *VulnerabilityRequest, opts ...grpc.CallOption) (*CpeResponse, error)
	// Get vulnerability details
	GetVulnerabilities(ctx context.Context, in *VulnerabilityRequest, opts ...grpc.CallOption) (*VulnerabilityResponse, error)
}

type vulnerabilitiesClient struct {
	cc grpc.ClientConnInterface
}

func NewVulnerabilitiesClient(cc grpc.ClientConnInterface) VulnerabilitiesClient {
	return &vulnerabilitiesClient{cc}
}

func (c *vulnerabilitiesClient) Echo(ctx context.Context, in *commonv2.EchoRequest, opts ...grpc.CallOption) (*commonv2.EchoResponse, error) {
	out := new(commonv2.EchoResponse)
	err := c.cc.Invoke(ctx, "/scanoss.api.vulnerabilities.v2.Vulnerabilities/Echo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vulnerabilitiesClient) GetCpes(ctx context.Context, in *VulnerabilityRequest, opts ...grpc.CallOption) (*CpeResponse, error) {
	out := new(CpeResponse)
	err := c.cc.Invoke(ctx, "/scanoss.api.vulnerabilities.v2.Vulnerabilities/GetCpes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vulnerabilitiesClient) GetVulnerabilities(ctx context.Context, in *VulnerabilityRequest, opts ...grpc.CallOption) (*VulnerabilityResponse, error) {
	out := new(VulnerabilityResponse)
	err := c.cc.Invoke(ctx, "/scanoss.api.vulnerabilities.v2.Vulnerabilities/GetVulnerabilities", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VulnerabilitiesServer is the server API for Vulnerabilities service.
// All implementations must embed UnimplementedVulnerabilitiesServer
// for forward compatibility
type VulnerabilitiesServer interface {
	// Standard echo
	Echo(context.Context, *commonv2.EchoRequest) (*commonv2.EchoResponse, error)
	// Get CPEs associated with a PURL
	GetCpes(context.Context, *VulnerabilityRequest) (*CpeResponse, error)
	// Get vulnerability details
	GetVulnerabilities(context.Context, *VulnerabilityRequest) (*VulnerabilityResponse, error)
	mustEmbedUnimplementedVulnerabilitiesServer()
}

// UnimplementedVulnerabilitiesServer must be embedded to have forward compatible implementations.
type UnimplementedVulnerabilitiesServer struct {
}

func (UnimplementedVulnerabilitiesServer) Echo(context.Context, *commonv2.EchoRequest) (*commonv2.EchoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Echo not implemented")
}
func (UnimplementedVulnerabilitiesServer) GetCpes(context.Context, *VulnerabilityRequest) (*CpeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCpes not implemented")
}
func (UnimplementedVulnerabilitiesServer) GetVulnerabilities(context.Context, *VulnerabilityRequest) (*VulnerabilityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVulnerabilities not implemented")
}
func (UnimplementedVulnerabilitiesServer) mustEmbedUnimplementedVulnerabilitiesServer() {}

// UnsafeVulnerabilitiesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VulnerabilitiesServer will
// result in compilation errors.
type UnsafeVulnerabilitiesServer interface {
	mustEmbedUnimplementedVulnerabilitiesServer()
}

func RegisterVulnerabilitiesServer(s grpc.ServiceRegistrar, srv VulnerabilitiesServer) {
	s.RegisterService(&Vulnerabilities_ServiceDesc, srv)
}

func _Vulnerabilities_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(commonv2.EchoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VulnerabilitiesServer).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scanoss.api.vulnerabilities.v2.Vulnerabilities/Echo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VulnerabilitiesServer).Echo(ctx, req.(*commonv2.EchoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Vulnerabilities_GetCpes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VulnerabilityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VulnerabilitiesServer).GetCpes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scanoss.api.vulnerabilities.v2.Vulnerabilities/GetCpes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VulnerabilitiesServer).GetCpes(ctx, req.(*VulnerabilityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Vulnerabilities_GetVulnerabilities_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VulnerabilityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VulnerabilitiesServer).GetVulnerabilities(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scanoss.api.vulnerabilities.v2.Vulnerabilities/GetVulnerabilities",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VulnerabilitiesServer).GetVulnerabilities(ctx, req.(*VulnerabilityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Vulnerabilities_ServiceDesc is the grpc.ServiceDesc for Vulnerabilities service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Vulnerabilities_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "scanoss.api.vulnerabilities.v2.Vulnerabilities",
	HandlerType: (*VulnerabilitiesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Echo",
			Handler:    _Vulnerabilities_Echo_Handler,
		},
		{
			MethodName: "GetCpes",
			Handler:    _Vulnerabilities_GetCpes_Handler,
		},
		{
			MethodName: "GetVulnerabilities",
			Handler:    _Vulnerabilities_GetVulnerabilities_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "scanoss/api/vulnerabilities/v2/scanoss-vulnerabilities.proto",
}