package core

import (
	"github.com/matthewchivers/dodl/config"
	"github.com/matthewchivers/dodl/models"
	"github.com/matthewchivers/dodl/workspace"
)

func ExecuteCommand(cmdCtx *models.CommandContext) error {
	workspaceRoot, err := workspace.FindWorkspaceRoot(cmdCtx.EntryPoint)
	if err != nil && err != workspace.ErrNotInWorkspace {
		return err
	}

	appCtx := AppContext{
		CmdCtx:        cmdCtx,
		WorkspaceRoot: workspaceRoot,
	}

	cfgOpts := config.ConfigOptions{
		WorkspaceDodlDir: workspaceRoot,
	}

	switch appCtx.CmdCtx.Command {
	case "create":
		err := appCtx.LoadConfigurations(cfgOpts)
		if err != nil {
			return err
		}
		return create(appCtx)
	case "init":
		return initialise(appCtx)
	case "status":
		err := appCtx.LoadConfigurations(cfgOpts)
		if err != nil {
			return err
		}
		return status(appCtx)
	}
	return nil
}
