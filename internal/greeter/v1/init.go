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

// Package greeter provides the initialization logic for the greeter service.
package greeter

import (
	"github.com/celestinals/celestinal/internal/greeter/v1/controllers"
	"github.com/celestinals/celestinal/internal/greeter/v1/domain"
	"github.com/celestinals/celestinal/pkg/striker"
)

var (
	// handlers/controllers layer
	_ = striker.Inject(controllers.New)

	// domain layer
	_ = striker.Inject(domain.New)

	// repo layer
	//_ = striker.Inject(repos.NewAuthor)

	// data layer
	//_ = striker.Inject(authors.New)
)
