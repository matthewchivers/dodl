package core

import (
	"github.com/matthewchivers/dodl/config"
	"github.com/matthewchivers/dodl/models"
)

type AppContext struct {
	CmdCtx        *models.CommandContext
	WorkspaceRoot string
	Config        *config.Config
}

func (c *AppContext) LoadConfigurations(opts config.ConfigOptions) error {
	cfg, err := config.LoadConfigurations(opts)
	if err != nil {
		return err
	}
	c.Config = cfg
	return nil
}
