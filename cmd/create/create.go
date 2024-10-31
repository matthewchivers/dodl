package create

import (
	"fmt"
	"time"

	"github.com/matthewchivers/dodl/internal/core"
	"github.com/matthewchivers/dodl/pkg/config"
	wd "github.com/matthewchivers/dodl/pkg/workingdir"
	"github.com/matthewchivers/dodl/pkg/workspace"
	"github.com/spf13/cobra"
)

var (
	topic      string
	dateStr    string
	configPath string
	dryRun     bool
)

// NewCreateCmd initializes the 'create' command and its flags.
func NewCreateCmd(wdProv wd.WorkingDirProvider) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create [document type]",
		Short: "Create a new document",
		Long:  "Create a new document of the specified type using a predefined template.",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runCreateE(args, wdProv)
		},
	}

	cmd.Flags().StringVarP(&dateStr, "date", "d", "", "The date of the document (defaults to today's date)")
	cmd.Flags().BoolVar(&dryRun, "dry-run", false, "Perform a dry run (no changes made)")
	cmd.Flags().StringVarP(&topic, "topic", "t", "", "The topic of the document")
	cmd.Flags().StringVarP(&configPath, "config", "c", "", "The path to the configuration file")

	return cmd
}

// runCreateE executes the logic to create a new document.
func runCreateE(args []string, wdProv wd.WorkingDirProvider) error {
	workingDir, err := wdProv.GetWorkingDir()
	if err != nil {
		return fmt.Errorf("failed to get working directory: %w", err)
	}

	wsp, err := workspace.NewWorkspace(workingDir)
	if err != nil {
		return err
	}

	cfg, err := loadConfig(wsp.DodlPath())
	if err != nil {
		return err
	}

	docTypeName, err := getDocumentTypeName(args, cfg.DefaultDocumentType)
	if err != nil {
		return err
	}

	docType, exists := cfg.DocumentTypes[docTypeName]
	if !exists {
		return fmt.Errorf("document type %s not found in configuration", docTypeName)
	}

	startTime := time.Now()
	if dateStr != "" {
		startTime, err = time.Parse("02-01-2006", dateStr)
		if err != nil {
			return fmt.Errorf("failed to parse date: %w", err)
		}
	}

	appCtx := &core.AppContext{
		WorkingDir: workingDir,
		StartTime:  startTime,
	}

	createCmd := core.CreateCommand{
		DocName:      docTypeName,
		DocType:      docType,
		CustomFields: cfg.CustomFields,
		Topic:        topic,
		AppCtx:       appCtx,
		Workspace:    wsp,
	}

	return createCmd.Execute()
}

// loadConfig loads the configuration file from the workspace root for the given working directory.
// Returns the loaded configuration or an error if the configuration could not be loaded.
func loadConfig(dodlPath string) (*config.Config, error) {

	configOptions := config.ConfigOptions{
		CustomConfigFilePath: configPath,
		WorkspaceDodlDir:     dodlPath,
	}

	cfg, err := config.LoadConfigurations(configOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to load configurations: %w", err)
	}
	return cfg, nil
}

// getDocumentTypeName determines the document type name based on user input or config defaults.
// Returns the document type name or an error if no document type is specified (by argument or configuration default).
func getDocumentTypeName(args []string, defaultDocName string) (string, error) {
	if len(args) > 0 {
		return args[0], nil
	}

	if defaultDocName != "" {
		return defaultDocName, nil
	}

	return "", fmt.Errorf("no document type specified")
}
