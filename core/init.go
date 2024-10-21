package core

import (
	"fmt"

	"github.com/matthewchivers/dodl/workspace"
)

type ErrAlreadyInWorkspace struct{}

func (e ErrAlreadyInWorkspace) Error() string {
	return "already in a workspace"
}

func initialise(appCtx AppContext) error {
	workingDir := appCtx.cmdCtx.EntryPoint
	// don't allow initialisation of a workspace within a workspace
	if _, err := workspace.FindWorkspaceRoot(workingDir); err != nil {
		if err != workspace.ErrNotInWorkspace {
			return err
		}
	} else {
		return ErrAlreadyInWorkspace{}
	}
	fmt.Printf("Initialising a new workspace in %s\n", appCtx.cmdCtx.Flags["directory"])
	return nil
}
