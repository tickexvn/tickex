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

// Package visitor provides an implementation of the visitor pattern.
package visitor

import (
	"context"

	"github.com/celestinals/celestinal/internal/pkg/eventq"
	"github.com/celestinals/celestinal/pkg/striker/skhttp"
	"github.com/celestinals/celestinal/pkg/striker/skutils"

	"github.com/celestinals/celestinal/pkg/logger"
	"github.com/celestinals/celestinal/pkg/names"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// VisitServiceFromEndpoint visits and registers a service from an endpoint
func VisitServiceFromEndpoint(ctx context.Context, ns names.Namespace, server skhttp.Server,
	service skutils.ServiceRegistrar) error {

	eventq.Subscribe(ctx, ns.String(), func(endpoint string) error {
		opts := []grpc.DialOption{
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		}

		logger.Infof("[visitor.VisitServiceFromEndpoint] %s %s", ns.String(), "******")
		return skutils.RegisterServiceFromEndpoint(ctx, server, service, endpoint, opts)
	})

	return nil
}

// VisitService visits and registers a service
func VisitService(ctx context.Context, server skhttp.Server, service skutils.ServiceRegistrar) error {
	return skutils.RegisterService(ctx, server, service)
}
