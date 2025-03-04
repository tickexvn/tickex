// Copyright 2025 The Tickex Authors.

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.21.12
// source: types.proto

package types

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Pages struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Index int32 `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	Size  int32 `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	Total bool  `protobuf:"varint,3,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *Pages) Reset() {
	*x = Pages{}
	mi := &file_types_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Pages) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pages) ProtoMessage() {}

func (x *Pages) ProtoReflect() protoreflect.Message {
	mi := &file_types_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pages.ProtoReflect.Descriptor instead.
func (*Pages) Descriptor() ([]byte, []int) {
	return file_types_proto_rawDescGZIP(), []int{0}
}

func (x *Pages) GetIndex() int32 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *Pages) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *Pages) GetTotal() bool {
	if x != nil {
		return x.Total
	}
	return false
}

type Metadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	Author    string                 `protobuf:"bytes,2,opt,name=author,proto3" json:"author,omitempty"`
}

func (x *Metadata) Reset() {
	*x = Metadata{}
	mi := &file_types_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Metadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Metadata) ProtoMessage() {}

func (x *Metadata) ProtoReflect() protoreflect.Message {
	mi := &file_types_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Metadata.ProtoReflect.Descriptor instead.
func (*Metadata) Descriptor() ([]byte, []int) {
	return file_types_proto_rawDescGZIP(), []int{1}
}

func (x *Metadata) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Metadata) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

type RobotMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metadata *Metadata `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Header   string    `protobuf:"bytes,2,opt,name=header,proto3" json:"header,omitempty"`
	Body     string    `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
	Footer   string    `protobuf:"bytes,4,opt,name=footer,proto3" json:"footer,omitempty"`
}

func (x *RobotMessage) Reset() {
	*x = RobotMessage{}
	mi := &file_types_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RobotMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RobotMessage) ProtoMessage() {}

func (x *RobotMessage) ProtoReflect() protoreflect.Message {
	mi := &file_types_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RobotMessage.ProtoReflect.Descriptor instead.
func (*RobotMessage) Descriptor() ([]byte, []int) {
	return file_types_proto_rawDescGZIP(), []int{2}
}

func (x *RobotMessage) GetMetadata() *Metadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *RobotMessage) GetHeader() string {
	if x != nil {
		return x.Header
	}
	return ""
}

func (x *RobotMessage) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

func (x *RobotMessage) GetFooter() string {
	if x != nil {
		return x.Footer
	}
	return ""
}

type Context struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TraceId       string `protobuf:"bytes,1,opt,name=trace_id,json=traceId,proto3" json:"trace_id,omitempty"`                   // Tracks each request across microservices.
	RequestId     string `protobuf:"bytes,2,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`             // Unique identifier for each request. Useful for debugging and logging.
	UserId        string `protobuf:"bytes,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`                      // Identifies the user making the request.
	Authorization string `protobuf:"bytes,4,opt,name=authorization,proto3" json:"authorization,omitempty"`                      // Passes tokens or authentication information through the context.
	Locale        string `protobuf:"bytes,5,opt,name=locale,proto3" json:"locale,omitempty"`                                    // Passes language information for multilingual systems.
	CorrelationId string `protobuf:"bytes,6,opt,name=correlation_id,json=correlationId,proto3" json:"correlation_id,omitempty"` // Links multiple related requests in a workflow.
	ServiceName   string `protobuf:"bytes,7,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`       // Identifies the current service processing the request.
	RetryCount    uint32 `protobuf:"varint,8,opt,name=retry_count,json=retryCount,proto3" json:"retry_count,omitempty"`         // Tracks the number of retries to prevent infinite loops.
	SpanId        string `protobuf:"bytes,9,opt,name=span_id,json=spanId,proto3" json:"span_id,omitempty"`                      // Identifies individual parts within a request (tracing).
	Ip            string `protobuf:"bytes,10,opt,name=ip,proto3" json:"ip,omitempty"`                                           // Identifies the IP address of the client sending the request.
	Environment   string `protobuf:"bytes,11,opt,name=environment,proto3" json:"environment,omitempty"`                         // Identifies the environment in which the service is running.
}

func (x *Context) Reset() {
	*x = Context{}
	mi := &file_types_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Context) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Context) ProtoMessage() {}

func (x *Context) ProtoReflect() protoreflect.Message {
	mi := &file_types_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Context.ProtoReflect.Descriptor instead.
func (*Context) Descriptor() ([]byte, []int) {
	return file_types_proto_rawDescGZIP(), []int{3}
}

func (x *Context) GetTraceId() string {
	if x != nil {
		return x.TraceId
	}
	return ""
}

func (x *Context) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *Context) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Context) GetAuthorization() string {
	if x != nil {
		return x.Authorization
	}
	return ""
}

func (x *Context) GetLocale() string {
	if x != nil {
		return x.Locale
	}
	return ""
}

func (x *Context) GetCorrelationId() string {
	if x != nil {
		return x.CorrelationId
	}
	return ""
}

func (x *Context) GetServiceName() string {
	if x != nil {
		return x.ServiceName
	}
	return ""
}

func (x *Context) GetRetryCount() uint32 {
	if x != nil {
		return x.RetryCount
	}
	return 0
}

func (x *Context) GetSpanId() string {
	if x != nil {
		return x.SpanId
	}
	return ""
}

func (x *Context) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *Context) GetEnvironment() string {
	if x != nil {
		return x.Environment
	}
	return ""
}

type Config struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServiceRegistryAddress string `protobuf:"bytes,1,opt,name=service_registry_address,json=serviceRegistryAddress,proto3" json:"service_registry_address,omitempty"`
	GatewayAddress         string `protobuf:"bytes,2,opt,name=gateway_address,json=gatewayAddress,proto3" json:"gateway_address,omitempty"`
	Env                    string `protobuf:"bytes,3,opt,name=env,proto3" json:"env,omitempty"`
	BotToken               string `protobuf:"bytes,4,opt,name=bot_token,json=botToken,proto3" json:"bot_token,omitempty"`
	ChatId                 int64  `protobuf:"varint,5,opt,name=chat_id,json=chatId,proto3" json:"chat_id,omitempty"`
}

func (x *Config) Reset() {
	*x = Config{}
	mi := &file_types_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_types_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Config.ProtoReflect.Descriptor instead.
func (*Config) Descriptor() ([]byte, []int) {
	return file_types_proto_rawDescGZIP(), []int{4}
}

func (x *Config) GetServiceRegistryAddress() string {
	if x != nil {
		return x.ServiceRegistryAddress
	}
	return ""
}

func (x *Config) GetGatewayAddress() string {
	if x != nil {
		return x.GatewayAddress
	}
	return ""
}

func (x *Config) GetEnv() string {
	if x != nil {
		return x.Env
	}
	return ""
}

func (x *Config) GetBotToken() string {
	if x != nil {
		return x.BotToken
	}
	return ""
}

func (x *Config) GetChatId() int64 {
	if x != nil {
		return x.ChatId
	}
	return 0
}

type Service struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Host string   `protobuf:"bytes,3,opt,name=host,proto3" json:"host,omitempty"`
	Port uint32   `protobuf:"varint,4,opt,name=port,proto3" json:"port,omitempty"`
	Tags []string `protobuf:"bytes,5,rep,name=tags,proto3" json:"tags,omitempty"`
}

func (x *Service) Reset() {
	*x = Service{}
	mi := &file_types_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Service) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Service) ProtoMessage() {}

func (x *Service) ProtoReflect() protoreflect.Message {
	mi := &file_types_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Service.ProtoReflect.Descriptor instead.
func (*Service) Descriptor() ([]byte, []int) {
	return file_types_proto_rawDescGZIP(), []int{5}
}

func (x *Service) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Service) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Service) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *Service) GetPort() uint32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *Service) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

type Flags struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TurnOnBots bool   `protobuf:"varint,1,opt,name=turn_on_bots,json=turnOnBots,proto3" json:"turn_on_bots,omitempty"`
	Name       string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Address    string `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *Flags) Reset() {
	*x = Flags{}
	mi := &file_types_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Flags) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Flags) ProtoMessage() {}

func (x *Flags) ProtoReflect() protoreflect.Message {
	mi := &file_types_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Flags.ProtoReflect.Descriptor instead.
func (*Flags) Descriptor() ([]byte, []int) {
	return file_types_proto_rawDescGZIP(), []int{6}
}

func (x *Flags) GetTurnOnBots() bool {
	if x != nil {
		return x.TurnOnBots
	}
	return false
}

func (x *Flags) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Flags) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

var File_types_proto protoreflect.FileDescriptor

var file_types_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x74,
	0x69, 0x63, 0x6b, 0x65, 0x78, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1b,
	0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x47, 0x0a, 0x05,
	0x50, 0x61, 0x67, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x12, 0x0a, 0x04, 0x73,
	0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x5d, 0x0a, 0x08, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x22, 0x89, 0x01, 0x0a, 0x0c, 0x52, 0x6f, 0x62, 0x6f, 0x74, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x35, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x78,
	0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x16, 0x0a, 0x06,
	0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x68, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x6f, 0x6f, 0x74,
	0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x6f, 0x6f, 0x74, 0x65, 0x72,
	0x22, 0x98, 0x03, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x22, 0x0a, 0x08,
	0x74, 0x72, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07,
	0xba, 0x48, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x07, 0x74, 0x72, 0x61, 0x63, 0x65, 0x49, 0x64,
	0x12, 0x26, 0x0a, 0x0a, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xba, 0x48, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x09, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xba, 0x48, 0x04, 0x72, 0x02,
	0x10, 0x01, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0d, 0x61, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x1f, 0x0a, 0x06, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x07, 0xba, 0x48, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x06, 0x6c, 0x6f, 0x63, 0x61, 0x6c,
	0x65, 0x12, 0x2e, 0x0a, 0x0e, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xba, 0x48, 0x04, 0x72, 0x02,
	0x10, 0x01, 0x52, 0x0d, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x64, 0x12, 0x2a, 0x0a, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xba, 0x48, 0x04, 0x72, 0x02, 0x10, 0x01,
	0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a,
	0x0b, 0x72, 0x65, 0x74, 0x72, 0x79, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0a, 0x72, 0x65, 0x74, 0x72, 0x79, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x17,
	0x0a, 0x07, 0x73, 0x70, 0x61, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x70, 0x61, 0x6e, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x07, 0xba, 0x48, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x02, 0x69, 0x70,
	0x12, 0x29, 0x0a, 0x0b, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xba, 0x48, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x0b,
	0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x98, 0x02, 0x0a, 0x06,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x5c, 0x0a, 0x18, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x22, 0xba, 0x48, 0x1f, 0x72, 0x1d, 0x32,
	0x1b, 0x5e, 0x5b, 0x61, 0x2d, 0x7a, 0x41, 0x2d, 0x5a, 0x30, 0x2d, 0x39, 0x2e, 0x2d, 0x5d, 0x2b,
	0x3a, 0x5b, 0x30, 0x2d, 0x39, 0x5d, 0x7b, 0x31, 0x2c, 0x35, 0x7d, 0x24, 0x52, 0x16, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x12, 0x4b, 0x0a, 0x0f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x5f,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x22, 0xba,
	0x48, 0x1f, 0x72, 0x1d, 0x32, 0x1b, 0x5e, 0x5b, 0x61, 0x2d, 0x7a, 0x41, 0x2d, 0x5a, 0x30, 0x2d,
	0x39, 0x2e, 0x2d, 0x5d, 0x2b, 0x3a, 0x5b, 0x30, 0x2d, 0x39, 0x5d, 0x7b, 0x31, 0x2c, 0x35, 0x7d,
	0x24, 0x52, 0x0e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x12, 0x2d, 0x0a, 0x03, 0x65, 0x6e, 0x76, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1b,
	0xba, 0x48, 0x18, 0x72, 0x16, 0x32, 0x14, 0x5e, 0x28, 0x64, 0x65, 0x76, 0x7c, 0x70, 0x72, 0x6f,
	0x64, 0x7c, 0x73, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x78, 0x29, 0x24, 0x52, 0x03, 0x65, 0x6e, 0x76,
	0x12, 0x1b, 0x0a, 0x09, 0x62, 0x6f, 0x74, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x6f, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x17, 0x0a,
	0x07, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x63, 0x68, 0x61, 0x74, 0x49, 0x64, 0x22, 0x8e, 0x01, 0x0a, 0x07, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x17, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07,
	0xba, 0x48, 0x04, 0x72, 0x02, 0x10, 0x02, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xba, 0x48, 0x04, 0x72, 0x02,
	0x10, 0x02, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xba, 0x48, 0x04, 0x72, 0x02, 0x10, 0x02, 0x52,
	0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0d, 0x42, 0x08, 0xba, 0x48, 0x05, 0x2a, 0x03, 0x28, 0xe8, 0x07, 0x52, 0x04, 0x70,
	0x6f, 0x72, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x22, 0x7b, 0x0a, 0x05, 0x46, 0x6c, 0x61, 0x67, 0x73,
	0x12, 0x20, 0x0a, 0x0c, 0x74, 0x75, 0x72, 0x6e, 0x5f, 0x6f, 0x6e, 0x5f, 0x62, 0x6f, 0x74, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x74, 0x75, 0x72, 0x6e, 0x4f, 0x6e, 0x42, 0x6f,
	0x74, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x3c, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x22, 0xba, 0x48, 0x1f, 0x72, 0x1d, 0x32, 0x1b,
	0x5e, 0x5b, 0x61, 0x2d, 0x7a, 0x41, 0x2d, 0x5a, 0x30, 0x2d, 0x39, 0x2e, 0x2d, 0x5d, 0x2b, 0x3a,
	0x5b, 0x30, 0x2d, 0x39, 0x5d, 0x7b, 0x31, 0x2c, 0x35, 0x7d, 0x24, 0x52, 0x07, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x42, 0x36, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x78, 0x76, 0x6e, 0x2f, 0x74, 0x69, 0x63, 0x6b,
	0x65, 0x78, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x74, 0x79, 0x70, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_types_proto_rawDescOnce sync.Once
	file_types_proto_rawDescData = file_types_proto_rawDesc
)

func file_types_proto_rawDescGZIP() []byte {
	file_types_proto_rawDescOnce.Do(func() {
		file_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_types_proto_rawDescData)
	})
	return file_types_proto_rawDescData
}

var file_types_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_types_proto_goTypes = []any{
	(*Pages)(nil),                 // 0: tickex.types.v1.Pages
	(*Metadata)(nil),              // 1: tickex.types.v1.Metadata
	(*RobotMessage)(nil),          // 2: tickex.types.v1.RobotMessage
	(*Context)(nil),               // 3: tickex.types.v1.Context
	(*Config)(nil),                // 4: tickex.types.v1.Config
	(*Service)(nil),               // 5: tickex.types.v1.Service
	(*Flags)(nil),                 // 6: tickex.types.v1.Flags
	(*timestamppb.Timestamp)(nil), // 7: google.protobuf.Timestamp
}
var file_types_proto_depIdxs = []int32{
	7, // 0: tickex.types.v1.Metadata.created_at:type_name -> google.protobuf.Timestamp
	1, // 1: tickex.types.v1.RobotMessage.metadata:type_name -> tickex.types.v1.Metadata
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_types_proto_init() }
func file_types_proto_init() {
	if File_types_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_types_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_types_proto_goTypes,
		DependencyIndexes: file_types_proto_depIdxs,
		MessageInfos:      file_types_proto_msgTypes,
	}.Build()
	File_types_proto = out.File
	file_types_proto_rawDesc = nil
	file_types_proto_goTypes = nil
	file_types_proto_depIdxs = nil
}
