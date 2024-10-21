package cmd

import (
	"github.com/matthewchivers/dodl/cmd/wd"
	"github.com/matthewchivers/dodl/core"
	"github.com/matthewchivers/dodl/models"
	"github.com/spf13/cobra"
)

var (
	topic   string
	dateStr string
	dryRun  bool
)

var createCmd = NewCreateCmd(&wd.DefaultWorkingDirProvider{})

func NewCreateCmd(wdProv wd.WorkingDirProvider) *cobra.Command {
	return &cobra.Command{
		Use:   "create [document type]",
		Short: "Create a new document",
		Long:  `Create a new document of the specified type using a predefined template.`,
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runCreateE(args, wdProv)
		},
	}
}

func init() {
	createCmd.Flags().StringVarP(&dateStr, "date", "d", "", "The date of the document (defaults to today's date)")
	createCmd.Flags().BoolVar(&dryRun, "dry-run", false, "Perform a dry run (no changes made)")
	createCmd.Flags().StringVarP(&topic, "topic", "t", "", "The topic of the document")
}

func runCreateE(args []string, wdProv wd.WorkingDirProvider) error {
	workingDir, err := wdProv.GetWorkingDir()
	if err != nil {
		return err
	}

	docType := args[0]

	cmdCtx := &models.CommandContext{
		Command: "create",
		Args:    args,
		Flags: map[string]interface{}{
			"document_type": docType,
			"topic":         topic,
			"date":          dateStr,
			"dryRun":        dryRun,
		},
		EntryPoint: workingDir,
	}

	return core.ExecuteCommand(cmdCtx)
}
