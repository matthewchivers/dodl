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
	fmt.Printf("Command: %s\n", cmdCtx.Command)
	if len(cmdCtx.Args) == 0 {
		fmt.Println("No directory specified, using current directory.")
	} else {
		cmdCtx.Flags["directory"] = cmdCtx.Args[0]
		fmt.Printf("Directory: %s\n", cmdCtx.Args[0])
	}
	fmt.Printf("Directory Flag: %s\n", cmdCtx.Flags["directory"])
	fmt.Println("Initialising a new workspace...")
}
