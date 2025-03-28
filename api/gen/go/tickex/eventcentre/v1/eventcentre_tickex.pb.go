// Code generated by protoc-gen-go-tickex. DO NOT EDIT.
// Copyright 2025 Duc-Hung Ho

package eventcentre

import (
	"fmt"

	"github.com/tickexvn/tickex/api/gen/go/tickex/v1"
)

var (
	_ tickex.Empty
)

const ascii = `
 _______     __          
/_  __(_)___/ /_______ __	TICKEX // EVENTCENTRE
 / / / / __/  '_/ -_) \ /	--------------
/_/ /_/\__/_/\_\\__/_\_\	tickex.eventcentre.v1

`

// PrintASCII the ASCII art to the console.
func PrintASCII() {
	fmt.Print(ascii)
}

// HasRoleAtEventCentreService_SayHello checks if the role has access to the method
func HasRoleAtEventCentreService_SayHello(role tickex.Role) bool {
	return true
}

// HasRoleAtEventCentreService_Status checks if the role has access to the method
func HasRoleAtEventCentreService_Status(role tickex.Role) bool {
	return true
}
