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

// Package cli provide flag variable props
package cli

import (
	"flag"
	"fmt"
	"sync"

	"github.com/tickexvn/tickex/api/gen/go/common/flags/v1"
)

var once sync.Once

// Flags global variable
var Flags = &flags.Flags{
	IsTurnOnBots: false,
	Name:         "Tickex mesh server",
	Address:      "0.0.0.0:9000",
}

// Parse flag args
func Parse() *flags.Flags {
	once.Do(func() {
		flag.BoolVar(&Flags.IsTurnOnBots, "bot", Flags.GetIsTurnOnBots(), "turn on bots?")
		flag.StringVar(&Flags.Name, "name", Flags.GetName(), "hostname?")
		flag.StringVar(&Flags.Address, "address", Flags.GetAddress(), "host address?")

		flag.Usage = func() {
			fmt.Println("Usage: tickex [Flags]")
			flag.PrintDefaults()
		}
		flag.Parse()
	})

	return Flags
}
