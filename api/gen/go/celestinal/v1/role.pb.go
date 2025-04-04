// Copyright 2025 The Celestinal Authors.
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
// source: celestinal/v1/role.proto

package celestinal

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Role int32

const (
	Role_ROLE_UNSPECIFIED Role = 0
	Role_ROLE_CUSTOMER    Role = 1
	Role_ROLE_SELLER      Role = 2
	Role_ROLE_ADMIN       Role = 3
)

// Enum value maps for Role.
var (
	Role_name = map[int32]string{
		0: "ROLE_UNSPECIFIED",
		1: "ROLE_CUSTOMER",
		2: "ROLE_SELLER",
		3: "ROLE_ADMIN",
	}
	Role_value = map[string]int32{
		"ROLE_UNSPECIFIED": 0,
		"ROLE_CUSTOMER":    1,
		"ROLE_SELLER":      2,
		"ROLE_ADMIN":       3,
	}
)

func (x Role) Enum() *Role {
	p := new(Role)
	*p = x
	return p
}

func (x Role) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Role) Descriptor() protoreflect.EnumDescriptor {
	return file_celestinal_v1_role_proto_enumTypes[0].Descriptor()
}

func (Role) Type() protoreflect.EnumType {
	return &file_celestinal_v1_role_proto_enumTypes[0]
}

func (x Role) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Role.Descriptor instead.
func (Role) EnumDescriptor() ([]byte, []int) {
	return file_celestinal_v1_role_proto_rawDescGZIP(), []int{0}
}

var File_celestinal_v1_role_proto protoreflect.FileDescriptor

var file_celestinal_v1_role_proto_rawDesc = []byte{
	0x0a, 0x18, 0x63, 0x65, 0x6c, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x6c, 0x2f, 0x76, 0x31, 0x2f,
	0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x63, 0x65, 0x6c, 0x65,
	0x73, 0x74, 0x69, 0x6e, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0x50, 0x0a, 0x04, 0x52,
	0x6f, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x10, 0x52, 0x4f, 0x4c, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50,
	0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x52, 0x4f, 0x4c,
	0x45, 0x5f, 0x43, 0x55, 0x53, 0x54, 0x4f, 0x4d, 0x45, 0x52, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b,
	0x52, 0x4f, 0x4c, 0x45, 0x5f, 0x53, 0x45, 0x4c, 0x4c, 0x45, 0x52, 0x10, 0x02, 0x12, 0x0e, 0x0a,
	0x0a, 0x52, 0x4f, 0x4c, 0x45, 0x5f, 0x41, 0x44, 0x4d, 0x49, 0x4e, 0x10, 0x03, 0x42, 0x47, 0x5a,
	0x45, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x65, 0x6c, 0x65,
	0x73, 0x74, 0x69, 0x6e, 0x61, 0x6c, 0x73, 0x2f, 0x63, 0x65, 0x6c, 0x65, 0x73, 0x74, 0x69, 0x6e,
	0x61, 0x6c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x65,
	0x6c, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x6c, 0x2f, 0x76, 0x31, 0x3b, 0x63, 0x65, 0x6c, 0x65,
	0x73, 0x74, 0x69, 0x6e, 0x61, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_celestinal_v1_role_proto_rawDescOnce sync.Once
	file_celestinal_v1_role_proto_rawDescData = file_celestinal_v1_role_proto_rawDesc
)

func file_celestinal_v1_role_proto_rawDescGZIP() []byte {
	file_celestinal_v1_role_proto_rawDescOnce.Do(func() {
		file_celestinal_v1_role_proto_rawDescData = protoimpl.X.CompressGZIP(file_celestinal_v1_role_proto_rawDescData)
	})
	return file_celestinal_v1_role_proto_rawDescData
}

var file_celestinal_v1_role_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_celestinal_v1_role_proto_goTypes = []any{
	(Role)(0), // 0: celestinal.v1.Role
}
var file_celestinal_v1_role_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_celestinal_v1_role_proto_init() }
func file_celestinal_v1_role_proto_init() {
	if File_celestinal_v1_role_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_celestinal_v1_role_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_celestinal_v1_role_proto_goTypes,
		DependencyIndexes: file_celestinal_v1_role_proto_depIdxs,
		EnumInfos:         file_celestinal_v1_role_proto_enumTypes,
	}.Build()
	File_celestinal_v1_role_proto = out.File
	file_celestinal_v1_role_proto_rawDesc = nil
	file_celestinal_v1_role_proto_goTypes = nil
	file_celestinal_v1_role_proto_depIdxs = nil
}
