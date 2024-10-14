package core

import "fmt"

// create is a function that creates a new document.
func create(appCtx AppContext) {
	fmt.Println("Creating a new document...")
	fmt.Printf("Document type: %s\n", appCtx.cmdCtx.Args[0])
	fmt.Printf("Topic: %s\n", appCtx.cmdCtx.Flags["topic"])
	fmt.Printf("Date: %s\n", appCtx.cmdCtx.Flags["date"])
	fmt.Printf("Dry run: %t\n", appCtx.cmdCtx.Flags["dryRun"])
}
