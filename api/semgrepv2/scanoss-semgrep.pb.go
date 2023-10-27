//
//SPDX-License-Identifier: MIT
//
//Copyright (c) 2023, SCANOSS
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
// semgrep definition details
//*

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: scanoss/api/semgrep/v2/scanoss-semgrep.proto

package semgrepv2

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	commonv2 "github.com/scanoss/papi/api/commonv2"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Semgrep issue response data (JSON payload)
type SemgrepResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Cryptography details
	Purls []*SemgrepResponse_Purls `protobuf:"bytes,1,rep,name=purls,proto3" json:"purls,omitempty"`
	// Response status
	Status *commonv2.StatusResponse `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *SemgrepResponse) Reset() {
	*x = SemgrepResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scanoss_api_semgrep_v2_scanoss_semgrep_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SemgrepResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SemgrepResponse) ProtoMessage() {}

func (x *SemgrepResponse) ProtoReflect() protoreflect.Message {
	mi := &file_scanoss_api_semgrep_v2_scanoss_semgrep_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SemgrepResponse.ProtoReflect.Descriptor instead.
func (*SemgrepResponse) Descriptor() ([]byte, []int) {
	return file_scanoss_api_semgrep_v2_scanoss_semgrep_proto_rawDescGZIP(), []int{0}
}

func (x *SemgrepResponse) GetPurls() []*SemgrepResponse_Purls {
	if x != nil {
		return x.Purls
	}
	return nil
}

func (x *SemgrepResponse) GetStatus() *commonv2.StatusResponse {
	if x != nil {
		return x.Status
	}
	return nil
}

type SemgrepResponse_Issue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RuleID   string `protobuf:"bytes,1,opt,name=ruleID,proto3" json:"ruleID,omitempty"`
	From     string `protobuf:"bytes,2,opt,name=from,proto3" json:"from,omitempty"`
	To       string `protobuf:"bytes,3,opt,name=to,proto3" json:"to,omitempty"`
	Severity string `protobuf:"bytes,4,opt,name=severity,proto3" json:"severity,omitempty"`
}

func (x *SemgrepResponse_Issue) Reset() {
	*x = SemgrepResponse_Issue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scanoss_api_semgrep_v2_scanoss_semgrep_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SemgrepResponse_Issue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SemgrepResponse_Issue) ProtoMessage() {}

func (x *SemgrepResponse_Issue) ProtoReflect() protoreflect.Message {
	mi := &file_scanoss_api_semgrep_v2_scanoss_semgrep_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SemgrepResponse_Issue.ProtoReflect.Descriptor instead.
func (*SemgrepResponse_Issue) Descriptor() ([]byte, []int) {
	return file_scanoss_api_semgrep_v2_scanoss_semgrep_proto_rawDescGZIP(), []int{0, 0}
}

func (x *SemgrepResponse_Issue) GetRuleID() string {
	if x != nil {
		return x.RuleID
	}
	return ""
}

func (x *SemgrepResponse_Issue) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *SemgrepResponse_Issue) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *SemgrepResponse_Issue) GetSeverity() string {
	if x != nil {
		return x.Severity
	}
	return ""
}

type SemgrepResponse_File struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileMD5 string                   `protobuf:"bytes,1,opt,name=fileMD5,proto3" json:"fileMD5,omitempty"`
	Path    string                   `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	Issues  []*SemgrepResponse_Issue `protobuf:"bytes,3,rep,name=issues,proto3" json:"issues,omitempty"`
}

func (x *SemgrepResponse_File) Reset() {
	*x = SemgrepResponse_File{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scanoss_api_semgrep_v2_scanoss_semgrep_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SemgrepResponse_File) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SemgrepResponse_File) ProtoMessage() {}

func (x *SemgrepResponse_File) ProtoReflect() protoreflect.Message {
	mi := &file_scanoss_api_semgrep_v2_scanoss_semgrep_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SemgrepResponse_File.ProtoReflect.Descriptor instead.
func (*SemgrepResponse_File) Descriptor() ([]byte, []int) {
	return file_scanoss_api_semgrep_v2_scanoss_semgrep_proto_rawDescGZIP(), []int{0, 1}
}

func (x *SemgrepResponse_File) GetFileMD5() string {
	if x != nil {
		return x.FileMD5
	}
	return ""
}

func (x *SemgrepResponse_File) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *SemgrepResponse_File) GetIssues() []*SemgrepResponse_Issue {
	if x != nil {
		return x.Issues
	}
	return nil
}

type SemgrepResponse_Purls struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Purl    string                  `protobuf:"bytes,1,opt,name=purl,proto3" json:"purl,omitempty"`
	Version string                  `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	Files   []*SemgrepResponse_File `protobuf:"bytes,3,rep,name=files,proto3" json:"files,omitempty"`
}

func (x *SemgrepResponse_Purls) Reset() {
	*x = SemgrepResponse_Purls{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scanoss_api_semgrep_v2_scanoss_semgrep_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SemgrepResponse_Purls) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SemgrepResponse_Purls) ProtoMessage() {}

func (x *SemgrepResponse_Purls) ProtoReflect() protoreflect.Message {
	mi := &file_scanoss_api_semgrep_v2_scanoss_semgrep_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SemgrepResponse_Purls.ProtoReflect.Descriptor instead.
func (*SemgrepResponse_Purls) Descriptor() ([]byte, []int) {
	return file_scanoss_api_semgrep_v2_scanoss_semgrep_proto_rawDescGZIP(), []int{0, 2}
}

func (x *SemgrepResponse_Purls) GetPurl() string {
	if x != nil {
		return x.Purl
	}
	return ""
}

func (x *SemgrepResponse_Purls) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *SemgrepResponse_Purls) GetFiles() []*SemgrepResponse_File {
	if x != nil {
		return x.Files
	}
	return nil
}

var File_scanoss_api_semgrep_v2_scanoss_semgrep_proto protoreflect.FileDescriptor

var file_scanoss_api_semgrep_v2_scanoss_semgrep_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x73, 0x63, 0x61, 0x6e, 0x6f, 0x73, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x65,
	0x6d, 0x67, 0x72, 0x65, 0x70, 0x2f, 0x76, 0x32, 0x2f, 0x73, 0x63, 0x61, 0x6e, 0x6f, 0x73, 0x73,
	0x2d, 0x73, 0x65, 0x6d, 0x67, 0x72, 0x65, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16,
	0x73, 0x63, 0x61, 0x6e, 0x6f, 0x73, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x6d, 0x67,
	0x72, 0x65, 0x70, 0x2e, 0x76, 0x32, 0x1a, 0x2a, 0x73, 0x63, 0x61, 0x6e, 0x6f, 0x73, 0x73, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x32, 0x2f, 0x73, 0x63,
	0x61, 0x6e, 0x6f, 0x73, 0x73, 0x2d, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x2c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x73, 0x77, 0x61,
	0x67, 0x67, 0x65, 0x72, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xee,
	0x03, 0x0a, 0x0f, 0x53, 0x65, 0x6d, 0x67, 0x72, 0x65, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x43, 0x0a, 0x05, 0x70, 0x75, 0x72, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x2d, 0x2e, 0x73, 0x63, 0x61, 0x6e, 0x6f, 0x73, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x73, 0x65, 0x6d, 0x67, 0x72, 0x65, 0x70, 0x2e, 0x76, 0x32, 0x2e, 0x53, 0x65, 0x6d, 0x67, 0x72,
	0x65, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x50, 0x75, 0x72, 0x6c, 0x73,
	0x52, 0x05, 0x70, 0x75, 0x72, 0x6c, 0x73, 0x12, 0x3d, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x73, 0x63, 0x61, 0x6e, 0x6f, 0x73,
	0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x32, 0x2e,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x1a, 0x5f, 0x0a, 0x05, 0x49, 0x73, 0x73, 0x75, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x72, 0x75, 0x6c, 0x65, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x72, 0x75, 0x6c, 0x65, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x74,
	0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x73,
	0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73,
	0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x1a, 0x7b, 0x0a, 0x04, 0x46, 0x69, 0x6c, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x66, 0x69, 0x6c, 0x65, 0x4d, 0x44, 0x35, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x66, 0x69, 0x6c, 0x65, 0x4d, 0x44, 0x35, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74,
	0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x45, 0x0a,
	0x06, 0x69, 0x73, 0x73, 0x75, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2d, 0x2e,
	0x73, 0x63, 0x61, 0x6e, 0x6f, 0x73, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x6d, 0x67,
	0x72, 0x65, 0x70, 0x2e, 0x76, 0x32, 0x2e, 0x53, 0x65, 0x6d, 0x67, 0x72, 0x65, 0x70, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x49, 0x73, 0x73, 0x75, 0x65, 0x52, 0x06, 0x69, 0x73,
	0x73, 0x75, 0x65, 0x73, 0x1a, 0x79, 0x0a, 0x05, 0x50, 0x75, 0x72, 0x6c, 0x73, 0x12, 0x12, 0x0a,
	0x04, 0x70, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x75, 0x72,
	0x6c, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x42, 0x0a, 0x05, 0x66,
	0x69, 0x6c, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x73, 0x63, 0x61,
	0x6e, 0x6f, 0x73, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x6d, 0x67, 0x72, 0x65, 0x70,
	0x2e, 0x76, 0x32, 0x2e, 0x53, 0x65, 0x6d, 0x67, 0x72, 0x65, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x32,
	0xf8, 0x01, 0x0a, 0x07, 0x53, 0x65, 0x6d, 0x67, 0x72, 0x65, 0x70, 0x12, 0x70, 0x0a, 0x04, 0x45,
	0x63, 0x68, 0x6f, 0x12, 0x22, 0x2e, 0x73, 0x63, 0x61, 0x6e, 0x6f, 0x73, 0x73, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x32, 0x2e, 0x45, 0x63, 0x68, 0x6f,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x73, 0x63, 0x61, 0x6e, 0x6f, 0x73,
	0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x32, 0x2e,
	0x45, 0x63, 0x68, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1f, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x19, 0x22, 0x14, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x32, 0x2f, 0x73, 0x65,
	0x6d, 0x67, 0x72, 0x65, 0x70, 0x2f, 0x65, 0x63, 0x68, 0x6f, 0x3a, 0x01, 0x2a, 0x12, 0x7b, 0x0a,
	0x09, 0x47, 0x65, 0x74, 0x49, 0x73, 0x73, 0x75, 0x65, 0x73, 0x12, 0x22, 0x2e, 0x73, 0x63, 0x61,
	0x6e, 0x6f, 0x73, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x76, 0x32, 0x2e, 0x50, 0x75, 0x72, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27,
	0x2e, 0x73, 0x63, 0x61, 0x6e, 0x6f, 0x73, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x6d,
	0x67, 0x72, 0x65, 0x70, 0x2e, 0x76, 0x32, 0x2e, 0x53, 0x65, 0x6d, 0x67, 0x72, 0x65, 0x70, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x22,
	0x16, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x32, 0x2f, 0x73, 0x65, 0x6d, 0x67, 0x72, 0x65, 0x70,
	0x2f, 0x69, 0x73, 0x73, 0x75, 0x65, 0x73, 0x3a, 0x01, 0x2a, 0x42, 0x85, 0x02, 0x5a, 0x2f, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x63, 0x61, 0x6e, 0x6f, 0x73,
	0x73, 0x2f, 0x70, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x65, 0x6d, 0x67, 0x72,
	0x65, 0x70, 0x76, 0x32, 0x3b, 0x73, 0x65, 0x6d, 0x67, 0x72, 0x65, 0x70, 0x76, 0x32, 0x92, 0x41,
	0xd0, 0x01, 0x12, 0x6a, 0x0a, 0x17, 0x53, 0x43, 0x41, 0x4e, 0x4f, 0x53, 0x53, 0x20, 0x53, 0x65,
	0x6d, 0x67, 0x72, 0x65, 0x70, 0x20, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x4a, 0x0a,
	0x0f, 0x73, 0x63, 0x61, 0x6e, 0x6f, 0x73, 0x73, 0x2d, 0x73, 0x65, 0x6d, 0x67, 0x72, 0x65, 0x70,
	0x12, 0x22, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x63, 0x61, 0x6e, 0x6f, 0x73, 0x73, 0x2f, 0x73, 0x65, 0x6d,
	0x67, 0x72, 0x65, 0x70, 0x1a, 0x13, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x40, 0x73, 0x63,
	0x61, 0x6e, 0x6f, 0x73, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x32, 0x03, 0x32, 0x2e, 0x30, 0x2a, 0x01,
	0x01, 0x32, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a,
	0x73, 0x6f, 0x6e, 0x3a, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x52, 0x3b, 0x0a, 0x03, 0x34, 0x30, 0x34, 0x12, 0x34, 0x0a, 0x2a,
	0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x65, 0x64, 0x20, 0x77, 0x68, 0x65, 0x6e, 0x20, 0x74, 0x68,
	0x65, 0x20, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x20, 0x64, 0x6f, 0x65, 0x73, 0x20,
	0x6e, 0x6f, 0x74, 0x20, 0x65, 0x78, 0x69, 0x73, 0x74, 0x2e, 0x12, 0x06, 0x0a, 0x04, 0x9a, 0x02,
	0x01, 0x07, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_scanoss_api_semgrep_v2_scanoss_semgrep_proto_rawDescOnce sync.Once
	file_scanoss_api_semgrep_v2_scanoss_semgrep_proto_rawDescData = file_scanoss_api_semgrep_v2_scanoss_semgrep_proto_rawDesc
)

func file_scanoss_api_semgrep_v2_scanoss_semgrep_proto_rawDescGZIP() []byte {
	file_scanoss_api_semgrep_v2_scanoss_semgrep_proto_rawDescOnce.Do(func() {
		file_scanoss_api_semgrep_v2_scanoss_semgrep_proto_rawDescData = protoimpl.X.CompressGZIP(file_scanoss_api_semgrep_v2_scanoss_semgrep_proto_rawDescData)
	})
	return file_scanoss_api_semgrep_v2_scanoss_semgrep_proto_rawDescData
}

var file_scanoss_api_semgrep_v2_scanoss_semgrep_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_scanoss_api_semgrep_v2_scanoss_semgrep_proto_goTypes = []interface{}{
	(*SemgrepResponse)(nil),         // 0: scanoss.api.semgrep.v2.SemgrepResponse
	(*SemgrepResponse_Issue)(nil),   // 1: scanoss.api.semgrep.v2.SemgrepResponse.Issue
	(*SemgrepResponse_File)(nil),    // 2: scanoss.api.semgrep.v2.SemgrepResponse.File
	(*SemgrepResponse_Purls)(nil),   // 3: scanoss.api.semgrep.v2.SemgrepResponse.Purls
	(*commonv2.StatusResponse)(nil), // 4: scanoss.api.common.v2.StatusResponse
	(*commonv2.EchoRequest)(nil),    // 5: scanoss.api.common.v2.EchoRequest
	(*commonv2.PurlRequest)(nil),    // 6: scanoss.api.common.v2.PurlRequest
	(*commonv2.EchoResponse)(nil),   // 7: scanoss.api.common.v2.EchoResponse
}
var file_scanoss_api_semgrep_v2_scanoss_semgrep_proto_depIdxs = []int32{
	3, // 0: scanoss.api.semgrep.v2.SemgrepResponse.purls:type_name -> scanoss.api.semgrep.v2.SemgrepResponse.Purls
	4, // 1: scanoss.api.semgrep.v2.SemgrepResponse.status:type_name -> scanoss.api.common.v2.StatusResponse
	1, // 2: scanoss.api.semgrep.v2.SemgrepResponse.File.issues:type_name -> scanoss.api.semgrep.v2.SemgrepResponse.Issue
	2, // 3: scanoss.api.semgrep.v2.SemgrepResponse.Purls.files:type_name -> scanoss.api.semgrep.v2.SemgrepResponse.File
	5, // 4: scanoss.api.semgrep.v2.Semgrep.Echo:input_type -> scanoss.api.common.v2.EchoRequest
	6, // 5: scanoss.api.semgrep.v2.Semgrep.GetIssues:input_type -> scanoss.api.common.v2.PurlRequest
	7, // 6: scanoss.api.semgrep.v2.Semgrep.Echo:output_type -> scanoss.api.common.v2.EchoResponse
	0, // 7: scanoss.api.semgrep.v2.Semgrep.GetIssues:output_type -> scanoss.api.semgrep.v2.SemgrepResponse
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_scanoss_api_semgrep_v2_scanoss_semgrep_proto_init() }
func file_scanoss_api_semgrep_v2_scanoss_semgrep_proto_init() {
	if File_scanoss_api_semgrep_v2_scanoss_semgrep_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_scanoss_api_semgrep_v2_scanoss_semgrep_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SemgrepResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_scanoss_api_semgrep_v2_scanoss_semgrep_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SemgrepResponse_Issue); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_scanoss_api_semgrep_v2_scanoss_semgrep_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SemgrepResponse_File); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_scanoss_api_semgrep_v2_scanoss_semgrep_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SemgrepResponse_Purls); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_scanoss_api_semgrep_v2_scanoss_semgrep_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_scanoss_api_semgrep_v2_scanoss_semgrep_proto_goTypes,
		DependencyIndexes: file_scanoss_api_semgrep_v2_scanoss_semgrep_proto_depIdxs,
		MessageInfos:      file_scanoss_api_semgrep_v2_scanoss_semgrep_proto_msgTypes,
	}.Build()
	File_scanoss_api_semgrep_v2_scanoss_semgrep_proto = out.File
	file_scanoss_api_semgrep_v2_scanoss_semgrep_proto_rawDesc = nil
	file_scanoss_api_semgrep_v2_scanoss_semgrep_proto_goTypes = nil
	file_scanoss_api_semgrep_v2_scanoss_semgrep_proto_depIdxs = nil
}
