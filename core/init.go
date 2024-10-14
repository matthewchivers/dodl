package core

import "fmt"

func initialise(appCtx AppContext) {
	if len(appCtx.cmdCtx.Args) == 0 {
		fmt.Println("No directory specified, using current directory.")
	} else {
		appCtx.cmdCtx.Flags["directory"] = appCtx.cmdCtx.Args[0]
		fmt.Printf("Directory: %s\n", appCtx.cmdCtx.Args[0])
	}
	fmt.Printf("Directory Flag: %s\n", appCtx.cmdCtx.Flags["directory"])
	fmt.Println("Initialising a new workspace...")
}
