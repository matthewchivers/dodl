package core

import (
	"fmt"

	"github.com/matthewchivers/dodl/pkg/workspace"
)

// InitialiseCommand holds all the data/context required to initialise a new workspace
type InitialiseCommand struct {
	AppCtx          *AppContext
	TargetDirectory string
}

// Execute initialises a new workspace based on the data/context in the InitialiseCommand
func (c *InitialiseCommand) Execute() error {
	err := workspace.Initialise(c.TargetDirectory)
	if err != nil {
		return err
	}

	wsp, err := workspace.NewWorkspace(c.TargetDirectory)
	if err != nil {
		return err
	}
	status := "Initialised new"
	if wsp.RootPath() == c.TargetDirectory {
		status = "Re-initialised"
	}
	fmt.Printf("%s dodl workspace at %s\n", status, wsp.RootPath())

	return nil
}
