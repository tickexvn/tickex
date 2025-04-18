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

// Package skutils provides utility functions for striker.
package skutils

import (
	"context"

	"github.com/celestinals/celestinal/pkg/striker/skhttp"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

// ServiceRegistrar is an interface for registering a gRPC service. Not a server !!!
type ServiceRegistrar interface {
	Accept(context.Context, skhttp.Server) error
	RegisterFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error
	Register(ctx context.Context, mux *runtime.ServeMux) error
}

// RegisterServiceFromEndpoint registers a service with the runtime from endpoint
// The service is registered with the runtime mux.
//
// dependency:
//   - github.com/grpc-ecosystem/grpc-gateway/v2/runtime
//   - google.golang.org/grpc
func RegisterServiceFromEndpoint(
	ctx context.Context, httpServer skhttp.Server, service ServiceRegistrar, endpoint string,
	opts []grpc.DialOption) error {

	if err := service.RegisterFromEndpoint(
		ctx, httpServer.RuntimeMux(), endpoint, opts); err != nil {
		return err
	}

	return nil
}

// RegisterService registers a service with the run time
func RegisterService(ctx context.Context, httpServer skhttp.Server, service ServiceRegistrar) error {
	return service.Register(ctx, httpServer.RuntimeMux())
}
