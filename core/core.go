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
	}

	fmt.Println("Command executed successfully.")
}
