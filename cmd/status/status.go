package status

import (
	"time"

	"github.com/matthewchivers/dodl/internal/core"
	"github.com/matthewchivers/dodl/pkg/config"
	wd "github.com/matthewchivers/dodl/pkg/workingdir"
	"github.com/matthewchivers/dodl/pkg/workspace"
	"github.com/spf13/cobra"
)

// NewStatusCmd returns the configured cobra.Command for the status command.
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

// createAppContext initialises the core.AppContext with the provided working directory and workspace root.
// Returns a new AppContext with the working directory and start time set.
func createAppContext(workingDir string) *core.AppContext {
	startTime := time.Now()
	return &core.AppContext{
		WorkingDir: workingDir,
		StartTime:  startTime,
	}
}

// loadConfig loads the dodl configuration from the specified workspace root.
// Returns the loaded configuration, or an error if the configuration cannot be loaded.
func loadConfig(workspaceRoot string) (*config.Config, error) {
	cfgOpts := config.ConfigOptions{
		WorkspaceDodlDir: workspaceRoot,
	}
	return config.LoadConfigurations(cfgOpts)
}

// executeStatus runs the StatusCommand to display the workspace status.
// Returns an error if the status command fails to execute.
func executeStatus(appCtx *core.AppContext, cfg *config.Config, wsp *workspace.Workspace) error {
	statusCmd := core.StatusCommand{
		AppCtx:    appCtx,
		Config:    cfg,
		Workspace: wsp,
	}
	return statusCmd.Execute()
}
