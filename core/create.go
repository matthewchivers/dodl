package core

import (
	"fmt"

	"github.com/matthewchivers/dodl/config"
)

type CreateCommand struct {
	DocName string
	DocType config.DocumentType
	Topic   string
	AppCtx  *AppContext
}

func (c *CreateCommand) Execute() error {
	fmt.Println("Creating a new document...")
	fmt.Printf("Start time: %s\n", c.AppCtx.StartTime)
	fmt.Printf("Working directory: %s\n", c.AppCtx.WorkingDir)
	fmt.Printf("Workspace root: %s\n", c.AppCtx.WorkspaceRoot)
	fmt.Printf("Document type: %s\n", c.DocType)
	fmt.Printf("Topic: %s\n", c.Topic)
	docType := c.DocType
	fmt.Printf("Document filename pattern: %v\n", docType.FileNamePattern)
	fmt.Printf("Document directory pattern: %v\n", docType.DirectoryPattern)
	fmt.Printf("Document template: %v\n", docType.TemplateFile)
	fmt.Println("Custom values:")
	for k, v := range docType.CustomValues {
		fmt.Printf(" - Custom value: %s = %v\n", k, v)
	}
	return nil
}
