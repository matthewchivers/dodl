package cmd

import (
	"fmt"

	"github.com/matthewchivers/dodl/cmd/wd"
	"github.com/matthewchivers/dodl/config"
	"github.com/matthewchivers/dodl/core"
	"github.com/matthewchivers/dodl/workspace"
	"github.com/spf13/cobra"
)

var (
	topic      string
	dateStr    string
	configPath string
	dryRun     bool
)

var createCmd = NewCreateCmd(&wd.DefaultWorkingDirProvider{})

func NewCreateCmd(wdProv wd.WorkingDirProvider) *cobra.Command {
	createCmd := cobra.Command{
		Use:   "create [document type]",
		Short: "Create a new document",
		Long:  `Create a new document of the specified type using a predefined template.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runCreateE(args, wdProv)
		},
	}

	createCmd.Flags().StringVarP(&dateStr, "date", "d", "", "The date of the document (defaults to today's date)")
	createCmd.Flags().BoolVar(&dryRun, "dry-run", false, "Perform a dry run (no changes made)")
	createCmd.Flags().StringVarP(&topic, "topic", "t", "", "The topic of the document")
	createCmd.Flags().StringVarP(&configPath, "config", "c", "", "The path to the configuration file")

	return &createCmd
}

func runCreateE(args []string, wdProv wd.WorkingDirProvider) error {
	workingDir, err := wdProv.GetWorkingDir()
	if err != nil {
		return err
	}

	workspaceRoot, err := workspace.FindWorkspaceRoot(workingDir)
	if err != nil {
		return err
	}

	// fetch config
	cfg, err := config.LoadConfigurations(config.ConfigOptions{
		CustomConfigFilePath: configPath,
		WorkspaceDodlDir:     workspace.GetDodlDirPath(workspaceRoot),
	})
	if err != nil {
		return err
	}

	// create app context
	appCtx := &core.AppContext{
		WorkingDir:    workingDir, // find if this needs to be the .dodl dir?
		WorkspaceRoot: workspaceRoot,
		StartTime:     startTime,
	}

	docTypeName := ""
	if len(args) > 0 {
		docTypeName = args[0]
	}

	if docTypeName == "" {
		if cfg.DefaultDocumentType != "" {
			docTypeName = cfg.DefaultDocumentType
		} else {
			return fmt.Errorf("no document type specified")
		}
	}

	docType, ok := cfg.DocumentTypes[docTypeName]
	if !ok {
		return fmt.Errorf("document type %s not found in configuration", docTypeName)
	}

	createCmd := &core.CreateCommand{
		DocName:      docTypeName,
		DocType:      docType,
		CustomFields: cfg.CustomValues,
		Topic:        topic,
		AppCtx:       appCtx,
	}

	return createCmd.Execute()
}
