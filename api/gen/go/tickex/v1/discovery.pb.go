// Copyright 2025 The Tickex Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.21.12
// source: tickex/v1/discovery.proto

package tickex

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
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

// Detailed service information
type RegisterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Service      *Service      `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	ServiceCheck *ServiceCheck `protobuf:"bytes,2,opt,name=service_check,json=serviceCheck,proto3" json:"service_check,omitempty"`
}

func (x *RegisterRequest) Reset() {
	*x = RegisterRequest{}
	mi := &file_tickex_v1_discovery_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RegisterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterRequest) ProtoMessage() {}

func (x *RegisterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tickex_v1_discovery_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterRequest.ProtoReflect.Descriptor instead.
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return file_tickex_v1_discovery_proto_rawDescGZIP(), []int{0}
}

func (x *RegisterRequest) GetService() *Service {
	if x != nil {
		return x.Service
	}
	return nil
}

func (x *RegisterRequest) GetServiceCheck() *ServiceCheck {
	if x != nil {
		return x.ServiceCheck
	}
	return nil
}

type HeartbeatRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *HeartbeatRequest) Reset() {
	*x = HeartbeatRequest{}
	mi := &file_tickex_v1_discovery_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *HeartbeatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeartbeatRequest) ProtoMessage() {}

func (x *HeartbeatRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tickex_v1_discovery_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeartbeatRequest.ProtoReflect.Descriptor instead.
func (*HeartbeatRequest) Descriptor() ([]byte, []int) {
	return file_tickex_v1_discovery_proto_rawDescGZIP(), []int{1}
}

func (x *HeartbeatRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ServiceCheck struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeregisterCriticalServiceAfter string `protobuf:"bytes,1,opt,name=deregister_critical_service_after,json=deregisterCriticalServiceAfter,proto3" json:"deregister_critical_service_after,omitempty"`
	Ttl                            string `protobuf:"bytes,2,opt,name=ttl,proto3" json:"ttl,omitempty"`
	TlsSkipVerify                  bool   `protobuf:"varint,3,opt,name=tls_skip_verify,json=tlsSkipVerify,proto3" json:"tls_skip_verify,omitempty"`
}

func (x *ServiceCheck) Reset() {
	*x = ServiceCheck{}
	mi := &file_tickex_v1_discovery_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ServiceCheck) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceCheck) ProtoMessage() {}

func (x *ServiceCheck) ProtoReflect() protoreflect.Message {
	mi := &file_tickex_v1_discovery_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceCheck.ProtoReflect.Descriptor instead.
func (*ServiceCheck) Descriptor() ([]byte, []int) {
	return file_tickex_v1_discovery_proto_rawDescGZIP(), []int{2}
}

func (x *ServiceCheck) GetDeregisterCriticalServiceAfter() string {
	if x != nil {
		return x.DeregisterCriticalServiceAfter
	}
	return ""
}

func (x *ServiceCheck) GetTtl() string {
	if x != nil {
		return x.Ttl
	}
	return ""
}

func (x *ServiceCheck) GetTlsSkipVerify() bool {
	if x != nil {
		return x.TlsSkipVerify
	}
	return false
}

// Request to find a service
type DiscoverRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"` // Name of the service to discover
}

func (x *DiscoverRequest) Reset() {
	*x = DiscoverRequest{}
	mi := &file_tickex_v1_discovery_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DiscoverRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DiscoverRequest) ProtoMessage() {}

func (x *DiscoverRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tickex_v1_discovery_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DiscoverRequest.ProtoReflect.Descriptor instead.
func (*DiscoverRequest) Descriptor() ([]byte, []int) {
	return file_tickex_v1_discovery_proto_rawDescGZIP(), []int{3}
}

func (x *DiscoverRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// Response from the discovery process
type DiscoverResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Services []*Service `protobuf:"bytes,1,rep,name=services,proto3" json:"services,omitempty"` // List of service instances
}

func (x *DiscoverResponse) Reset() {
	*x = DiscoverResponse{}
	mi := &file_tickex_v1_discovery_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DiscoverResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DiscoverResponse) ProtoMessage() {}

func (x *DiscoverResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tickex_v1_discovery_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DiscoverResponse.ProtoReflect.Descriptor instead.
func (*DiscoverResponse) Descriptor() ([]byte, []int) {
	return file_tickex_v1_discovery_proto_rawDescGZIP(), []int{4}
}

func (x *DiscoverResponse) GetServices() []*Service {
	if x != nil {
		return x.Services
	}
	return nil
}

// Service registration response
type RegisterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"` // Result message
}

func (x *RegisterResponse) Reset() {
	*x = RegisterResponse{}
	mi := &file_tickex_v1_discovery_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RegisterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterResponse) ProtoMessage() {}

func (x *RegisterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tickex_v1_discovery_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterResponse.ProtoReflect.Descriptor instead.
func (*RegisterResponse) Descriptor() ([]byte, []int) {
	return file_tickex_v1_discovery_proto_rawDescGZIP(), []int{5}
}

func (x *RegisterResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// Heartbeat response
type HeartbeatResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"` // Heartbeat status
}

func (x *HeartbeatResponse) Reset() {
	*x = HeartbeatResponse{}
	mi := &file_tickex_v1_discovery_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *HeartbeatResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HeartbeatResponse) ProtoMessage() {}

func (x *HeartbeatResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tickex_v1_discovery_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HeartbeatResponse.ProtoReflect.Descriptor instead.
func (*HeartbeatResponse) Descriptor() ([]byte, []int) {
	return file_tickex_v1_discovery_proto_rawDescGZIP(), []int{6}
}

func (x *HeartbeatResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_tickex_v1_discovery_proto protoreflect.FileDescriptor

var file_tickex_v1_discovery_proto_rawDesc = []byte{
	0x0a, 0x19, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x78, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x69, 0x73, 0x63,
	0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x74, 0x69, 0x63,
	0x6b, 0x65, 0x78, 0x2e, 0x76, 0x31, 0x1a, 0x17, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x78, 0x2f, 0x76,
	0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7d, 0x0a, 0x0f,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x2c, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x12, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3c, 0x0a,
	0x0d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x78, 0x2e, 0x76, 0x31,
	0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x0c, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x22, 0x2b, 0x0a, 0x10, 0x48,
	0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x17, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xba, 0x48, 0x04,
	0x72, 0x02, 0x10, 0x02, 0x52, 0x02, 0x69, 0x64, 0x22, 0x93, 0x01, 0x0a, 0x0c, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x49, 0x0a, 0x21, 0x64, 0x65, 0x72,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x63, 0x72, 0x69, 0x74, 0x69, 0x63, 0x61, 0x6c,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x61, 0x66, 0x74, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x1e, 0x64, 0x65, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x43, 0x72, 0x69, 0x74, 0x69, 0x63, 0x61, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41,
	0x66, 0x74, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x74, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x74, 0x74, 0x6c, 0x12, 0x26, 0x0a, 0x0f, 0x74, 0x6c, 0x73, 0x5f, 0x73, 0x6b,
	0x69, 0x70, 0x5f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0d, 0x74, 0x6c, 0x73, 0x53, 0x6b, 0x69, 0x70, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x22, 0x25,
	0x0a, 0x0f, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x42, 0x0a, 0x10, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x08, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x74, 0x69,
	0x63, 0x6b, 0x65, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52,
	0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x22, 0x2c, 0x0a, 0x10, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x2d, 0x0a, 0x11, 0x48, 0x65, 0x61, 0x72, 0x74,
	0x62, 0x65, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x32, 0xe4, 0x01, 0x0a, 0x10, 0x44, 0x69, 0x73, 0x63, 0x6f,
	0x76, 0x65, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x43, 0x0a, 0x08, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x1a, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x78,
	0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x78, 0x2e, 0x76, 0x31, 0x2e,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x43, 0x0a, 0x08, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x12, 0x1a, 0x2e, 0x74,
	0x69, 0x63, 0x6b, 0x65, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65,
	0x78, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x46, 0x0a, 0x09, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65,
	0x61, 0x74, 0x12, 0x1b, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x48,
	0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1c, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x65, 0x61, 0x72,
	0x74, 0x62, 0x65, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x38, 0x5a,
	0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x69, 0x63, 0x6b,
	0x65, 0x78, 0x76, 0x6e, 0x2f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x78, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x78, 0x2f, 0x76, 0x31,
	0x3b, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x78, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tickex_v1_discovery_proto_rawDescOnce sync.Once
	file_tickex_v1_discovery_proto_rawDescData = file_tickex_v1_discovery_proto_rawDesc
)

func file_tickex_v1_discovery_proto_rawDescGZIP() []byte {
	file_tickex_v1_discovery_proto_rawDescOnce.Do(func() {
		file_tickex_v1_discovery_proto_rawDescData = protoimpl.X.CompressGZIP(file_tickex_v1_discovery_proto_rawDescData)
	})
	return file_tickex_v1_discovery_proto_rawDescData
}

var file_tickex_v1_discovery_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_tickex_v1_discovery_proto_goTypes = []any{
	(*RegisterRequest)(nil),   // 0: tickex.v1.RegisterRequest
	(*HeartbeatRequest)(nil),  // 1: tickex.v1.HeartbeatRequest
	(*ServiceCheck)(nil),      // 2: tickex.v1.ServiceCheck
	(*DiscoverRequest)(nil),   // 3: tickex.v1.DiscoverRequest
	(*DiscoverResponse)(nil),  // 4: tickex.v1.DiscoverResponse
	(*RegisterResponse)(nil),  // 5: tickex.v1.RegisterResponse
	(*HeartbeatResponse)(nil), // 6: tickex.v1.HeartbeatResponse
	(*Service)(nil),           // 7: tickex.v1.Service
}
var file_tickex_v1_discovery_proto_depIdxs = []int32{
	7, // 0: tickex.v1.RegisterRequest.service:type_name -> tickex.v1.Service
	2, // 1: tickex.v1.RegisterRequest.service_check:type_name -> tickex.v1.ServiceCheck
	7, // 2: tickex.v1.DiscoverResponse.services:type_name -> tickex.v1.Service
	0, // 3: tickex.v1.DiscoveryService.Register:input_type -> tickex.v1.RegisterRequest
	3, // 4: tickex.v1.DiscoveryService.Discover:input_type -> tickex.v1.DiscoverRequest
	1, // 5: tickex.v1.DiscoveryService.Heartbeat:input_type -> tickex.v1.HeartbeatRequest
	5, // 6: tickex.v1.DiscoveryService.Register:output_type -> tickex.v1.RegisterResponse
	4, // 7: tickex.v1.DiscoveryService.Discover:output_type -> tickex.v1.DiscoverResponse
	6, // 8: tickex.v1.DiscoveryService.Heartbeat:output_type -> tickex.v1.HeartbeatResponse
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_tickex_v1_discovery_proto_init() }
func file_tickex_v1_discovery_proto_init() {
	if File_tickex_v1_discovery_proto != nil {
		return
	}
	file_tickex_v1_service_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tickex_v1_discovery_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_tickex_v1_discovery_proto_goTypes,
		DependencyIndexes: file_tickex_v1_discovery_proto_depIdxs,
		MessageInfos:      file_tickex_v1_discovery_proto_msgTypes,
	}.Build()
	File_tickex_v1_discovery_proto = out.File
	file_tickex_v1_discovery_proto_rawDesc = nil
	file_tickex_v1_discovery_proto_goTypes = nil
	file_tickex_v1_discovery_proto_depIdxs = nil
}
