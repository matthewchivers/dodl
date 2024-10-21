package core

import (
	"fmt"

	"github.com/matthewchivers/dodl/workspace"
)

func status(appCtx AppContext) error {
	root, err := workspace.FindWorkspaceRoot(appCtx.cmdCtx.EntryPoint)
	if err != nil {
		if err == workspace.ErrNotInWorkspace {
			fmt.Println("Not in a dodl workspace")
			return nil
		}
		return err
	}
	fmt.Printf("In workspace: %s\n", root)
	fmt.Printf("Working directory: %s\n", appCtx.cmdCtx.EntryPoint)

	return nil
}
