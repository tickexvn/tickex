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

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: tickex/transfermarkt/v1/transfermarkt.proto

package transfermarkt

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	TransfermarktService_SayHello_FullMethodName = "/tickex.transfermarkt.v1.TransfermarktService/SayHello"
	TransfermarktService_Status_FullMethodName   = "/tickex.transfermarkt.v1.TransfermarktService/Status"
)

// TransfermarktServiceClient is the client API for TransfermarktService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TransfermarktServiceClient interface {
	SayHello(ctx context.Context, in *SayHelloRequest, opts ...grpc.CallOption) (*SayHelloResponse, error)
	Status(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*StatusResponse, error)
}

type transfermarktServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTransfermarktServiceClient(cc grpc.ClientConnInterface) TransfermarktServiceClient {
	return &transfermarktServiceClient{cc}
}

func (c *transfermarktServiceClient) SayHello(ctx context.Context, in *SayHelloRequest, opts ...grpc.CallOption) (*SayHelloResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SayHelloResponse)
	err := c.cc.Invoke(ctx, TransfermarktService_SayHello_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transfermarktServiceClient) Status(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*StatusResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(StatusResponse)
	err := c.cc.Invoke(ctx, TransfermarktService_Status_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TransfermarktServiceServer is the server API for TransfermarktService service.
// All implementations must embed UnimplementedTransfermarktServiceServer
// for forward compatibility.
type TransfermarktServiceServer interface {
	SayHello(context.Context, *SayHelloRequest) (*SayHelloResponse, error)
	Status(context.Context, *StatusRequest) (*StatusResponse, error)
	mustEmbedUnimplementedTransfermarktServiceServer()
}

// UnimplementedTransfermarktServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedTransfermarktServiceServer struct{}

func (UnimplementedTransfermarktServiceServer) SayHello(context.Context, *SayHelloRequest) (*SayHelloResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedTransfermarktServiceServer) Status(context.Context, *StatusRequest) (*StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Status not implemented")
}
func (UnimplementedTransfermarktServiceServer) mustEmbedUnimplementedTransfermarktServiceServer() {}
func (UnimplementedTransfermarktServiceServer) testEmbeddedByValue()                              {}

// UnsafeTransfermarktServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TransfermarktServiceServer will
// result in compilation errors.
type UnsafeTransfermarktServiceServer interface {
	mustEmbedUnimplementedTransfermarktServiceServer()
}

func RegisterTransfermarktServiceServer(s grpc.ServiceRegistrar, srv TransfermarktServiceServer) {
	// If the following call pancis, it indicates UnimplementedTransfermarktServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&TransfermarktService_ServiceDesc, srv)
}

func _TransfermarktService_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SayHelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransfermarktServiceServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TransfermarktService_SayHello_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransfermarktServiceServer).SayHello(ctx, req.(*SayHelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransfermarktService_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransfermarktServiceServer).Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TransfermarktService_Status_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransfermarktServiceServer).Status(ctx, req.(*StatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TransfermarktService_ServiceDesc is the grpc.ServiceDesc for TransfermarktService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TransfermarktService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tickex.transfermarkt.v1.TransfermarktService",
	HandlerType: (*TransfermarktServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _TransfermarktService_SayHello_Handler,
		},
		{
			MethodName: "Status",
			Handler:    _TransfermarktService_Status_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tickex/transfermarkt/v1/transfermarkt.proto",
}
