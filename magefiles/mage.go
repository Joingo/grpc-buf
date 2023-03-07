//go:build mage
// +build mage

package main

import (
	"fmt"
	"os/exec"

	"github.com/magefile/mage/mg" // mg contains helpful utility functions, like Deps
	"github.com/magefile/mage/sh"
)

// Default target to run when none is specified
// If not set, running mage will list available targets
// var Default = Build

const BUFF_VERSION = "v1.15.0"

// A build step that requires additional params, or platform specific steps for example
func Build() error {
	mg.Deps(InstallDeps)
	fmt.Println("Building...")
	cmd := exec.Command("go", "build", "-o", "MyApp", ".")
	return cmd.Run()
}

func InstallDeps() error {
	return nil
}

func InstallBuf() error {
	println("Installing Buf ...")
	err := sh.RunV("go", "install", fmt.Sprintf("github.com/bufbuild/buf/cmd/buf@%s", BUFF_VERSION))
	if err != nil {
		return logAndReturn(err)
	}
	return nil
}

func logAndReturn(err error) error {
	println(err)
	return err
}
