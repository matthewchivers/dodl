package workspace

import (
	"path/filepath"

	"github.com/matthewchivers/dodl/filesystem"
)

func Initialise(targetDir string) error {
	workspaceRoot := getRootPath(targetDir)

	err := createWorkspaceRoot(targetDir)
	if err != nil {
		return err
	}

	err = createTemplatesDir(workspaceRoot)
	if err != nil {
		return err
	}

	return nil
}

func getRootPath(targetDir string) string {
	return filepath.Join(targetDir, ".dodl")
}

func createWorkspaceRoot(targetDir string) error {
	workspaceRootPath := filepath.Join(targetDir, ".dodl")
	return filesystem.MkDir(workspaceRootPath)
}

func createTemplatesDir(workspaceRoot string) error {
	templatesDirPath := filepath.Join(workspaceRoot, "templates")
	return filesystem.MkDir(templatesDirPath)
}
