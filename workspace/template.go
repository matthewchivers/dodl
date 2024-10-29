package workspace

import (
	"os"
	"path/filepath"
)

func LoadTemplate(entryPoint, templateFile string) ([]byte, error) {
	// Get workspace root / dodl directory
	workspaceRoot, err := FindWorkspaceRoot(entryPoint)
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
