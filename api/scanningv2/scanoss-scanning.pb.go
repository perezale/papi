//
//SPDX-License-Identifier: MIT
//
//Copyright (c) 2021, SCANOSS
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
//Scanning definition details
//*

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: protobuf/scanoss/api/scanning/v2/scanoss-scanning.proto

package scanningv2

import (
	commonv2 "github.com/scanoss/papi/api/commonv2"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_protobuf_scanoss_api_scanning_v2_scanoss_scanning_proto protoreflect.FileDescriptor

var file_protobuf_scanoss_api_scanning_v2_scanoss_scanning_proto_rawDesc = []byte{
	0x0a, 0x37, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x63, 0x61, 0x6e, 0x6f,
	0x73, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x63, 0x61, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x2f,
	0x76, 0x32, 0x2f, 0x73, 0x63, 0x61, 0x6e, 0x6f, 0x73, 0x73, 0x2d, 0x73, 0x63, 0x61, 0x6e, 0x6e,
	0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x73, 0x63, 0x61, 0x6e, 0x6f,
	0x73, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x63, 0x61, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x2e,
	0x76, 0x32, 0x1a, 0x33, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x63, 0x61,
	0x6e, 0x6f, 0x73, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f,
	0x76, 0x32, 0x2f, 0x73, 0x63, 0x61, 0x6e, 0x6f, 0x73, 0x73, 0x2d, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x5b, 0x0a, 0x08, 0x53, 0x63, 0x61, 0x6e, 0x6e,
	0x69, 0x6e, 0x67, 0x12, 0x4f, 0x0a, 0x04, 0x45, 0x63, 0x68, 0x6f, 0x12, 0x22, 0x2e, 0x73, 0x63,
	0x61, 0x6e, 0x6f, 0x73, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x76, 0x32, 0x2e, 0x45, 0x63, 0x68, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x23, 0x2e, 0x73, 0x63, 0x61, 0x6e, 0x6f, 0x73, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x32, 0x2e, 0x45, 0x63, 0x68, 0x6f, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x73, 0x63, 0x61, 0x6e, 0x6f, 0x73, 0x73, 0x2f, 0x70, 0x61, 0x70, 0x69, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x73, 0x63, 0x61, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x76, 0x32, 0x3b, 0x73,
	0x63, 0x61, 0x6e, 0x6e, 0x69, 0x6e, 0x67, 0x76, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var file_protobuf_scanoss_api_scanning_v2_scanoss_scanning_proto_goTypes = []interface{}{
	(*commonv2.EchoRequest)(nil),  // 0: scanoss.api.common.v2.EchoRequest
	(*commonv2.EchoResponse)(nil), // 1: scanoss.api.common.v2.EchoResponse
}
var file_protobuf_scanoss_api_scanning_v2_scanoss_scanning_proto_depIdxs = []int32{
	0, // 0: scanoss.api.scanning.v2.Scanning.Echo:input_type -> scanoss.api.common.v2.EchoRequest
	1, // 1: scanoss.api.scanning.v2.Scanning.Echo:output_type -> scanoss.api.common.v2.EchoResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protobuf_scanoss_api_scanning_v2_scanoss_scanning_proto_init() }
func file_protobuf_scanoss_api_scanning_v2_scanoss_scanning_proto_init() {
	if File_protobuf_scanoss_api_scanning_v2_scanoss_scanning_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protobuf_scanoss_api_scanning_v2_scanoss_scanning_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protobuf_scanoss_api_scanning_v2_scanoss_scanning_proto_goTypes,
		DependencyIndexes: file_protobuf_scanoss_api_scanning_v2_scanoss_scanning_proto_depIdxs,
	}.Build()
	File_protobuf_scanoss_api_scanning_v2_scanoss_scanning_proto = out.File
	file_protobuf_scanoss_api_scanning_v2_scanoss_scanning_proto_rawDesc = nil
	file_protobuf_scanoss_api_scanning_v2_scanoss_scanning_proto_goTypes = nil
	file_protobuf_scanoss_api_scanning_v2_scanoss_scanning_proto_depIdxs = nil
}
