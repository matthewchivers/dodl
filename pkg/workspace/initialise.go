package workspace

import (
	"path/filepath"

	"github.com/matthewchivers/dodl/pkg/filesystem"
)

func Initialise(targetDir string) error {
	dodlDir := filepath.Join(targetDir, ".dodl")

	err := filesystem.EnsureDirExists(dodlDir)
	if err != nil {
		return err
	}

	configYamlPath := filepath.Join(dodlDir, "config.yaml")
	err = filesystem.EnsureFileExists(configYamlPath, []byte(""))
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
