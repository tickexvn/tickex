/*
 * Copyright 2025 The Tickex Authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Package services provides all service declare.
package services

import (
	"context"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	greeterpb "github.com/tickexvn/tickex/api/gen/go/greeter/v1"
	"github.com/tickexvn/tickex/internal/edge/services"
	"github.com/tickexvn/tickex/internal/edge/types"
	"github.com/tickexvn/tickex/pkg/core"
	"github.com/tickexvn/tickex/pkg/namespace"
	"google.golang.org/grpc"
)

var _ types.IService = (*greeter)(nil)

// NewGreeter creates a new greeter service to register handler to gateway
func NewGreeter() types.IService {
	return greeter{}
}

// greeter represents the greeter service
type greeter struct {
	services.Base
}

// Accept accepts the greeter service
func (g greeter) Accept(
	ctx context.Context, edge core.Edge, v types.IVisitor) error {

	return v.VisitService(ctx, namespace.GreeterV1, edge, g)
}

// Register registers the greeter service
func (g greeter) Register(ctx context.Context, mux *runtime.ServeMux,
	endpoint string, opts []grpc.DialOption) error {

	return greeterpb.
		RegisterGreeterServiceHandlerFromEndpoint(ctx, mux, endpoint, opts)
}
