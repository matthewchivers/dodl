package cmd

import (
	"fmt"

	"github.com/matthewchivers/dodl/models"
	"github.com/spf13/cobra"
)

var (
	topic   string
	dateStr string
	dryRun  bool
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create [document_type]",
	Short: "Create a new document",
	Long:  `Create a new document of the specified type using a predefined template.`,
	Args:  cobra.MinimumNArgs(1),
	Run:   runCreate,
}

func init() {
	createCmd.Flags().StringVarP(&topic, "topic", "t", "", "The topic of the document")
	createCmd.Flags().StringVarP(&dateStr, "date", "d", "", "The date of the document (defaults to today's date)")
	createCmd.Flags().BoolVar(&dryRun, "dry-run", false, "Perform a dry run (no changes made)")
}

func runCreate(cmd *cobra.Command, args []string) {
	cmdCtx := models.CommandContext{
		Command: "create",
		Args:    args,
		Flags: map[string]interface{}{
			"topic":  topic,
			"date":   dateStr,
			"dryRun": dryRun,
		},
	}

	// call core logic once implemented
	fmt.Printf("Command: %s\n", cmdCtx.Command)
	fmt.Println("Creating a new document...")
	fmt.Printf("Document type: %s\n", cmdCtx.Args[0])
	fmt.Printf("Topic: %s\n", cmdCtx.Flags["topic"])
	fmt.Printf("Date: %s\n", cmdCtx.Flags["date"])
	fmt.Printf("Dry run: %t\n", cmdCtx.Flags["dryRun"])
}
