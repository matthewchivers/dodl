package workspace

import (
	"path/filepath"

	"github.com/matthewchivers/dodl/pkg/filesystem"
)

// Initialise ensures a dodl workspace exists in the target directory.
// The workspace will contain a .dodl directory with a config.yaml file and a templates directory.
// It does this idempotently, so it will not overwrite any existing files or directories.
func Initialise(targetDir string) error {
	dodlDir := filepath.Join(targetDir, ".dodl")

	err := filesystem.EnsureDirExists(dodlDir)
	if err != nil {
		return err
	}

	configYamlPath := filepath.Join(dodlDir, "config.yaml")
	_, err = filesystem.EnsureFileExists(configYamlPath, []byte(""))
	if err != nil {
		return err
	}

	templatesDirPath := filepath.Join(dodlDir, "templates")
	err = filesystem.EnsureDirExists(templatesDirPath)
	if err != nil {
		return err
	}

	return nil
}
