package status

import (
	"time"

	"github.com/matthewchivers/dodl/cmd/wd"
	"github.com/matthewchivers/dodl/internal/core"
	"github.com/matthewchivers/dodl/pkg/config"
	"github.com/matthewchivers/dodl/pkg/workspace"
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

	wsp, err := workspace.NewWorkspace(workingDir)
	if err != nil {
		return err
	}

	appCtx := createAppContext(workingDir)
	cfg, err := loadConfig(workingDir)
	if err != nil {
		return err
	}

	return executeStatus(appCtx, cfg, wsp)
}

// createAppContext initializes the core.AppContext with the provided working directory and workspace root.
func createAppContext(workingDir string) *core.AppContext {
	startTime := time.Now()
	return &core.AppContext{
		WorkingDir: workingDir,
		StartTime:  startTime,
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
func executeStatus(appCtx *core.AppContext, cfg *config.Config, wsp *workspace.Workspace) error {
	statusCmd := core.StatusCommand{
		AppCtx:    appCtx,
		Config:    cfg,
		Workspace: wsp,
	}
	return statusCmd.Execute()
}
