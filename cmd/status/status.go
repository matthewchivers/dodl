package status

import (
	"time"

	"github.com/matthewchivers/dodl/cmd/wd"
	"github.com/matthewchivers/dodl/config"
	"github.com/matthewchivers/dodl/core"
	"github.com/matthewchivers/dodl/workspace"
	"github.com/spf13/cobra"
)

// NewStatusCmd creates the 'status' command for checking the workspace status.
func NewStatusCmd(wdProv wd.WorkingDirProvider) *cobra.Command {
	return &cobra.Command{
		Use:   "status",
		Short: "Show the status of the current workspace",
		Long:  "Displays the current configuration and status of the dodl workspace.",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runStatusE(wdProv)
		},
	}
}

// runStatusE executes the status command logic, checking workspace status.
func runStatusE(wdProv wd.WorkingDirProvider) error {
	workingDir, err := wdProv.GetWorkingDir()
	if err != nil {
		return err
	}

	workspaceRoot, err := workspace.FindWorkspaceRoot(workingDir)
	if err != nil {
		return err
	}

	appCtx := createAppContext(workingDir, workspaceRoot)
	cfg, err := loadConfig(workspaceRoot)
	if err != nil {
		return err
	}

	return executeStatus(appCtx, cfg)
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

// loadConfig loads the dodl configuration from the specified workspace root.
func loadConfig(workspaceRoot string) (*config.Config, error) {
	cfgOpts := config.ConfigOptions{
		WorkspaceDodlDir: workspaceRoot,
	}
	return config.LoadConfigurations(cfgOpts)
}

// executeStatus runs the StatusCommand to display the workspace status.
func executeStatus(appCtx *core.AppContext, cfg *config.Config) error {
	statusCmd := core.StatusCommand{
		AppCtx: appCtx,
		Config: cfg,
	}
	return statusCmd.Execute()
}