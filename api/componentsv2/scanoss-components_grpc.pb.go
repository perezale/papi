//
//SPDX-License-Identifier: MIT
//
//Copyright (c) 2022, SCANOSS
//
//Permission is hereby granted, free of charge, to any person obtaining a copy
//of this software and associated documentation files (the "Software"), to deal
//in the Software without restriction, including without limitation the rights
//to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
//copies of the Software, and to permit persons to whom the Software is
//furnished to do so, subject to the following conditions:
//
//The above copyright notice and this permission notice shall be included in
//all copies or substantial portions of the Software.
//
//THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
//IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
//FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
//AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
//LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
//OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
//THE SOFTWARE.

//**
// Component definition details
//*

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: scanoss/api/components/v2/scanoss-components.proto

package componentsv2

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

const (
	Components_Echo_FullMethodName                   = "/scanoss.api.components.v2.Components/Echo"
	Components_SearchComponents_FullMethodName       = "/scanoss.api.components.v2.Components/SearchComponents"
	Components_GetComponentVersions_FullMethodName   = "/scanoss.api.components.v2.Components/GetComponentVersions"
	Components_GetComponentStatistics_FullMethodName = "/scanoss.api.components.v2.Components/GetComponentStatistics"
)

// ComponentsClient is the client API for Components service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ComponentsClient interface {
	// Standard echo
	Echo(ctx context.Context, in *commonv2.EchoRequest, opts ...grpc.CallOption) (*commonv2.EchoResponse, error)
	// Search for components
	SearchComponents(ctx context.Context, in *CompSearchRequest, opts ...grpc.CallOption) (*CompSearchResponse, error)
	// Get all version information for a specific component
	GetComponentVersions(ctx context.Context, in *CompVersionRequest, opts ...grpc.CallOption) (*CompVersionResponse, error)
	// Get the statistics for the specified components
	GetComponentStatistics(ctx context.Context, in *commonv2.PurlRequest, opts ...grpc.CallOption) (*CompStatisticResponse, error)
}

type componentsClient struct {
	cc grpc.ClientConnInterface
}

func NewComponentsClient(cc grpc.ClientConnInterface) ComponentsClient {
	return &componentsClient{cc}
}

func (c *componentsClient) Echo(ctx context.Context, in *commonv2.EchoRequest, opts ...grpc.CallOption) (*commonv2.EchoResponse, error) {
	out := new(commonv2.EchoResponse)
	err := c.cc.Invoke(ctx, Components_Echo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *componentsClient) SearchComponents(ctx context.Context, in *CompSearchRequest, opts ...grpc.CallOption) (*CompSearchResponse, error) {
	out := new(CompSearchResponse)
	err := c.cc.Invoke(ctx, Components_SearchComponents_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *componentsClient) GetComponentVersions(ctx context.Context, in *CompVersionRequest, opts ...grpc.CallOption) (*CompVersionResponse, error) {
	out := new(CompVersionResponse)
	err := c.cc.Invoke(ctx, Components_GetComponentVersions_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *componentsClient) GetComponentStatistics(ctx context.Context, in *commonv2.PurlRequest, opts ...grpc.CallOption) (*CompStatisticResponse, error) {
	out := new(CompStatisticResponse)
	err := c.cc.Invoke(ctx, Components_GetComponentStatistics_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ComponentsServer is the server API for Components service.
// All implementations must embed UnimplementedComponentsServer
// for forward compatibility
type ComponentsServer interface {
	// Standard echo
	Echo(context.Context, *commonv2.EchoRequest) (*commonv2.EchoResponse, error)
	// Search for components
	SearchComponents(context.Context, *CompSearchRequest) (*CompSearchResponse, error)
	// Get all version information for a specific component
	GetComponentVersions(context.Context, *CompVersionRequest) (*CompVersionResponse, error)
	// Get the statistics for the specified components
	GetComponentStatistics(context.Context, *commonv2.PurlRequest) (*CompStatisticResponse, error)
	mustEmbedUnimplementedComponentsServer()
}

// UnimplementedComponentsServer must be embedded to have forward compatible implementations.
type UnimplementedComponentsServer struct {
}

func (UnimplementedComponentsServer) Echo(context.Context, *commonv2.EchoRequest) (*commonv2.EchoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Echo not implemented")
}
func (UnimplementedComponentsServer) SearchComponents(context.Context, *CompSearchRequest) (*CompSearchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchComponents not implemented")
}
func (UnimplementedComponentsServer) GetComponentVersions(context.Context, *CompVersionRequest) (*CompVersionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetComponentVersions not implemented")
}
func (UnimplementedComponentsServer) GetComponentStatistics(context.Context, *commonv2.PurlRequest) (*CompStatisticResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetComponentStatistics not implemented")
}
func (UnimplementedComponentsServer) mustEmbedUnimplementedComponentsServer() {}

// UnsafeComponentsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ComponentsServer will
// result in compilation errors.
type UnsafeComponentsServer interface {
	mustEmbedUnimplementedComponentsServer()
}

func RegisterComponentsServer(s grpc.ServiceRegistrar, srv ComponentsServer) {
	s.RegisterService(&Components_ServiceDesc, srv)
}

func _Components_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(commonv2.EchoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComponentsServer).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Components_Echo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComponentsServer).Echo(ctx, req.(*commonv2.EchoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Components_SearchComponents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CompSearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComponentsServer).SearchComponents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Components_SearchComponents_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComponentsServer).SearchComponents(ctx, req.(*CompSearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Components_GetComponentVersions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CompVersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComponentsServer).GetComponentVersions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Components_GetComponentVersions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComponentsServer).GetComponentVersions(ctx, req.(*CompVersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Components_GetComponentStatistics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(commonv2.PurlRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComponentsServer).GetComponentStatistics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Components_GetComponentStatistics_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComponentsServer).GetComponentStatistics(ctx, req.(*commonv2.PurlRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Components_ServiceDesc is the grpc.ServiceDesc for Components service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Components_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "scanoss.api.components.v2.Components",
	HandlerType: (*ComponentsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Echo",
			Handler:    _Components_Echo_Handler,
		},
		{
			MethodName: "SearchComponents",
			Handler:    _Components_SearchComponents_Handler,
		},
		{
			MethodName: "GetComponentVersions",
			Handler:    _Components_GetComponentVersions_Handler,
		},
		{
			MethodName: "GetComponentStatistics",
			Handler:    _Components_GetComponentStatistics_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "scanoss/api/components/v2/scanoss-components.proto",
}
