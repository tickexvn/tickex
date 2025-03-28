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
// source: tickex/v1/options.proto

package tickex

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TickexMethodOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ignore  bool       `protobuf:"varint,1,opt,name=ignore,proto3" json:"ignore,omitempty"`
	Require []*Require `protobuf:"bytes,2,rep,name=require,proto3" json:"require,omitempty"`
}

func (x *TickexMethodOptions) Reset() {
	*x = TickexMethodOptions{}
	mi := &file_tickex_v1_options_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TickexMethodOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TickexMethodOptions) ProtoMessage() {}

func (x *TickexMethodOptions) ProtoReflect() protoreflect.Message {
	mi := &file_tickex_v1_options_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TickexMethodOptions.ProtoReflect.Descriptor instead.
func (*TickexMethodOptions) Descriptor() ([]byte, []int) {
	return file_tickex_v1_options_proto_rawDescGZIP(), []int{0}
}

func (x *TickexMethodOptions) GetIgnore() bool {
	if x != nil {
		return x.Ignore
	}
	return false
}

func (x *TickexMethodOptions) GetRequire() []*Require {
	if x != nil {
		return x.Require
	}
	return nil
}

type Require struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Role       Role       `protobuf:"varint,1,opt,name=role,proto3,enum=tickex.v1.Role" json:"role,omitempty"`
	Permission Permission `protobuf:"varint,2,opt,name=permission,proto3,enum=tickex.v1.Permission" json:"permission,omitempty"`
}

func (x *Require) Reset() {
	*x = Require{}
	mi := &file_tickex_v1_options_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Require) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Require) ProtoMessage() {}

func (x *Require) ProtoReflect() protoreflect.Message {
	mi := &file_tickex_v1_options_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Require.ProtoReflect.Descriptor instead.
func (*Require) Descriptor() ([]byte, []int) {
	return file_tickex_v1_options_proto_rawDescGZIP(), []int{1}
}

func (x *Require) GetRole() Role {
	if x != nil {
		return x.Role
	}
	return Role_ROLE_UNSPECIFIED
}

func (x *Require) GetPermission() Permission {
	if x != nil {
		return x.Permission
	}
	return Permission_PERMISSION_UNSPECIFIED
}

var file_tickex_v1_options_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*TickexMethodOptions)(nil),
		Field:         50001,
		Name:          "tickex.v1.options",
		Tag:           "bytes,50001,opt,name=options",
		Filename:      "tickex/v1/options.proto",
	},
}

// Extension fields to descriptorpb.MethodOptions.
var (
	// optional tickex.v1.TickexMethodOptions options = 50001;
	E_Options = &file_tickex_v1_options_proto_extTypes[0]
)

var File_tickex_v1_options_proto protoreflect.FileDescriptor

var file_tickex_v1_options_proto_rawDesc = []byte{
	0x0a, 0x17, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x78, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x74, 0x69, 0x63, 0x6b, 0x65,
	0x78, 0x2e, 0x76, 0x31, 0x1a, 0x14, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x78, 0x2f, 0x76, 0x31, 0x2f,
	0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x74, 0x69, 0x63, 0x6b,
	0x65, 0x78, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5b, 0x0a, 0x13, 0x54, 0x69, 0x63, 0x6b,
	0x65, 0x78, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12,
	0x16, 0x0a, 0x06, 0x69, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x06, 0x69, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x12, 0x2c, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x69,
	0x72, 0x65, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65,
	0x78, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x52, 0x07, 0x72, 0x65,
	0x71, 0x75, 0x69, 0x72, 0x65, 0x22, 0x65, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65,
	0x12, 0x23, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f,
	0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x52,
	0x04, 0x72, 0x6f, 0x6c, 0x65, 0x12, 0x35, 0x0a, 0x0a, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x74, 0x69, 0x63, 0x6b,
	0x65, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x52, 0x0a, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x3a, 0x5a, 0x0a, 0x07,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xd1, 0x86, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1e, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x69, 0x63, 0x6b,
	0x65, 0x78, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52,
	0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x38, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x78, 0x76, 0x6e, 0x2f,
	0x74, 0x69, 0x63, 0x6b, 0x65, 0x78, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67,
	0x6f, 0x2f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x78, 0x2f, 0x76, 0x31, 0x3b, 0x74, 0x69, 0x63, 0x6b,
	0x65, 0x78, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tickex_v1_options_proto_rawDescOnce sync.Once
	file_tickex_v1_options_proto_rawDescData = file_tickex_v1_options_proto_rawDesc
)

func file_tickex_v1_options_proto_rawDescGZIP() []byte {
	file_tickex_v1_options_proto_rawDescOnce.Do(func() {
		file_tickex_v1_options_proto_rawDescData = protoimpl.X.CompressGZIP(file_tickex_v1_options_proto_rawDescData)
	})
	return file_tickex_v1_options_proto_rawDescData
}

var file_tickex_v1_options_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_tickex_v1_options_proto_goTypes = []any{
	(*TickexMethodOptions)(nil),        // 0: tickex.v1.TickexMethodOptions
	(*Require)(nil),                    // 1: tickex.v1.Require
	(Role)(0),                          // 2: tickex.v1.Role
	(Permission)(0),                    // 3: tickex.v1.Permission
	(*descriptorpb.MethodOptions)(nil), // 4: google.protobuf.MethodOptions
}
var file_tickex_v1_options_proto_depIdxs = []int32{
	1, // 0: tickex.v1.TickexMethodOptions.require:type_name -> tickex.v1.Require
	2, // 1: tickex.v1.Require.role:type_name -> tickex.v1.Role
	3, // 2: tickex.v1.Require.permission:type_name -> tickex.v1.Permission
	4, // 3: tickex.v1.options:extendee -> google.protobuf.MethodOptions
	0, // 4: tickex.v1.options:type_name -> tickex.v1.TickexMethodOptions
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	4, // [4:5] is the sub-list for extension type_name
	3, // [3:4] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_tickex_v1_options_proto_init() }
func file_tickex_v1_options_proto_init() {
	if File_tickex_v1_options_proto != nil {
		return
	}
	file_tickex_v1_role_proto_init()
	file_tickex_v1_permission_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tickex_v1_options_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 1,
			NumServices:   0,
		},
		GoTypes:           file_tickex_v1_options_proto_goTypes,
		DependencyIndexes: file_tickex_v1_options_proto_depIdxs,
		MessageInfos:      file_tickex_v1_options_proto_msgTypes,
		ExtensionInfos:    file_tickex_v1_options_proto_extTypes,
	}.Build()
	File_tickex_v1_options_proto = out.File
	file_tickex_v1_options_proto_rawDesc = nil
	file_tickex_v1_options_proto_goTypes = nil
	file_tickex_v1_options_proto_depIdxs = nil
}
