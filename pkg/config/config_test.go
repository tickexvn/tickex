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

package config

import (
	"testing"

	"github.com/tickexvn/tickex/api/gen/go/tickex/v1"
	"github.com/tickexvn/tickex/pkg/protobuf"
)

func TestConfig(t *testing.T) {
	conf := tickex.Config{
		ServiceRegistryAddr: "0.0.0.0:8500",
		ApiAddr:             "0.0.0.0:9000",
		Env:                 "prod",
	}

	if err := protobuf.Validate(&conf); err != nil {
		t.Error(err)
	}
}

func TestConfigEnv(t *testing.T) {
	conf := Default()

	if err := protobuf.Validate(conf); err != nil {
		return
	}

	t.Error("should not validate env")
}

func BenchmarkConfigHeapAllocation(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = Default()
	}
}
