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
// Cryptography definition details
//*

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.21.12
// source: scanoss/api/cryptography/v2/scanoss-cryptography.proto

package cryptographyv2

import (
	commonv2 "github.com/scanoss/papi/api/commonv2"
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

//
// Cryptography Algorithm request data (JSON payload)
type CryptographyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// JSON cryptography request purls
	Purls []*CryptographyRequest_Purls `protobuf:"bytes,1,rep,name=purls,proto3" json:"purls,omitempty"`
}

func (x *CryptographyRequest) Reset() {
	*x = CryptographyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scanoss_api_cryptography_v2_scanoss_cryptography_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CryptographyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CryptographyRequest) ProtoMessage() {}

func (x *CryptographyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_scanoss_api_cryptography_v2_scanoss_cryptography_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CryptographyRequest.ProtoReflect.Descriptor instead.
func (*CryptographyRequest) Descriptor() ([]byte, []int) {
	return file_scanoss_api_cryptography_v2_scanoss_cryptography_proto_rawDescGZIP(), []int{0}
}

func (x *CryptographyRequest) GetPurls() []*CryptographyRequest_Purls {
	if x != nil {
		return x.Purls
	}
	return nil
}

//
// Cryptography Algorithm response data (JSON payload)
type CryptographyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Cryptography details
	Purls []*CryptographyResponse_Purls `protobuf:"bytes,1,rep,name=purls,proto3" json:"purls,omitempty"`
	// Response status
	Status *commonv2.StatusResponse `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *CryptographyResponse) Reset() {
	*x = CryptographyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scanoss_api_cryptography_v2_scanoss_cryptography_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CryptographyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CryptographyResponse) ProtoMessage() {}

func (x *CryptographyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_scanoss_api_cryptography_v2_scanoss_cryptography_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CryptographyResponse.ProtoReflect.Descriptor instead.
func (*CryptographyResponse) Descriptor() ([]byte, []int) {
	return file_scanoss_api_cryptography_v2_scanoss_cryptography_proto_rawDescGZIP(), []int{1}
}

func (x *CryptographyResponse) GetPurls() []*CryptographyResponse_Purls {
	if x != nil {
		return x.Purls
	}
	return nil
}

func (x *CryptographyResponse) GetStatus() *commonv2.StatusResponse {
	if x != nil {
		return x.Status
	}
	return nil
}

type CryptographyRequest_Purls struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Purl        string `protobuf:"bytes,1,opt,name=purl,proto3" json:"purl,omitempty"`
	Requirement string `protobuf:"bytes,2,opt,name=requirement,proto3" json:"requirement,omitempty"`
}

func (x *CryptographyRequest_Purls) Reset() {
	*x = CryptographyRequest_Purls{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scanoss_api_cryptography_v2_scanoss_cryptography_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CryptographyRequest_Purls) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CryptographyRequest_Purls) ProtoMessage() {}

func (x *CryptographyRequest_Purls) ProtoReflect() protoreflect.Message {
	mi := &file_scanoss_api_cryptography_v2_scanoss_cryptography_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CryptographyRequest_Purls.ProtoReflect.Descriptor instead.
func (*CryptographyRequest_Purls) Descriptor() ([]byte, []int) {
	return file_scanoss_api_cryptography_v2_scanoss_cryptography_proto_rawDescGZIP(), []int{0, 0}
}

func (x *CryptographyRequest_Purls) GetPurl() string {
	if x != nil {
		return x.Purl
	}
	return ""
}

func (x *CryptographyRequest_Purls) GetRequirement() string {
	if x != nil {
		return x.Requirement
	}
	return ""
}

type CryptographyResponse_Algorithms struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Algorithm string `protobuf:"bytes,1,opt,name=algorithm,proto3" json:"algorithm,omitempty"`
	Strength  string `protobuf:"bytes,2,opt,name=strength,proto3" json:"strength,omitempty"`
	Count     int32  `protobuf:"varint,3,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *CryptographyResponse_Algorithms) Reset() {
	*x = CryptographyResponse_Algorithms{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scanoss_api_cryptography_v2_scanoss_cryptography_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CryptographyResponse_Algorithms) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CryptographyResponse_Algorithms) ProtoMessage() {}

func (x *CryptographyResponse_Algorithms) ProtoReflect() protoreflect.Message {
	mi := &file_scanoss_api_cryptography_v2_scanoss_cryptography_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CryptographyResponse_Algorithms.ProtoReflect.Descriptor instead.
func (*CryptographyResponse_Algorithms) Descriptor() ([]byte, []int) {
	return file_scanoss_api_cryptography_v2_scanoss_cryptography_proto_rawDescGZIP(), []int{1, 0}
}

func (x *CryptographyResponse_Algorithms) GetAlgorithm() string {
	if x != nil {
		return x.Algorithm
	}
	return ""
}

func (x *CryptographyResponse_Algorithms) GetStrength() string {
	if x != nil {
		return x.Strength
	}
	return ""
}

func (x *CryptographyResponse_Algorithms) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

type CryptographyResponse_Purls struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Purl      string                             `protobuf:"bytes,1,opt,name=purl,proto3" json:"purl,omitempty"`
	Algorithm []*CryptographyResponse_Algorithms `protobuf:"bytes,2,rep,name=algorithm,proto3" json:"algorithm,omitempty"`
}

func (x *CryptographyResponse_Purls) Reset() {
	*x = CryptographyResponse_Purls{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scanoss_api_cryptography_v2_scanoss_cryptography_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CryptographyResponse_Purls) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CryptographyResponse_Purls) ProtoMessage() {}

func (x *CryptographyResponse_Purls) ProtoReflect() protoreflect.Message {
	mi := &file_scanoss_api_cryptography_v2_scanoss_cryptography_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CryptographyResponse_Purls.ProtoReflect.Descriptor instead.
func (*CryptographyResponse_Purls) Descriptor() ([]byte, []int) {
	return file_scanoss_api_cryptography_v2_scanoss_cryptography_proto_rawDescGZIP(), []int{1, 1}
}

func (x *CryptographyResponse_Purls) GetPurl() string {
	if x != nil {
		return x.Purl
	}
	return ""
}

func (x *CryptographyResponse_Purls) GetAlgorithm() []*CryptographyResponse_Algorithms {
	if x != nil {
		return x.Algorithm
	}
	return nil
}

var File_scanoss_api_cryptography_v2_scanoss_cryptography_proto protoreflect.FileDescriptor

var file_scanoss_api_cryptography_v2_scanoss_cryptography_proto_rawDesc = []byte{
	0x0a, 0x36, 0x73, 0x63, 0x61, 0x6e, 0x6f, 0x73, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x72,
	0x79, 0x70, 0x74, 0x6f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x79, 0x2f, 0x76, 0x32, 0x2f, 0x73, 0x63,
	0x61, 0x6e, 0x6f, 0x73, 0x73, 0x2d, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x67, 0x72, 0x61, 0x70,
	0x68, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x73, 0x63, 0x61, 0x6e, 0x6f, 0x73,
	0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x67, 0x72, 0x61, 0x70,
	0x68, 0x79, 0x2e, 0x76, 0x32, 0x1a, 0x2a, 0x73, 0x63, 0x61, 0x6e, 0x6f, 0x73, 0x73, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x32, 0x2f, 0x73, 0x63, 0x61,
	0x6e, 0x6f, 0x73, 0x73, 0x2d, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xa2, 0x01, 0x0a, 0x13, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x67, 0x72, 0x61, 0x70,
	0x68, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x4c, 0x0a, 0x05, 0x70, 0x75, 0x72,
	0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x73, 0x63, 0x61, 0x6e, 0x6f,
	0x73, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x67, 0x72, 0x61,
	0x70, 0x68, 0x79, 0x2e, 0x76, 0x32, 0x2e, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x67, 0x72, 0x61,
	0x70, 0x68, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x50, 0x75, 0x72, 0x6c, 0x73,
	0x52, 0x05, 0x70, 0x75, 0x72, 0x6c, 0x73, 0x1a, 0x3d, 0x0a, 0x05, 0x50, 0x75, 0x72, 0x6c, 0x73,
	0x12, 0x12, 0x0a, 0x04, 0x70, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x70, 0x75, 0x72, 0x6c, 0x12, 0x20, 0x0a, 0x0b, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x71, 0x75, 0x69,
	0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0xfb, 0x02, 0x0a, 0x14, 0x43, 0x72, 0x79, 0x70, 0x74,
	0x6f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x4d, 0x0a, 0x05, 0x70, 0x75, 0x72, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x37,
	0x2e, 0x73, 0x63, 0x61, 0x6e, 0x6f, 0x73, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x72, 0x79,
	0x70, 0x74, 0x6f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x79, 0x2e, 0x76, 0x32, 0x2e, 0x43, 0x72, 0x79,
	0x70, 0x74, 0x6f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x2e, 0x50, 0x75, 0x72, 0x6c, 0x73, 0x52, 0x05, 0x70, 0x75, 0x72, 0x6c, 0x73, 0x12, 0x3d,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25,
	0x2e, 0x73, 0x63, 0x61, 0x6e, 0x6f, 0x73, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x32, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x1a, 0x5c, 0x0a,
	0x0a, 0x41, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x61,
	0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x61, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x74, 0x72,
	0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x74, 0x72,
	0x65, 0x6e, 0x67, 0x74, 0x68, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x1a, 0x77, 0x0a, 0x05, 0x50,
	0x75, 0x72, 0x6c, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x70, 0x75, 0x72, 0x6c, 0x12, 0x5a, 0x0a, 0x09, 0x61, 0x6c, 0x67, 0x6f,
	0x72, 0x69, 0x74, 0x68, 0x6d, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3c, 0x2e, 0x73, 0x63,
	0x61, 0x6e, 0x6f, 0x73, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f,
	0x67, 0x72, 0x61, 0x70, 0x68, 0x79, 0x2e, 0x76, 0x32, 0x2e, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f,
	0x67, 0x72, 0x61, 0x70, 0x68, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x41,
	0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x73, 0x52, 0x09, 0x61, 0x6c, 0x67, 0x6f, 0x72,
	0x69, 0x74, 0x68, 0x6d, 0x32, 0xd5, 0x01, 0x0a, 0x0c, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x67,
	0x72, 0x61, 0x70, 0x68, 0x79, 0x12, 0x4f, 0x0a, 0x04, 0x45, 0x63, 0x68, 0x6f, 0x12, 0x22, 0x2e,
	0x73, 0x63, 0x61, 0x6e, 0x6f, 0x73, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x76, 0x32, 0x2e, 0x45, 0x63, 0x68, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x23, 0x2e, 0x73, 0x63, 0x61, 0x6e, 0x6f, 0x73, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x32, 0x2e, 0x45, 0x63, 0x68, 0x6f, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x74, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x67,
	0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x73, 0x12, 0x30, 0x2e, 0x73, 0x63, 0x61, 0x6e, 0x6f, 0x73,
	0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x67, 0x72, 0x61, 0x70,
	0x68, 0x79, 0x2e, 0x76, 0x32, 0x2e, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x67, 0x72, 0x61, 0x70,
	0x68, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x31, 0x2e, 0x73, 0x63, 0x61, 0x6e,
	0x6f, 0x73, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x67, 0x72,
	0x61, 0x70, 0x68, 0x79, 0x2e, 0x76, 0x32, 0x2e, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x67, 0x72,
	0x61, 0x70, 0x68, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x3b, 0x5a, 0x39,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x63, 0x61, 0x6e, 0x6f,
	0x73, 0x73, 0x2f, 0x70, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x72, 0x79, 0x70,
	0x74, 0x6f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x79, 0x76, 0x32, 0x3b, 0x63, 0x72, 0x79, 0x70, 0x74,
	0x6f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x79, 0x76, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_scanoss_api_cryptography_v2_scanoss_cryptography_proto_rawDescOnce sync.Once
	file_scanoss_api_cryptography_v2_scanoss_cryptography_proto_rawDescData = file_scanoss_api_cryptography_v2_scanoss_cryptography_proto_rawDesc
)

func file_scanoss_api_cryptography_v2_scanoss_cryptography_proto_rawDescGZIP() []byte {
	file_scanoss_api_cryptography_v2_scanoss_cryptography_proto_rawDescOnce.Do(func() {
		file_scanoss_api_cryptography_v2_scanoss_cryptography_proto_rawDescData = protoimpl.X.CompressGZIP(file_scanoss_api_cryptography_v2_scanoss_cryptography_proto_rawDescData)
	})
	return file_scanoss_api_cryptography_v2_scanoss_cryptography_proto_rawDescData
}

var file_scanoss_api_cryptography_v2_scanoss_cryptography_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_scanoss_api_cryptography_v2_scanoss_cryptography_proto_goTypes = []interface{}{
	(*CryptographyRequest)(nil),             // 0: scanoss.api.cryptography.v2.CryptographyRequest
	(*CryptographyResponse)(nil),            // 1: scanoss.api.cryptography.v2.CryptographyResponse
	(*CryptographyRequest_Purls)(nil),       // 2: scanoss.api.cryptography.v2.CryptographyRequest.Purls
	(*CryptographyResponse_Algorithms)(nil), // 3: scanoss.api.cryptography.v2.CryptographyResponse.Algorithms
	(*CryptographyResponse_Purls)(nil),      // 4: scanoss.api.cryptography.v2.CryptographyResponse.Purls
	(*commonv2.StatusResponse)(nil),         // 5: scanoss.api.common.v2.StatusResponse
	(*commonv2.EchoRequest)(nil),            // 6: scanoss.api.common.v2.EchoRequest
	(*commonv2.EchoResponse)(nil),           // 7: scanoss.api.common.v2.EchoResponse
}
var file_scanoss_api_cryptography_v2_scanoss_cryptography_proto_depIdxs = []int32{
	2, // 0: scanoss.api.cryptography.v2.CryptographyRequest.purls:type_name -> scanoss.api.cryptography.v2.CryptographyRequest.Purls
	4, // 1: scanoss.api.cryptography.v2.CryptographyResponse.purls:type_name -> scanoss.api.cryptography.v2.CryptographyResponse.Purls
	5, // 2: scanoss.api.cryptography.v2.CryptographyResponse.status:type_name -> scanoss.api.common.v2.StatusResponse
	3, // 3: scanoss.api.cryptography.v2.CryptographyResponse.Purls.algorithm:type_name -> scanoss.api.cryptography.v2.CryptographyResponse.Algorithms
	6, // 4: scanoss.api.cryptography.v2.Cryptography.Echo:input_type -> scanoss.api.common.v2.EchoRequest
	0, // 5: scanoss.api.cryptography.v2.Cryptography.GetAlgorithms:input_type -> scanoss.api.cryptography.v2.CryptographyRequest
	7, // 6: scanoss.api.cryptography.v2.Cryptography.Echo:output_type -> scanoss.api.common.v2.EchoResponse
	1, // 7: scanoss.api.cryptography.v2.Cryptography.GetAlgorithms:output_type -> scanoss.api.cryptography.v2.CryptographyResponse
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_scanoss_api_cryptography_v2_scanoss_cryptography_proto_init() }
func file_scanoss_api_cryptography_v2_scanoss_cryptography_proto_init() {
	if File_scanoss_api_cryptography_v2_scanoss_cryptography_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_scanoss_api_cryptography_v2_scanoss_cryptography_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CryptographyRequest); i {
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
		file_scanoss_api_cryptography_v2_scanoss_cryptography_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CryptographyResponse); i {
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
		file_scanoss_api_cryptography_v2_scanoss_cryptography_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CryptographyRequest_Purls); i {
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
		file_scanoss_api_cryptography_v2_scanoss_cryptography_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CryptographyResponse_Algorithms); i {
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
		file_scanoss_api_cryptography_v2_scanoss_cryptography_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CryptographyResponse_Purls); i {
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
			RawDescriptor: file_scanoss_api_cryptography_v2_scanoss_cryptography_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_scanoss_api_cryptography_v2_scanoss_cryptography_proto_goTypes,
		DependencyIndexes: file_scanoss_api_cryptography_v2_scanoss_cryptography_proto_depIdxs,
		MessageInfos:      file_scanoss_api_cryptography_v2_scanoss_cryptography_proto_msgTypes,
	}.Build()
	File_scanoss_api_cryptography_v2_scanoss_cryptography_proto = out.File
	file_scanoss_api_cryptography_v2_scanoss_cryptography_proto_rawDesc = nil
	file_scanoss_api_cryptography_v2_scanoss_cryptography_proto_goTypes = nil
	file_scanoss_api_cryptography_v2_scanoss_cryptography_proto_depIdxs = nil
}
