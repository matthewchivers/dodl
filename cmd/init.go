package cmd

import (
	"path/filepath"

	"github.com/matthewchivers/dodl/cmd/wd"
	"github.com/matthewchivers/dodl/core"
	"github.com/matthewchivers/dodl/models"
	"github.com/spf13/cobra"
)

var initCmd = NewInitCmd(&wd.DefaultWorkingDirProvider{})

func NewInitCmd(wdProv wd.WorkingDirProvider) *cobra.Command {
	return &cobra.Command{
		Use:   "init [directory]",
		Short: "Initialise a new dodl workspace",
		Long:  `Creates a new dodl workspace in the specified directory (defaults to current working directory).`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runInitE(args, wdProv)
		},
	}
}

func runInitE(args []string, wdProv wd.WorkingDirProvider) error {
	workingDir, err := wdProv.GetWorkingDir()
	if err != nil {
		return err
	}

	targetDir, err := getTargetDir(workingDir, args)
	if err != nil {
		return err
	}

	cmdCtx := &models.CommandContext{
		Command: "init",
		Args:    args,
		Flags: map[string]interface{}{
			"targetDirectory": targetDir,
		},
		EntryPoint: workingDir,
	}

	return core.ExecuteCommand(cmdCtx)
}

func getTargetDir(wd string, args []string) (string, error) {
	targetDirectory := wd
	if len(args) > 0 {
		targetDirectory = args[0]
	}

	absTargetDir, err := filepath.Abs(targetDirectory)
	if err != nil {
		return "", err
	}
	return absTargetDir, nil
}
