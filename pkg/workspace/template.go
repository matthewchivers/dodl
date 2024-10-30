package workspace

import (
	"os"
	"path/filepath"
)

// LoadTemplate reads a template file from the workspace.
// Returns the template data as a byte slice.
func (w *Workspace) LoadTemplate(templateFile string) ([]byte, error) {
	workspaceRoot := w.RootPath()

	dodlDir := ".dodl"
	templateDir := "templates"
	templatePath := filepath.Join(workspaceRoot, dodlDir, templateDir, templateFile)

	templateData, err := os.ReadFile(templatePath)
	if err != nil {
		return nil, err
	}

	return templateData, nil
}
