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

// Package gateway provides the gateway of tickex.
package gateway

import (
	"context"

	typepb "github.com/tickexvn/tickex/api/gen/go/types/v1"
	"github.com/tickexvn/tickex/internal/gateway/services/greeter"
	"github.com/tickexvn/tickex/internal/gateway/types"
	"github.com/tickexvn/tickex/internal/gateway/visitor"
	"github.com/tickexvn/tickex/pkg/core"
	"github.com/tickexvn/tickex/pkg/core/syslog"
	"github.com/tickexvn/tickex/pkg/errors"
	"github.com/tickexvn/tickex/pkg/logger"
	"github.com/tickexvn/tickex/pkg/msgf"
	"github.com/tickexvn/tickex/pkg/pbtools"
	"github.com/tickexvn/tickex/pkg/utils"
)

var _ core.Server = (*Engine)(nil)

// Engine represents the gateway app
type Engine struct {
	config  *typepb.Config
	mux     core.IServeMux
	visitor types.IVisitor
}

func (e *Engine) visit(ctx context.Context, services ...types.IService) error {
	for _, service := range services {
		if err := service.Accept(ctx, e.mux, e.visitor); err != nil {
			return err
		}
	}

	return nil
}

// Declare function in gateway/types at types.IVisitor interface
//
// Ex:
//
//	type IVisitor interface {
//		VisitGreeterService(ctx context.Context, mux core.IServeMux, service IService) error
//	}
//
// Implement function at visitor.Visitor:
//
// Ex:
//
//	func (v *Visitor) VisitGreeterService(ctx context.Context, mux core.IServeMux, service types.IService) error {
//		opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
//
//		greeterAddr := ":8000"
//		if err := core.RegisterService(ctx, mux, service, greeterAddr, opts); err != nil {
//			return err
//		}
//
//		return nil
//	}
func (e *Engine) register(ctx context.Context) error {
	// TODO: Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	// Create folder at services, inherit base package, override function, implement business logic
	// See: gateway/services/greeter
	services := []types.IService{
		// Example: register the greeter service to the gateway
		&greeter.Greeter{},
		// Add more services here ...
	}

	return e.visit(ctx, services...)
}

// ListenAndServe the gateway app
func (e *Engine) ListenAndServe() error {
	if err := pbtools.Validate(e.config); err != nil {
		errs := errors.New(typepb.Errors_ERRORS_INVALID_DATA, "validation failed", err)
		syslog.Error(errs.Error())

		return errs.Combine()
	}

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	if err := e.register(ctx); err != nil {
		return err
	}

	// Listen HTTP server (and mux calls to gRPC server endpoint)
	log := logger.New()
	log.Sugar().Infof(msgf.InfoHTTPServer, e.config.GetGatewayAddress())
	defer utils.CallBack(log.Sync)

	return e.mux.Listen(e.config.GetGatewayAddress())
}

// New creates a new gateway app
func New(conf *typepb.Config) core.Server {
	return &Engine{
		mux:     core.NewServeMux(),
		visitor: visitor.New(),
		config:  conf,
	}
}
