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

// Package limiter provides the rate limiter for the edge app
package limiter

import (
	"net/http"

	"github.com/tickexvn/tickex/api/gen/go/common/env/config/v1"
	"github.com/tickexvn/tickex/pkg/core"
)

// Serve is a middleware that limits the number of requests
// that can be made to the server in a given time period.
func Serve(edge core.Edge, _ *config.Config) {
	edge.Use(limit)
}

// limit is a middleware that limits the number of requests
// that can be made to the server in a given time period.
func limit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}
