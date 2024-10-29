package core

import (
	"fmt"

	"github.com/matthewchivers/dodl/workspace"
)

type InitialiseCommand struct {
	AppCtx          *AppContext
	TargetDirectory string
}

func (c *InitialiseCommand) Execute() error {
	err := workspace.Initialise(c.TargetDirectory)
	if err != nil {
		return err
	}

	status := "Initialised new"
	if c.AppCtx.WorkspaceRoot == c.TargetDirectory {
		status = "Re-initialised"
	}
	fmt.Printf("%s dodl workspace at %s\n", status, c.TargetDirectory)

	return nil
}
