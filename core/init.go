package core

import "fmt"

func initialise(appCtx AppContext) error {
	fmt.Printf("Initialising a new workspace in %s\n", appCtx.cmdCtx.Flags["directory"])
	return nil
}
