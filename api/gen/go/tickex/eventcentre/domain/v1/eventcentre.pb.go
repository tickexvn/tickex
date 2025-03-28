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
// source: tickex/eventcentre/domain/v1/eventcentre.proto

package domain

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	v11 "github.com/tickexvn/tickex/api/gen/go/tickex/eventcentre/shared/v1"
	v1 "github.com/tickexvn/tickex/api/gen/go/tickex/v1"
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

type SayHelloRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page *v1.Pages `protobuf:"bytes,1,opt,name=page,proto3" json:"page,omitempty"`
	Name string    `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *SayHelloRequest) Reset() {
	*x = SayHelloRequest{}
	mi := &file_tickex_eventcentre_domain_v1_eventcentre_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SayHelloRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SayHelloRequest) ProtoMessage() {}

func (x *SayHelloRequest) ProtoReflect() protoreflect.Message {
	mi := &file_tickex_eventcentre_domain_v1_eventcentre_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SayHelloRequest.ProtoReflect.Descriptor instead.
func (*SayHelloRequest) Descriptor() ([]byte, []int) {
	return file_tickex_eventcentre_domain_v1_eventcentre_proto_rawDescGZIP(), []int{0}
}

func (x *SayHelloRequest) GetPage() *v1.Pages {
	if x != nil {
		return x.Page
	}
	return nil
}

func (x *SayHelloRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type SayHelloResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response *v11.SayHelloResponse `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *SayHelloResponse) Reset() {
	*x = SayHelloResponse{}
	mi := &file_tickex_eventcentre_domain_v1_eventcentre_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SayHelloResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SayHelloResponse) ProtoMessage() {}

func (x *SayHelloResponse) ProtoReflect() protoreflect.Message {
	mi := &file_tickex_eventcentre_domain_v1_eventcentre_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SayHelloResponse.ProtoReflect.Descriptor instead.
func (*SayHelloResponse) Descriptor() ([]byte, []int) {
	return file_tickex_eventcentre_domain_v1_eventcentre_proto_rawDescGZIP(), []int{1}
}

func (x *SayHelloResponse) GetResponse() *v11.SayHelloResponse {
	if x != nil {
		return x.Response
	}
	return nil
}

var File_tickex_eventcentre_domain_v1_eventcentre_proto protoreflect.FileDescriptor

var file_tickex_eventcentre_domain_v1_eventcentre_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x78, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x63, 0x65,
	0x6e, 0x74, 0x72, 0x65, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x1c, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x78, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x63, 0x65,
	0x6e, 0x74, 0x72, 0x65, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x1a, 0x1b,
	0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x74, 0x69, 0x63,
	0x6b, 0x65, 0x78, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x78, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x63,
	0x65, 0x6e, 0x74, 0x72, 0x65, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x76, 0x31, 0x2f,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x54, 0x0a, 0x0f, 0x53, 0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x10, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x50,
	0x61, 0x67, 0x65, 0x73, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x1b, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xba, 0x48, 0x04, 0x72, 0x02, 0x10,
	0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x5e, 0x0a, 0x10, 0x53, 0x61, 0x79, 0x48, 0x65,
	0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4a, 0x0a, 0x08, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e,
	0x74, 0x69, 0x63, 0x6b, 0x65, 0x78, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x63, 0x65, 0x6e, 0x74,
	0x72, 0x65, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x61, 0x79,
	0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x08, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x87, 0x01, 0x0a, 0x18, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x43, 0x65, 0x6e, 0x74, 0x72, 0x65, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x6b, 0x0a, 0x08, 0x53, 0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f,
	0x12, 0x2d, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x78, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x63,
	0x65, 0x6e, 0x74, 0x72, 0x65, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x2e, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x78, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x63, 0x65,
	0x6e, 0x74, 0x72, 0x65, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53,
	0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x4b, 0x5a, 0x49, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x74, 0x69, 0x63, 0x6b, 0x65, 0x78, 0x76, 0x6e, 0x2f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x78, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x74, 0x69, 0x63, 0x6b, 0x65,
	0x78, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x65, 0x2f, 0x64, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x3b, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tickex_eventcentre_domain_v1_eventcentre_proto_rawDescOnce sync.Once
	file_tickex_eventcentre_domain_v1_eventcentre_proto_rawDescData = file_tickex_eventcentre_domain_v1_eventcentre_proto_rawDesc
)

func file_tickex_eventcentre_domain_v1_eventcentre_proto_rawDescGZIP() []byte {
	file_tickex_eventcentre_domain_v1_eventcentre_proto_rawDescOnce.Do(func() {
		file_tickex_eventcentre_domain_v1_eventcentre_proto_rawDescData = protoimpl.X.CompressGZIP(file_tickex_eventcentre_domain_v1_eventcentre_proto_rawDescData)
	})
	return file_tickex_eventcentre_domain_v1_eventcentre_proto_rawDescData
}

var file_tickex_eventcentre_domain_v1_eventcentre_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_tickex_eventcentre_domain_v1_eventcentre_proto_goTypes = []any{
	(*SayHelloRequest)(nil),      // 0: tickex.eventcentre.domain.v1.SayHelloRequest
	(*SayHelloResponse)(nil),     // 1: tickex.eventcentre.domain.v1.SayHelloResponse
	(*v1.Pages)(nil),             // 2: tickex.v1.Pages
	(*v11.SayHelloResponse)(nil), // 3: tickex.eventcentre.shared.v1.SayHelloResponse
}
var file_tickex_eventcentre_domain_v1_eventcentre_proto_depIdxs = []int32{
	2, // 0: tickex.eventcentre.domain.v1.SayHelloRequest.page:type_name -> tickex.v1.Pages
	3, // 1: tickex.eventcentre.domain.v1.SayHelloResponse.response:type_name -> tickex.eventcentre.shared.v1.SayHelloResponse
	0, // 2: tickex.eventcentre.domain.v1.EventCentreDomainService.SayHello:input_type -> tickex.eventcentre.domain.v1.SayHelloRequest
	1, // 3: tickex.eventcentre.domain.v1.EventCentreDomainService.SayHello:output_type -> tickex.eventcentre.domain.v1.SayHelloResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_tickex_eventcentre_domain_v1_eventcentre_proto_init() }
func file_tickex_eventcentre_domain_v1_eventcentre_proto_init() {
	if File_tickex_eventcentre_domain_v1_eventcentre_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tickex_eventcentre_domain_v1_eventcentre_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_tickex_eventcentre_domain_v1_eventcentre_proto_goTypes,
		DependencyIndexes: file_tickex_eventcentre_domain_v1_eventcentre_proto_depIdxs,
		MessageInfos:      file_tickex_eventcentre_domain_v1_eventcentre_proto_msgTypes,
	}.Build()
	File_tickex_eventcentre_domain_v1_eventcentre_proto = out.File
	file_tickex_eventcentre_domain_v1_eventcentre_proto_rawDesc = nil
	file_tickex_eventcentre_domain_v1_eventcentre_proto_goTypes = nil
	file_tickex_eventcentre_domain_v1_eventcentre_proto_depIdxs = nil
}
