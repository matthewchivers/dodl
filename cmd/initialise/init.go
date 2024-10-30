package initialise

import (
	"path/filepath"
	"time"

	"github.com/matthewchivers/dodl/cmd/wd"
	"github.com/matthewchivers/dodl/internal/core"
	"github.com/matthewchivers/dodl/workspace"
	"github.com/spf13/cobra"
)

// NewInitCmd creates the 'init' command for initializing a new dodl workspace.
func NewInitCmd(wdProv wd.WorkingDirProvider) *cobra.Command {
	return &cobra.Command{
		Use:   "init [directory]",
		Short: "Initialise a new dodl workspace",
		Long:  "Creates a new dodl workspace in the specified directory (defaults to the current working directory).",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runInitE(args, wdProv)
		},
	}
}

// runInitE executes the init command, setting up a new dodl workspace.
func runInitE(args []string, wdProv wd.WorkingDirProvider) error {
	workingDir, err := wdProv.GetWorkingDir()
	if err != nil {
		return err
	}

	targetDir, err := resolveTargetDir(workingDir, args)
	if err != nil {
		return err
	}

	workspaceRoot, err := locateWorkspaceRoot(workingDir)
	if err != nil {
		return err
	}

	appCtx := createAppContext(workingDir, workspaceRoot)
	initCmd := core.InitialiseCommand{
		AppCtx:          appCtx,
		TargetDirectory: targetDir,
	}

	return initCmd.Execute()
}

// resolveTargetDir determines the target directory for the workspace, defaulting to the working directory if not provided.
func resolveTargetDir(workingDir string, args []string) (string, error) {
	targetDir := workingDir
	if len(args) > 0 {
		targetDir = args[0]
	}

	return filepath.Abs(targetDir)
}

// locateWorkspaceRoot checks if the current working directory is part of an existing workspace and returns the root.
func locateWorkspaceRoot(workingDir string) (string, error) {
	workspaceRoot, err := workspace.FindWorkspaceRoot(workingDir)
	if err != nil && err != workspace.ErrNotInWorkspace {
		return "", err
	}
	return workspaceRoot, nil
}

// createAppContext initializes the core.AppContext with the provided working directory and workspace root.
func createAppContext(workingDir, workspaceRoot string) *core.AppContext {
	startTime := time.Now()
	return &core.AppContext{
		WorkingDir:    workingDir,
		StartTime:     startTime,
		WorkspaceRoot: workspaceRoot,
	}
}
