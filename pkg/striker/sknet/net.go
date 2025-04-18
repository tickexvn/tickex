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

// Package sknet provide functions extended of http/net
package sknet

import (
	"net"
	"strconv"
)

// ListenNetworkTCP listens on the TCP network address addr and
// returns a net.Listener.
func ListenNetworkTCP(addr string) (net.Listener, error) {
	return net.Listen("tcp", addr)
}

// SplitHostPortListener return host, port of the net.Listener instance
func SplitHostPortListener(listener net.Listener) (string, uint32, error) {
	host, port, err := net.SplitHostPort(listener.Addr().String())
	if err != nil {
		return "", 0, err
	}

	portNumber, err := strconv.Atoi(port)
	if err != nil {
		return "", 0, err
	}

	return host, uint32(portNumber), nil
}
