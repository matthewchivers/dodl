package core

import (
	"github.com/matthewchivers/dodl/models"
)

type AppContext struct {
	cmdCtx *models.CommandContext
}

func ExecuteCommand(cmdCtx *models.CommandContext) error {
	appCtx := AppContext{
		cmdCtx: cmdCtx,
	}

	switch appCtx.cmdCtx.Command {
	case "create":
		return create(appCtx)
	case "init":
		return initialise(appCtx)
	case "status":
		return status(appCtx)
	}
	return nil
}
