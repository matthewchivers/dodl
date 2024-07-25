/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/matthewchivers/dodl/handlers/create"
	"github.com/spf13/cobra"
)

var (
	createHandler = create.NewCreateCommandHandler()
)

// createCmd represents the new command
var createCmd = &cobra.Command{
	Use:   "create [type]",
	Short: "Create a new document of a specified type",
	Long: `Create a new document of a specified type.

Arguments:
- type: The specific type of document as specified in the config (e.g., note).
    - If no type is provided, a default will be used provided the config has specified one.

Examples:
- create note
  - Creates a new document of type note.
- create
  - Creates a new document using the default type.`,
	Run: createHandler.Handle,
}

func init() {
	rootCmd.AddCommand(createCmd)
}
