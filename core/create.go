package core

import "fmt"

// create is a function that creates a new document.
func create(appCtx AppContext) error {
	cmdCtx := appCtx.cmdCtx
	fmt.Println("Creating a new document...")
	fmt.Printf("Document type: %s\n", cmdCtx.Flags["document_type"])
	fmt.Printf("Topic: %s\n", cmdCtx.Flags["topic"])
	fmt.Printf("Date: %s\n", cmdCtx.Flags["date"])
	fmt.Printf("Dry run: %t\n", cmdCtx.Flags["dryRun"])
	return nil
}
