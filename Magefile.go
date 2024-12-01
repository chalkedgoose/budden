//go:build mage
// +build mage

package main

import (
	"fmt"
	"os"
	"os/exec"
)

// Run starts the application.
func Run() error {
	fmt.Println("Running the application...")
	cmd := exec.Command("go", "run", "cmd/budden/main.go")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// Test runs the tests for the application.
func Test() error {
	fmt.Println("Running tests...")
	cmd := exec.Command("go", "test", "./...")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
