// Code generated by protoc-gen-go-tickex. DO NOT EDIT.
// Copyright 2025 Duc-Hung Ho

package authzcentre

import (
	"fmt"

	"github.com/tickexvn/tickex/api/gen/go/stdx/v1"
)

const ascii = `
 _______     __          
/_  __(_)___/ /_______ __	TICKEX // AUTHZCENTRE
 / / / / __/  '_/ -_) \ /	--------------
/_/ /_/\__/_/\_\\__/_\_\	tickex.authzcentre.v1

`

// PrintASCII the ASCII art to the console.
func PrintASCII() {
	fmt.Print(ascii)
}

// StdxOptsOfAuthzCentreService_SayHello get options from service method
func StdxOptsOfAuthzCentreService_SayHello() *stdx.Options {
	options := stdx.Options{}
	return &options
}

// StdxOptsOfAuthzCentreService_Status get options from service method
func StdxOptsOfAuthzCentreService_Status() *stdx.Options {
	options := stdx.Options{}
	return &options
}
