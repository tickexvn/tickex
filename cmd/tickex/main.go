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

// Package main provides the entry point for the Tickex.
package main

import (
	"context"

	"github.com/tickexvn/tickex/internal/edge"
	"github.com/tickexvn/tickex/pkg/cli"
	"github.com/tickexvn/tickex/pkg/configs"
	"github.com/tickexvn/tickex/pkg/core"
	"github.com/tickexvn/tickex/pkg/logger"
	"github.com/tickexvn/tickex/pkg/namespace"
)

// Build and run the main application with environment variables.
// Remember to inject all layers of the application using the
// core.Inject() function.
//
// Example:
//
//	_ = core.Inject(controllers.New)
//
// This is the Tickex edge application, it will automatically connect to
// other services via gRPC. Start the application along with other services
// in the x/ directory.The application provides APIs for users through a
// single HTTP gateway following the RESTful API standard. The application
// uses gRPC to connect to other services.Additionally, the system provides
// a Swagger UI interface for users to easily interact with the system
// through a web interface.
//
// Start the application using the Makefile command
//
//	make run.tickex // start tickex edge
//	make run.x.<service> // start service
func main() {
	cli.ServiceFlags.Name = namespace.Edge
	_ = cli.ParseEdge()

	app := core.Build(edge.New, configs.Default)
	logger.Fatal(app.Start(context.Background()))
}
