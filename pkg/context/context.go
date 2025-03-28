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

// Package context provides a context with a message.
package context

import (
	"context"

	"github.com/tickexvn/tickex/api/gen/go/tickex/v1"
	"github.com/tickexvn/tickex/pkg/protobuf/proto"
)

// Key is the key type for the context.
type Key int

const contextKey Key = 0

// New returns a new context with the given message
func New(msg *tickex.Context) context.Context {
	if msg == nil {
		msg = &tickex.Context{}
	}

	ctx := context.Background()
	msgBin, _ := proto.Marshal(msg)

	return context.WithValue(ctx, contextKey, msgBin)
}

// Value returns the context value.
func Value(ctx context.Context) (*tickex.Context, bool) {
	msgBin, ok := ctx.Value(contextKey).([]byte)
	if !ok {
		return nil, false
	}

	var msg tickex.Context
	if err := proto.Unmarshal(msgBin, &msg); err != nil {
		return nil, false
	}

	return &msg, true
}
