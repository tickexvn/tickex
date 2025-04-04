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

// Package main provides the entry point for the greeter service.
package main

import (
	"context"

	"github.com/celestinals/celestinal/internal/greeter/v1/boots/server"

	"github.com/celestinals/celestinal/api/gen/go/celestinal/greeter/v1"
	_ "github.com/celestinals/celestinal/internal/greeter/v1/boots/init"

	cestconf "github.com/celestinals/celestinal/pkg/config"
	cestcore "github.com/celestinals/celestinal/pkg/core"
	cestflag "github.com/celestinals/celestinal/pkg/flag"
	cestlog "github.com/celestinals/celestinal/pkg/logger"
	cestns "github.com/celestinals/celestinal/pkg/names"
)

// Build and run main application with environment variable
// Remember to inject all layers of the application by
// cestcore.Inject() function
//
// Example:
//
// _ = cestcore.Inject(controllers.New)
func main() {
	cestflag.SetDefault(cestns.GreeterV1, "127.0.0.1:0", "dev")
	cestflag.SetConsole(greeter.ASCII)

	_ = cestflag.Parse()

	app := cestcore.Build(server.New, cestconf.Default)
	if err := app.Run(context.Background()); err != nil {
		cestlog.Fatal(err)
	}
}
