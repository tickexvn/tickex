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

// Package flag provide flag variable props
package flag

import (
	"fmt"
	"os"
	"sync"

	"github.com/celestinals/celestinal/api/gen/go/celestinal/v1"
	"github.com/celestinals/celestinal/internal/pkg/version"
	"github.com/celestinals/celestinal/pkg/names"

	"github.com/spf13/pflag"
)

var (
	once          sync.Once
	onceAPIServer sync.Once
)

var (
	// ascii art use in console with --help option
	asciiConsole = version.ASCIIArt
)

// flags global variable
var flags = &celestinal.Flag{
	Name:     "celestinal.server.default",
	Address:  "0.0.0.0:9000",
	Mode:     "dev",
	LogLevel: "debug",
}

// apiServerFlags global variable
var apiServerFlags = &celestinal.FlagAPIServer{
	Telegram:     false,
	ApiSpecsPath: "api/specs/v1",
	SwaggerPath:  "api/ui/swagger",
}

// Parse flag args
func Parse() *celestinal.Flag {
	once.Do(func() {
		pflag.StringVarP(&flags.Address, "address", "a", flags.GetAddress(), "host address ?")
		pflag.StringVarP(&flags.Mode, "mode", "m", flags.GetMode(), "run mode (dev|prod|sandbox) ?")
		pflag.StringVar(&flags.LogLevel, "log-level", flags.GetLogLevel(), "log level (debug|info|warn|error) ?")

		pflag.Usage = func() {
			fmt.Print(asciiConsole)
			fmt.Println("Usage: <service> [Flags]")
			pflag.PrintDefaults()
			os.Exit(0)
		}

		pflag.Parse()
	})

	return flags
}

// ParseAPIServer flag args for apiserver service
func ParseAPIServer() *celestinal.FlagAPIServer {
	onceAPIServer.Do(func() {
		pflag.BoolVarP(&apiServerFlags.Telegram, "telegram", "t", apiServerFlags.GetTelegram(), "turn on telegram system log ?")
		pflag.StringVar(&apiServerFlags.ApiSpecsPath, "api-specs", apiServerFlags.GetApiSpecsPath(), "openapi specification path ?")
		pflag.StringVar(&apiServerFlags.SwaggerPath, "swagger-ui", apiServerFlags.GetSwaggerPath(), "swagger ui path ?")
	})
	_ = Parse()

	return apiServerFlags
}

// SetDefault set default flag values
func SetDefault(name names.Namespace, address, mode string) {
	flags.Name = name.String()
	flags.Address = address
	flags.Mode = mode
}

// SetConsole set console log when --help or -h option
func SetConsole(ascii string) {
	asciiConsole = ascii
}
