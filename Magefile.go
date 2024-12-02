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

func Migrate() error {
	fmt.Println("Migrating the database...")
	cmd := exec.Command("goose", "-dir", "data/sql/migrations", "postgres", "user=postgres password=mysecretpassword host=localhost dbname=totp_practice sslmode=disable", "up")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
