//go:build mage
// +build mage

package main

import (
	"fmt"
	"github.com/magefile/mage/sh"
)

// Default target to run when none is specified
// If not set, running mage will list available targets
// var Default = Build

const BUFF_VERSION = "v1.15.0"

func InstallBuf() error {
	println("Installing Buf ...")
	err := sh.RunV("go", "install", fmt.Sprintf("github.com/bufbuild/buf/cmd/buf@%s", BUFF_VERSION))
	if err != nil {
		return logAndReturn(err)
	}
	return nil
}

/*

go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2
go install google.golang.org/protobuf/cmd/protoc-gen-go
go install google.golang.org/grpc/cmd/protoc-gen-go-grp

npm install protoc-gen-ts
npm install grpc-web
npm install google-protobuf

*/

func logAndReturn(err error) error {
	println(err)
	return err
}
