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
// source: ticket_domain.proto

package domain

import (
	v1 "github.com/tickexvn/tickex/api/gen/go/ticket/shared/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetTicketRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetTicketRequest) Reset() {
	*x = GetTicketRequest{}
	mi := &file_ticket_domain_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetTicketRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTicketRequest) ProtoMessage() {}

func (x *GetTicketRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ticket_domain_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTicketRequest.ProtoReflect.Descriptor instead.
func (*GetTicketRequest) Descriptor() ([]byte, []int) {
	return file_ticket_domain_proto_rawDescGZIP(), []int{0}
}

type CreateTicketRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateTicketRequest) Reset() {
	*x = CreateTicketRequest{}
	mi := &file_ticket_domain_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateTicketRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTicketRequest) ProtoMessage() {}

func (x *CreateTicketRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ticket_domain_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTicketRequest.ProtoReflect.Descriptor instead.
func (*CreateTicketRequest) Descriptor() ([]byte, []int) {
	return file_ticket_domain_proto_rawDescGZIP(), []int{1}
}

type DeleteTicketRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteTicketRequest) Reset() {
	*x = DeleteTicketRequest{}
	mi := &file_ticket_domain_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteTicketRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTicketRequest) ProtoMessage() {}

func (x *DeleteTicketRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ticket_domain_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTicketRequest.ProtoReflect.Descriptor instead.
func (*DeleteTicketRequest) Descriptor() ([]byte, []int) {
	return file_ticket_domain_proto_rawDescGZIP(), []int{2}
}

var File_ticket_domain_proto protoreflect.FileDescriptor

var file_ticket_domain_proto_rawDesc = []byte{
	0x0a, 0x13, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x78, 0x2e, 0x74, 0x69,
	0x63, 0x6b, 0x65, 0x74, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x1a, 0x2b,
	0x74, 0x69, 0x63, 0x6b, 0x65, 0x78, 0x2f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2f, 0x73, 0x68,
	0x61, 0x72, 0x65, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x73,
	0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70,
	0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x12, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x54,
	0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x15, 0x0a, 0x13,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0x15, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x69, 0x63,
	0x6b, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x32, 0xc2, 0x02, 0x0a, 0x13, 0x54,
	0x69, 0x63, 0x6b, 0x65, 0x74, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x6d, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x63, 0x6b,
	0x65, 0x74, 0x12, 0x2c, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x78, 0x2e, 0x74, 0x69, 0x63, 0x6b,
	0x65, 0x74, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x2d, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x78, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74,
	0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x64, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x29,
	0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x78, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x64,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x69, 0x63, 0x6b,
	0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x74, 0x69, 0x63, 0x6b,
	0x65, 0x78, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x56, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x2c, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x78,
	0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x76,
	0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42,
	0x3f, 0x5a, 0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x69,
	0x63, 0x6b, 0x65, 0x78, 0x76, 0x6e, 0x2f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x78, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2f,
	0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x3b, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ticket_domain_proto_rawDescOnce sync.Once
	file_ticket_domain_proto_rawDescData = file_ticket_domain_proto_rawDesc
)

func file_ticket_domain_proto_rawDescGZIP() []byte {
	file_ticket_domain_proto_rawDescOnce.Do(func() {
		file_ticket_domain_proto_rawDescData = protoimpl.X.CompressGZIP(file_ticket_domain_proto_rawDescData)
	})
	return file_ticket_domain_proto_rawDescData
}

var file_ticket_domain_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_ticket_domain_proto_goTypes = []any{
	(*GetTicketRequest)(nil),        // 0: tickex.ticket.domain.v1.GetTicketRequest
	(*CreateTicketRequest)(nil),     // 1: tickex.ticket.domain.v1.CreateTicketRequest
	(*DeleteTicketRequest)(nil),     // 2: tickex.ticket.domain.v1.DeleteTicketRequest
	(*v1.CreateTicketResponse)(nil), // 3: tickex.ticket.shared.v1.CreateTicketResponse
	(*v1.GetTicketResponse)(nil),    // 4: tickex.ticket.shared.v1.GetTicketResponse
	(*emptypb.Empty)(nil),           // 5: google.protobuf.Empty
}
var file_ticket_domain_proto_depIdxs = []int32{
	1, // 0: tickex.ticket.domain.v1.TicketDomainService.CreateTicket:input_type -> tickex.ticket.domain.v1.CreateTicketRequest
	0, // 1: tickex.ticket.domain.v1.TicketDomainService.GetTicket:input_type -> tickex.ticket.domain.v1.GetTicketRequest
	2, // 2: tickex.ticket.domain.v1.TicketDomainService.DeleteTicket:input_type -> tickex.ticket.domain.v1.DeleteTicketRequest
	3, // 3: tickex.ticket.domain.v1.TicketDomainService.CreateTicket:output_type -> tickex.ticket.shared.v1.CreateTicketResponse
	4, // 4: tickex.ticket.domain.v1.TicketDomainService.GetTicket:output_type -> tickex.ticket.shared.v1.GetTicketResponse
	5, // 5: tickex.ticket.domain.v1.TicketDomainService.DeleteTicket:output_type -> google.protobuf.Empty
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ticket_domain_proto_init() }
func file_ticket_domain_proto_init() {
	if File_ticket_domain_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_ticket_domain_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ticket_domain_proto_goTypes,
		DependencyIndexes: file_ticket_domain_proto_depIdxs,
		MessageInfos:      file_ticket_domain_proto_msgTypes,
	}.Build()
	File_ticket_domain_proto = out.File
	file_ticket_domain_proto_rawDesc = nil
	file_ticket_domain_proto_goTypes = nil
	file_ticket_domain_proto_depIdxs = nil
}
