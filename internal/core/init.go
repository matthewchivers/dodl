package core

import (
	"fmt"

	"github.com/matthewchivers/dodl/pkg/workspace"
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
	workspaceRoot, err := workspace.FindWorkspaceRoot(c.TargetDirectory)
	if err != nil {
		return err
	}
	if workspaceRoot == c.TargetDirectory {
		status = "Re-initialised"
	}
	fmt.Printf("%s dodl workspace at %s\n", status, c.TargetDirectory)

	return nil
}
