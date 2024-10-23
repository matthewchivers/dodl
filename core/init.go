package core

import (
	"fmt"

	"github.com/matthewchivers/dodl/workspace"
)

func initialise(appCtx AppContext) error {
	targetDirectory := appCtx.CmdCtx.Flags["targetDirectory"].(string)

	// Determine if the target directory is already the root of a workspace
	isReinitialise := false
	if appCtx.WorkspaceRoot == targetDirectory {
		isReinitialise = true
	}

	// Proceed with initialization (or re-initialization)
	err := workspace.Initialise(targetDirectory)
	if err != nil {
		return err
	}

	// Output the appropriate message
	if isReinitialise {
		fmt.Printf("Re-initialised dodl workspace at %s\n", targetDirectory)
	} else {
		fmt.Printf("Initialised new dodl workspace at %s\n", targetDirectory)
	}

	return nil
}
