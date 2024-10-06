package cmd

import (
	"fmt"

	"github.com/matthewchivers/dodl/models"
	"github.com/spf13/cobra"
)

var directory string

var initCmd = &cobra.Command{
	Use:   "init [directory]",
	Short: "Initialise a new dodl workspace",
	Long:  `Creates a new dodl workspace in the specified directory (defaults to current).`,
	Run:   runInit,
}

func runInit(cmd *cobra.Command, args []string) {
	// Command context preparation
	cmdCtx := models.CommandContext{
		Command: "init",
		Args:    args,
		Flags: map[string]interface{}{
			"directory": directory,
		},
	}

	// call core logic once implemented
	fmt.Println("Initialising a new workspace...")
	if len(cmdCtx.Args) == 0 {
		fmt.Println("No directory specified, using current directory.")
	} else {
		fmt.Printf("Directory: %s\n", cmdCtx.Args[0])
	}
}
