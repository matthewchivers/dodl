package core

import (
	"fmt"

	"github.com/matthewchivers/dodl/workspace"
)

func initialise(appCtx AppContext) error {
	targetDirectory := appCtx.cmdCtx.Flags["targetDirectory"].(string)

	// Determine if the target directory is already the root of a workspace
	root, err := workspace.FindWorkspaceRoot(targetDirectory)
	if err != nil && err != workspace.ErrNotInWorkspace {
		return err
	}

	isReinitialise := false
	if err == nil && root == targetDirectory {
		// The target directory is already a workspace root
		isReinitialise = true
	}

	// Proceed with initialization (or re-initialization)
	err = workspace.Initialise(targetDirectory)
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
