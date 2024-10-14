package cmd

import (
	"github.com/matthewchivers/dodl/core"
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
	cmdCtx := models.CommandContext{
		Command: "init",
		Args:    args,
		Flags: map[string]interface{}{
			"directory": directory,
		},
	}

	core.ExecuteCommand(cmdCtx)
}
