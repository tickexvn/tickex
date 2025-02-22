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

package core

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/fx"
)

type container struct {
	engine *fx.App
}

// Start implements IContainer.
func (c *container) Start(ctx context.Context) error {
	err := make(chan error)
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	go c.start(ctx, err)
	go c.stop(ctx, sig, err)

	return <-err
}

func (c *container) start(ctx context.Context, err chan<- error) {
	err <- c.engine.Start(ctx)
}

func (c *container) stop(ctx context.Context, sig <-chan os.Signal, err chan<- error) {
	<-sig
	err <- c.engine.Stop(ctx)
}
