package core

import (
	"fmt"

	"github.com/matthewchivers/dodl/models"
)

type AppContext struct {
	cmdCtx models.CommandContext
}

func ExecuteCommand(cmdCtx models.CommandContext) {
	appCtx := AppContext{
		cmdCtx: cmdCtx,
	}

	fmt.Printf("Command: %s\n", appCtx.cmdCtx.Command)

	switch cmdCtx.Command {
	case "create":
		fmt.Println("Creating a new document...")
		fmt.Printf("Document type: %s\n", appCtx.cmdCtx.Args[0])
		fmt.Printf("Topic: %s\n", appCtx.cmdCtx.Flags["topic"])
		fmt.Printf("Date: %s\n", appCtx.cmdCtx.Flags["date"])
		fmt.Printf("Dry run: %t\n", appCtx.cmdCtx.Flags["dryRun"])
	case "init":
		if len(cmdCtx.Args) == 0 {
			fmt.Println("No directory specified, using current directory.")
		} else {
			cmdCtx.Flags["directory"] = cmdCtx.Args[0]
			fmt.Printf("Directory: %s\n", cmdCtx.Args[0])
		}
		fmt.Printf("Directory Flag: %s\n", cmdCtx.Flags["directory"])
		fmt.Println("Initialising a new workspace...")
	}

	fmt.Println("Command executed successfully.")
}
