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

// Package internal provides the logger for the package.
package internal

import (
	"fmt"

	"github.com/celestinals/celestinal/internal/utils/version"

	cestcolor "github.com/celestinals/celestinal/pkg/color"
	"go.uber.org/zap"
)

const (
	// LevelDebug is the lowest level of verbosity.
	LevelDebug = 0

	// LevelInfo is the default level of verbosity.
	LevelInfo = 1

	// LevelWarning is the default level of verbosity.
	LevelWarning = 2

	// LevelError is the default level of verbosity.
	LevelError = 3

	// LevelFatal is the highest level of verbosity.
	LevelFatal = 4
)

var header = fmt.Sprintf("%s > ", cestcolor.Green.Add(fmt.Sprintf("[%s]", version.Name)))

// TxLogCore is the logger for the package.
type TxLogCore struct {
	Verbosity int
	Logger    *zap.SugaredLogger
}

func (c *TxLogCore) addPrefix(args ...any) []any {
	return append([]any{header}, args...)
}

func (c *TxLogCore) addPrefixFormat(format string) string {
	return header + format
}

// Debug logs a debug message.
func (c *TxLogCore) Debug(args ...any) {
	if c.V(LevelDebug) {
		c.Logger.Debug(c.addPrefix(args...)...)
	}
}

// Debugln logs a debug message.
func (c *TxLogCore) Debugln(args ...any) { c.Debug(args...) }

// Debugf logs a debug message with a format.
func (c *TxLogCore) Debugf(format string, args ...any) {
	if c.V(LevelDebug) {
		c.Logger.Debugf(c.addPrefixFormat(format), args...)
	}
}

// Info logs an info message.
func (c *TxLogCore) Info(args ...any) {
	if c.V(LevelInfo) {
		c.Logger.Info(c.addPrefix(args...)...)
	}
}

// Infoln logs an info message.
func (c *TxLogCore) Infoln(args ...any) {
	c.Info(args...)
}

// Infof logs an info message with a format.
func (c *TxLogCore) Infof(format string, args ...any) {
	if c.V(LevelInfo) {
		c.Logger.Infof(c.addPrefixFormat(format), args...)
	}
}

// Warning logs a warning message.
func (c *TxLogCore) Warning(args ...any) {
	if c.V(LevelWarning) {
		c.Logger.Warn(c.addPrefix(args...)...)
	}
}

// Warningln logs a warning message.
func (c *TxLogCore) Warningln(args ...any) {
	c.Warning(args...)
}

// Warningf logs a warning message with a format.s
func (c *TxLogCore) Warningf(format string, args ...any) {
	if c.V(LevelWarning) {
		c.Logger.Warnf(c.addPrefixFormat(format), args...)
	}
}

// Error logs an error message.
func (c *TxLogCore) Error(args ...any) {
	if c.V(LevelError) {
		c.Logger.Error(c.addPrefix(args...)...)
	}
}

// Errorln logs an error message.
func (c *TxLogCore) Errorln(args ...any) {
	c.Error(args...)
}

// Errorf logs an error message with a format.
func (c *TxLogCore) Errorf(format string, args ...any) {
	if c.V(LevelError) {
		c.Logger.Errorf(c.addPrefixFormat(format), args...)
	}
}

// Fatal logs a fatal message.
func (c *TxLogCore) Fatal(args ...any) {
	if c.V(LevelFatal) {
		c.Logger.Fatal(c.addPrefix(args...)...)
	}
}

// Fatalln logs a fatal message.
func (c *TxLogCore) Fatalln(args ...any) { c.Fatal(args...) }

// Fatalf logs a fatal message with a format.
func (c *TxLogCore) Fatalf(format string, args ...any) {
	if c.V(LevelFatal) {
		c.Logger.Fatalf(c.addPrefixFormat(format), args...)
	}
}

// V reports whether verbosity level l is at least the requested verbose level.
func (c *TxLogCore) V(l int) bool {
	return l >= c.Verbosity
}

// Sync flushes the log.
func (c *TxLogCore) Sync() error {
	return c.Logger.Sync()
}

// NewTxLogCore creates a new TxLogCore.
func NewTxLogCore(logger *zap.SugaredLogger, verbosity int) *TxLogCore {
	return &TxLogCore{Logger: logger, Verbosity: verbosity}
}
