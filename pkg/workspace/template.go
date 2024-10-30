package workspace

import (
	"os"
	"path/filepath"
)

func LoadTemplate(workingDirectory, templateFile string) ([]byte, error) {
	// Get workspace root / dodl directory
	workspaceRoot, err := FindWorkspaceRoot(workingDirectory)
	if err != nil {
		return nil, err
	}

	// Define the .dodl directory and template subdirectory
	dodlDir := ".dodl"
	templateDir := "templates"
	templatePath := filepath.Join(workspaceRoot, dodlDir, templateDir, templateFile)

	// Load the template file from the constructed path
	templateData, err := os.ReadFile(templatePath)
	if err != nil {
		return nil, err
	}

	return templateData, nil
}
