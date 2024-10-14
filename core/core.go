package core

import (
	"github.com/matthewchivers/dodl/models"
)

type AppContext struct {
	cmdCtx models.CommandContext
}

func ExecuteCommand(cmdCtx models.CommandContext) {
	appCtx := AppContext{
		cmdCtx: cmdCtx,
	}

	switch appCtx.cmdCtx.Command {
	case "create":
		create(appCtx)
	case "init":
		initialise(appCtx)
	}
}
