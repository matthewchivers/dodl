package core

import (
	"fmt"

	"github.com/matthewchivers/dodl/pkg/config"
)

type StatusCommand struct {
	AppCtx *AppContext
	Config *config.Config
}

func (c *StatusCommand) Execute() error {
	fmt.Printf("In workspace: %s\n", c.AppCtx.WorkspaceRoot)
	fmt.Printf("Working directory: %s\n", c.AppCtx.WorkingDir)
	fmt.Printf("Start time: %s\n", c.AppCtx.StartTime)
	fmt.Printf("Config:\n")
	if c.Config.DefaultDocumentType != "" {
		fmt.Printf("  Default document type: %s\n", c.Config.DefaultDocumentType)
	}
	for docName, dt := range c.Config.DocumentTypes {
		fmt.Printf("  Document type: %s\n", docName)
		fmt.Printf("    Filename Pattern: %s\n", dt.FileNamePattern)
		fmt.Printf("    Directory Pattern: %s\n", dt.DirectoryPattern)
		fmt.Printf("    Template Path: %s\n", dt.TemplateFile)
		for k, v := range dt.CustomFields {
			fmt.Printf("      Custom Field: %s: %s\n", k, v)
		}
	}
	return nil
}
