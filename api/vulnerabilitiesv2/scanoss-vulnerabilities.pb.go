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
// Vulnerability definition details
//*

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.21.12
// source: scanoss/api/vulnerabilities/v2/scanoss-vulnerabilities.proto

package vulnerabilitiesv2

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
// Vulnerability request data (JSON payload)
type VulnerabilityRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// JSON vulnerability request purls
	Purls []*VulnerabilityRequest_Purls `protobuf:"bytes,1,rep,name=purls,proto3" json:"purls,omitempty"`
}

func (x *VulnerabilityRequest) Reset() {
	*x = VulnerabilityRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scanoss_api_vulnerabilities_v2_scanoss_vulnerabilities_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VulnerabilityRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VulnerabilityRequest) ProtoMessage() {}

func (x *VulnerabilityRequest) ProtoReflect() protoreflect.Message {
	mi := &file_scanoss_api_vulnerabilities_v2_scanoss_vulnerabilities_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VulnerabilityRequest.ProtoReflect.Descriptor instead.
func (*VulnerabilityRequest) Descriptor() ([]byte, []int) {
	return file_scanoss_api_vulnerabilities_v2_scanoss_vulnerabilities_proto_rawDescGZIP(), []int{0}
}

func (x *VulnerabilityRequest) GetPurls() []*VulnerabilityRequest_Purls {
	if x != nil {
		return x.Purls
	}
	return nil
}

//
// Vulnerability CPE response data (JSON payload)
type CpeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Purl/CPE details
	Purls []*CpeResponse_Purls `protobuf:"bytes,1,rep,name=purls,proto3" json:"purls,omitempty"`
	// Response status
	Status *commonv2.StatusResponse `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *CpeResponse) Reset() {
	*x = CpeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scanoss_api_vulnerabilities_v2_scanoss_vulnerabilities_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CpeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CpeResponse) ProtoMessage() {}

func (x *CpeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_scanoss_api_vulnerabilities_v2_scanoss_vulnerabilities_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CpeResponse.ProtoReflect.Descriptor instead.
func (*CpeResponse) Descriptor() ([]byte, []int) {
	return file_scanoss_api_vulnerabilities_v2_scanoss_vulnerabilities_proto_rawDescGZIP(), []int{1}
}

func (x *CpeResponse) GetPurls() []*CpeResponse_Purls {
	if x != nil {
		return x.Purls
	}
	return nil
}

func (x *CpeResponse) GetStatus() *commonv2.StatusResponse {
	if x != nil {
		return x.Status
	}
	return nil
}

//
// Vulnerability response data (JSON payload)
type VulnerabilityResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Vulnerability response details
	Purls []*VulnerabilityResponse_Purls `protobuf:"bytes,1,rep,name=purls,proto3" json:"purls,omitempty"`
	// Response status
	Status *commonv2.StatusResponse `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *VulnerabilityResponse) Reset() {
	*x = VulnerabilityResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scanoss_api_vulnerabilities_v2_scanoss_vulnerabilities_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VulnerabilityResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VulnerabilityResponse) ProtoMessage() {}

func (x *VulnerabilityResponse) ProtoReflect() protoreflect.Message {
	mi := &file_scanoss_api_vulnerabilities_v2_scanoss_vulnerabilities_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VulnerabilityResponse.ProtoReflect.Descriptor instead.
func (*VulnerabilityResponse) Descriptor() ([]byte, []int) {
	return file_scanoss_api_vulnerabilities_v2_scanoss_vulnerabilities_proto_rawDescGZIP(), []int{2}
}

func (x *VulnerabilityResponse) GetPurls() []*VulnerabilityResponse_Purls {
	if x != nil {
		return x.Purls
	}
	return nil
}

func (x *VulnerabilityResponse) GetStatus() *commonv2.StatusResponse {
	if x != nil {
		return x.Status
	}
	return nil
}

type VulnerabilityRequest_Purls struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Purl        string `protobuf:"bytes,1,opt,name=purl,proto3" json:"purl,omitempty"`
	Requirement string `protobuf:"bytes,2,opt,name=requirement,proto3" json:"requirement,omitempty"`
}

func (x *VulnerabilityRequest_Purls) Reset() {
	*x = VulnerabilityRequest_Purls{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scanoss_api_vulnerabilities_v2_scanoss_vulnerabilities_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VulnerabilityRequest_Purls) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VulnerabilityRequest_Purls) ProtoMessage() {}

func (x *VulnerabilityRequest_Purls) ProtoReflect() protoreflect.Message {
	mi := &file_scanoss_api_vulnerabilities_v2_scanoss_vulnerabilities_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VulnerabilityRequest_Purls.ProtoReflect.Descriptor instead.
func (*VulnerabilityRequest_Purls) Descriptor() ([]byte, []int) {
	return file_scanoss_api_vulnerabilities_v2_scanoss_vulnerabilities_proto_rawDescGZIP(), []int{0, 0}
}

func (x *VulnerabilityRequest_Purls) GetPurl() string {
	if x != nil {
		return x.Purl
	}
	return ""
}

func (x *VulnerabilityRequest_Purls) GetRequirement() string {
	if x != nil {
		return x.Requirement
	}
	return ""
}

type CpeResponse_Purls struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Purl string   `protobuf:"bytes,1,opt,name=purl,proto3" json:"purl,omitempty"`
	Cpes []string `protobuf:"bytes,2,rep,name=cpes,proto3" json:"cpes,omitempty"`
}

func (x *CpeResponse_Purls) Reset() {
	*x = CpeResponse_Purls{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scanoss_api_vulnerabilities_v2_scanoss_vulnerabilities_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CpeResponse_Purls) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CpeResponse_Purls) ProtoMessage() {}

func (x *CpeResponse_Purls) ProtoReflect() protoreflect.Message {
	mi := &file_scanoss_api_vulnerabilities_v2_scanoss_vulnerabilities_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CpeResponse_Purls.ProtoReflect.Descriptor instead.
func (*CpeResponse_Purls) Descriptor() ([]byte, []int) {
	return file_scanoss_api_vulnerabilities_v2_scanoss_vulnerabilities_proto_rawDescGZIP(), []int{1, 0}
}

func (x *CpeResponse_Purls) GetPurl() string {
	if x != nil {
		return x.Purl
	}
	return ""
}

func (x *CpeResponse_Purls) GetCpes() []string {
	if x != nil {
		return x.Cpes
	}
	return nil
}

type VulnerabilityResponse_Vulnerabilities struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Cve       string `protobuf:"bytes,2,opt,name=cve,proto3" json:"cve,omitempty"`
	Url       string `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
	Summary   string `protobuf:"bytes,4,opt,name=summary,proto3" json:"summary,omitempty"`
	Severity  string `protobuf:"bytes,5,opt,name=severity,proto3" json:"severity,omitempty"`
	Published string `protobuf:"bytes,6,opt,name=published,proto3" json:"published,omitempty"`
	Modified  string `protobuf:"bytes,7,opt,name=modified,proto3" json:"modified,omitempty"`
	Source    string `protobuf:"bytes,8,opt,name=source,proto3" json:"source,omitempty"`
}

func (x *VulnerabilityResponse_Vulnerabilities) Reset() {
	*x = VulnerabilityResponse_Vulnerabilities{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scanoss_api_vulnerabilities_v2_scanoss_vulnerabilities_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VulnerabilityResponse_Vulnerabilities) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VulnerabilityResponse_Vulnerabilities) ProtoMessage() {}

func (x *VulnerabilityResponse_Vulnerabilities) ProtoReflect() protoreflect.Message {
	mi := &file_scanoss_api_vulnerabilities_v2_scanoss_vulnerabilities_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VulnerabilityResponse_Vulnerabilities.ProtoReflect.Descriptor instead.
func (*VulnerabilityResponse_Vulnerabilities) Descriptor() ([]byte, []int) {
	return file_scanoss_api_vulnerabilities_v2_scanoss_vulnerabilities_proto_rawDescGZIP(), []int{2, 0}
}

func (x *VulnerabilityResponse_Vulnerabilities) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *VulnerabilityResponse_Vulnerabilities) GetCve() string {
	if x != nil {
		return x.Cve
	}
	return ""
}

func (x *VulnerabilityResponse_Vulnerabilities) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *VulnerabilityResponse_Vulnerabilities) GetSummary() string {
	if x != nil {
		return x.Summary
	}
	return ""
}

func (x *VulnerabilityResponse_Vulnerabilities) GetSeverity() string {
	if x != nil {
		return x.Severity
	}
	return ""
}

func (x *VulnerabilityResponse_Vulnerabilities) GetPublished() string {
	if x != nil {
		return x.Published
	}
	return ""
}

func (x *VulnerabilityResponse_Vulnerabilities) GetModified() string {
	if x != nil {
		return x.Modified
	}
	return ""
}

func (x *VulnerabilityResponse_Vulnerabilities) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

type VulnerabilityResponse_Purls struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Purl            string                                   `protobuf:"bytes,1,opt,name=purl,proto3" json:"purl,omitempty"`
	Vulnerabilities []*VulnerabilityResponse_Vulnerabilities `protobuf:"bytes,2,rep,name=vulnerabilities,proto3" json:"vulnerabilities,omitempty"`
}

func (x *VulnerabilityResponse_Purls) Reset() {
	*x = VulnerabilityResponse_Purls{}
	if protoimpl.UnsafeEnabled {
		mi := &file_scanoss_api_vulnerabilities_v2_scanoss_vulnerabilities_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VulnerabilityResponse_Purls) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VulnerabilityResponse_Purls) ProtoMessage() {}

func (x *VulnerabilityResponse_Purls) ProtoReflect() protoreflect.Message {
	mi := &file_scanoss_api_vulnerabilities_v2_scanoss_vulnerabilities_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VulnerabilityResponse_Purls.ProtoReflect.Descriptor instead.
func (*VulnerabilityResponse_Purls) Descriptor() ([]byte, []int) {
	return file_scanoss_api_vulnerabilities_v2_scanoss_vulnerabilities_proto_rawDescGZIP(), []int{2, 1}
}

func (x *VulnerabilityResponse_Purls) GetPurl() string {
	if x != nil {
		return x.Purl
	}
	return ""
}

func (x *VulnerabilityResponse_Purls) GetVulnerabilities() []*VulnerabilityResponse_Vulnerabilities {
	if x != nil {
		return x.Vulnerabilities
	}
	return nil
}

var File_scanoss_api_vulnerabilities_v2_scanoss_vulnerabilities_proto protoreflect.FileDescriptor

var file_scanoss_api_vulnerabilities_v2_scanoss_vulnerabilities_proto_rawDesc = []byte{
	0x0a, 0x3c, 0x73, 0x63, 0x61, 0x6e, 0x6f, 0x73, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x75,
	0x6c, 0x6e, 0x65, 0x72, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2f, 0x76, 0x32,
	0x2f, 0x73, 0x63, 0x61, 0x6e, 0x6f, 0x73, 0x73, 0x2d, 0x76, 0x75, 0x6c, 0x6e, 0x65, 0x72, 0x61,
	0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e,
	0x73, 0x63, 0x61, 0x6e, 0x6f, 0x73, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x75, 0x6c, 0x6e,
	0x65, 0x72, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x32, 0x1a, 0x2a,
	0x73, 0x63, 0x61, 0x6e, 0x6f, 0x73, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2f, 0x76, 0x32, 0x2f, 0x73, 0x63, 0x61, 0x6e, 0x6f, 0x73, 0x73, 0x2d, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa7, 0x01, 0x0a, 0x14, 0x56,
	0x75, 0x6c, 0x6e, 0x65, 0x72, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x50, 0x0a, 0x05, 0x70, 0x75, 0x72, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x3a, 0x2e, 0x73, 0x63, 0x61, 0x6e, 0x6f, 0x73, 0x73, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x75, 0x6c, 0x6e, 0x65, 0x72, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73,
	0x2e, 0x76, 0x32, 0x2e, 0x56, 0x75, 0x6c, 0x6e, 0x65, 0x72, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x50, 0x75, 0x72, 0x6c, 0x73, 0x52, 0x05,
	0x70, 0x75, 0x72, 0x6c, 0x73, 0x1a, 0x3d, 0x0a, 0x05, 0x50, 0x75, 0x72, 0x6c, 0x73, 0x12, 0x12,
	0x0a, 0x04, 0x70, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x75,
	0x72, 0x6c, 0x12, 0x20, 0x0a, 0x0b, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x22, 0xc6, 0x01, 0x0a, 0x0b, 0x43, 0x70, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x47, 0x0a, 0x05, 0x70, 0x75, 0x72, 0x6c, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x73, 0x63, 0x61, 0x6e, 0x6f, 0x73, 0x73, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x75, 0x6c, 0x6e, 0x65, 0x72, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65,
	0x73, 0x2e, 0x76, 0x32, 0x2e, 0x43, 0x70, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x2e, 0x50, 0x75, 0x72, 0x6c, 0x73, 0x52, 0x05, 0x70, 0x75, 0x72, 0x6c, 0x73, 0x12, 0x3d, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e,
	0x73, 0x63, 0x61, 0x6e, 0x6f, 0x73, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x76, 0x32, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x1a, 0x2f, 0x0a, 0x05,
	0x50, 0x75, 0x72, 0x6c, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x75, 0x72, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x70, 0x65,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x63, 0x70, 0x65, 0x73, 0x22, 0x88, 0x04,
	0x0a, 0x15, 0x56, 0x75, 0x6c, 0x6e, 0x65, 0x72, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x51, 0x0a, 0x05, 0x70, 0x75, 0x72, 0x6c, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3b, 0x2e, 0x73, 0x63, 0x61, 0x6e, 0x6f, 0x73, 0x73,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x75, 0x6c, 0x6e, 0x65, 0x72, 0x61, 0x62, 0x69, 0x6c, 0x69,
	0x74, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x56, 0x75, 0x6c, 0x6e, 0x65, 0x72, 0x61, 0x62,
	0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x50, 0x75,
	0x72, 0x6c, 0x73, 0x52, 0x05, 0x70, 0x75, 0x72, 0x6c, 0x73, 0x12, 0x3d, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x73, 0x63, 0x61,
	0x6e, 0x6f, 0x73, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x76, 0x32, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x1a, 0xcd, 0x01, 0x0a, 0x0f, 0x56, 0x75,
	0x6c, 0x6e, 0x65, 0x72, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a,
	0x03, 0x63, 0x76, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x63, 0x76, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72,
	0x6c, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x73,
	0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73,
	0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69,
	0x73, 0x68, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x75, 0x62, 0x6c,
	0x69, 0x73, 0x68, 0x65, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65,
	0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x1a, 0x8c, 0x01, 0x0a, 0x05, 0x50, 0x75,
	0x72, 0x6c, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x70, 0x75, 0x72, 0x6c, 0x12, 0x6f, 0x0a, 0x0f, 0x76, 0x75, 0x6c, 0x6e, 0x65,
	0x72, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x45, 0x2e, 0x73, 0x63, 0x61, 0x6e, 0x6f, 0x73, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x75, 0x6c, 0x6e, 0x65, 0x72, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x76,
	0x32, 0x2e, 0x56, 0x75, 0x6c, 0x6e, 0x65, 0x72, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x56, 0x75, 0x6c, 0x6e, 0x65, 0x72, 0x61, 0x62,
	0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x52, 0x0f, 0x76, 0x75, 0x6c, 0x6e, 0x65, 0x72, 0x61,
	0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x32, 0xd4, 0x02, 0x0a, 0x0f, 0x56, 0x75, 0x6c,
	0x6e, 0x65, 0x72, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x12, 0x4f, 0x0a, 0x04,
	0x45, 0x63, 0x68, 0x6f, 0x12, 0x22, 0x2e, 0x73, 0x63, 0x61, 0x6e, 0x6f, 0x73, 0x73, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x32, 0x2e, 0x45, 0x63, 0x68,
	0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x73, 0x63, 0x61, 0x6e, 0x6f,
	0x73, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x32,
	0x2e, 0x45, 0x63, 0x68, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x6c, 0x0a,
	0x07, 0x47, 0x65, 0x74, 0x43, 0x70, 0x65, 0x73, 0x12, 0x34, 0x2e, 0x73, 0x63, 0x61, 0x6e, 0x6f,
	0x73, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x75, 0x6c, 0x6e, 0x65, 0x72, 0x61, 0x62, 0x69,
	0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x56, 0x75, 0x6c, 0x6e, 0x65, 0x72,
	0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b,
	0x2e, 0x73, 0x63, 0x61, 0x6e, 0x6f, 0x73, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x75, 0x6c,
	0x6e, 0x65, 0x72, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x32, 0x2e,
	0x43, 0x70, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x81, 0x01, 0x0a, 0x12,
	0x47, 0x65, 0x74, 0x56, 0x75, 0x6c, 0x6e, 0x65, 0x72, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69,
	0x65, 0x73, 0x12, 0x34, 0x2e, 0x73, 0x63, 0x61, 0x6e, 0x6f, 0x73, 0x73, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x75, 0x6c, 0x6e, 0x65, 0x72, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73,
	0x2e, 0x76, 0x32, 0x2e, 0x56, 0x75, 0x6c, 0x6e, 0x65, 0x72, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x35, 0x2e, 0x73, 0x63, 0x61, 0x6e, 0x6f,
	0x73, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x75, 0x6c, 0x6e, 0x65, 0x72, 0x61, 0x62, 0x69,
	0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x56, 0x75, 0x6c, 0x6e, 0x65, 0x72,
	0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x41, 0x5a, 0x3f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x63,
	0x61, 0x6e, 0x6f, 0x73, 0x73, 0x2f, 0x70, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x75, 0x6c, 0x6e, 0x65, 0x72, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x76, 0x32,
	0x3b, 0x76, 0x75, 0x6c, 0x6e, 0x65, 0x72, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73,
	0x76, 0x32, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_scanoss_api_vulnerabilities_v2_scanoss_vulnerabilities_proto_rawDescOnce sync.Once
	file_scanoss_api_vulnerabilities_v2_scanoss_vulnerabilities_proto_rawDescData = file_scanoss_api_vulnerabilities_v2_scanoss_vulnerabilities_proto_rawDesc
)

func file_scanoss_api_vulnerabilities_v2_scanoss_vulnerabilities_proto_rawDescGZIP() []byte {
	file_scanoss_api_vulnerabilities_v2_scanoss_vulnerabilities_proto_rawDescOnce.Do(func() {
		file_scanoss_api_vulnerabilities_v2_scanoss_vulnerabilities_proto_rawDescData = protoimpl.X.CompressGZIP(file_scanoss_api_vulnerabilities_v2_scanoss_vulnerabilities_proto_rawDescData)
	})
	return file_scanoss_api_vulnerabilities_v2_scanoss_vulnerabilities_proto_rawDescData
}

var file_scanoss_api_vulnerabilities_v2_scanoss_vulnerabilities_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_scanoss_api_vulnerabilities_v2_scanoss_vulnerabilities_proto_goTypes = []interface{}{
	(*VulnerabilityRequest)(nil),                  // 0: scanoss.api.vulnerabilities.v2.VulnerabilityRequest
	(*CpeResponse)(nil),                           // 1: scanoss.api.vulnerabilities.v2.CpeResponse
	(*VulnerabilityResponse)(nil),                 // 2: scanoss.api.vulnerabilities.v2.VulnerabilityResponse
	(*VulnerabilityRequest_Purls)(nil),            // 3: scanoss.api.vulnerabilities.v2.VulnerabilityRequest.Purls
	(*CpeResponse_Purls)(nil),                     // 4: scanoss.api.vulnerabilities.v2.CpeResponse.Purls
	(*VulnerabilityResponse_Vulnerabilities)(nil), // 5: scanoss.api.vulnerabilities.v2.VulnerabilityResponse.Vulnerabilities
	(*VulnerabilityResponse_Purls)(nil),           // 6: scanoss.api.vulnerabilities.v2.VulnerabilityResponse.Purls
	(*commonv2.StatusResponse)(nil),               // 7: scanoss.api.common.v2.StatusResponse
	(*commonv2.EchoRequest)(nil),                  // 8: scanoss.api.common.v2.EchoRequest
	(*commonv2.EchoResponse)(nil),                 // 9: scanoss.api.common.v2.EchoResponse
}
var file_scanoss_api_vulnerabilities_v2_scanoss_vulnerabilities_proto_depIdxs = []int32{
	3, // 0: scanoss.api.vulnerabilities.v2.VulnerabilityRequest.purls:type_name -> scanoss.api.vulnerabilities.v2.VulnerabilityRequest.Purls
	4, // 1: scanoss.api.vulnerabilities.v2.CpeResponse.purls:type_name -> scanoss.api.vulnerabilities.v2.CpeResponse.Purls
	7, // 2: scanoss.api.vulnerabilities.v2.CpeResponse.status:type_name -> scanoss.api.common.v2.StatusResponse
	6, // 3: scanoss.api.vulnerabilities.v2.VulnerabilityResponse.purls:type_name -> scanoss.api.vulnerabilities.v2.VulnerabilityResponse.Purls
	7, // 4: scanoss.api.vulnerabilities.v2.VulnerabilityResponse.status:type_name -> scanoss.api.common.v2.StatusResponse
	5, // 5: scanoss.api.vulnerabilities.v2.VulnerabilityResponse.Purls.vulnerabilities:type_name -> scanoss.api.vulnerabilities.v2.VulnerabilityResponse.Vulnerabilities
	8, // 6: scanoss.api.vulnerabilities.v2.Vulnerabilities.Echo:input_type -> scanoss.api.common.v2.EchoRequest
	0, // 7: scanoss.api.vulnerabilities.v2.Vulnerabilities.GetCpes:input_type -> scanoss.api.vulnerabilities.v2.VulnerabilityRequest
	0, // 8: scanoss.api.vulnerabilities.v2.Vulnerabilities.GetVulnerabilities:input_type -> scanoss.api.vulnerabilities.v2.VulnerabilityRequest
	9, // 9: scanoss.api.vulnerabilities.v2.Vulnerabilities.Echo:output_type -> scanoss.api.common.v2.EchoResponse
	1, // 10: scanoss.api.vulnerabilities.v2.Vulnerabilities.GetCpes:output_type -> scanoss.api.vulnerabilities.v2.CpeResponse
	2, // 11: scanoss.api.vulnerabilities.v2.Vulnerabilities.GetVulnerabilities:output_type -> scanoss.api.vulnerabilities.v2.VulnerabilityResponse
	9, // [9:12] is the sub-list for method output_type
	6, // [6:9] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_scanoss_api_vulnerabilities_v2_scanoss_vulnerabilities_proto_init() }
func file_scanoss_api_vulnerabilities_v2_scanoss_vulnerabilities_proto_init() {
	if File_scanoss_api_vulnerabilities_v2_scanoss_vulnerabilities_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_scanoss_api_vulnerabilities_v2_scanoss_vulnerabilities_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VulnerabilityRequest); i {
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
		file_scanoss_api_vulnerabilities_v2_scanoss_vulnerabilities_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CpeResponse); i {
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
		file_scanoss_api_vulnerabilities_v2_scanoss_vulnerabilities_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VulnerabilityResponse); i {
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
		file_scanoss_api_vulnerabilities_v2_scanoss_vulnerabilities_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VulnerabilityRequest_Purls); i {
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
		file_scanoss_api_vulnerabilities_v2_scanoss_vulnerabilities_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CpeResponse_Purls); i {
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
		file_scanoss_api_vulnerabilities_v2_scanoss_vulnerabilities_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VulnerabilityResponse_Vulnerabilities); i {
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
		file_scanoss_api_vulnerabilities_v2_scanoss_vulnerabilities_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VulnerabilityResponse_Purls); i {
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
			RawDescriptor: file_scanoss_api_vulnerabilities_v2_scanoss_vulnerabilities_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_scanoss_api_vulnerabilities_v2_scanoss_vulnerabilities_proto_goTypes,
		DependencyIndexes: file_scanoss_api_vulnerabilities_v2_scanoss_vulnerabilities_proto_depIdxs,
		MessageInfos:      file_scanoss_api_vulnerabilities_v2_scanoss_vulnerabilities_proto_msgTypes,
	}.Build()
	File_scanoss_api_vulnerabilities_v2_scanoss_vulnerabilities_proto = out.File
	file_scanoss_api_vulnerabilities_v2_scanoss_vulnerabilities_proto_rawDesc = nil
	file_scanoss_api_vulnerabilities_v2_scanoss_vulnerabilities_proto_goTypes = nil
	file_scanoss_api_vulnerabilities_v2_scanoss_vulnerabilities_proto_depIdxs = nil
}
