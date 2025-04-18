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

// Package uuid provides a simple UUID generator.
package uuid

import (
	"fmt"
	"strings"

	"github.com/celestinals/celestinal/internal/pkg/version"
	"github.com/google/uuid"
)

// Generate generates a new UUID and returns it as a string.
func Generate() string {
	id := uuid.New()
	return fmt.Sprintf("%s-%s", strings.ToLower(version.Code), id.String())
}
