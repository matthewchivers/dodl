package cmd

import (
	"github.com/matthewchivers/dodl/cmd/wd"
	"github.com/matthewchivers/dodl/core"
	"github.com/matthewchivers/dodl/models"
	"github.com/spf13/cobra"
)

var statusCmd = NewStatusCmd(&wd.DefaultWorkingDirProvider{})

func NewStatusCmd(wdProv wd.WorkingDirProvider) *cobra.Command {
	return &cobra.Command{
		Use:   "status",
		Short: "Show the status of the current workspace",
		Long:  `Show the status of the current workspace.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runStatusE(wdProv)
		},
	}
}

func runStatusE(wdProv wd.WorkingDirProvider) error {
	workingDir, err := wdProv.GetWorkingDir()
	if err != nil {
		return err
	}

	cmdCtx := &models.CommandContext{
		Command:    "status",
		EntryPoint: workingDir,
	}

	return core.ExecuteCommand(cmdCtx)
}
