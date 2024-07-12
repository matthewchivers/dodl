package main

import (
	"fmt"
	"os"

	"github.com/matthewchivers/dodl/cmd"
)

// dodl is a cli tool for file management
func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
