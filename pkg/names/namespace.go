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

// Package names provide name of all service at celestinal
package names

// Namespace is a type that represents the names of a service.
type Namespace string

func (ns Namespace) String() string {
	return string(ns)
}

const (
	// APIServer names info
	APIServer Namespace = "celestinal.apiserver"

	// GreeterV1 names info
	GreeterV1 Namespace = "celestinal.greeter.v1"

	// DiscoveryV1 names info
	DiscoveryV1 Namespace = "celestinal.discovery.v1"
)
