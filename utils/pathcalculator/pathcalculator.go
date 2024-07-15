package pathcalculator

import (
	"os"
	"path/filepath"
)

const (
	// dodlDirName is the name of the directory for the workspace
	dodlDirName = ".dodl"

	// configFile is the name of the configuration file for the workspace
	configFile = "config.yaml"

	// metaFile is the name of the meta file for the workspace
	metaFile = "dodl_meta.json"
)

// CalculateDodlDir calculates the path to the .dodl directory
func CalculateDodlDir(workspaceDir string) string {
	return filepath.Join(workspaceDir, dodlDirName)
}

// CalculateConfigFile calculates the path to the configuration file
func CalculateConfigFile(dodlDir string) string {
	return filepath.Join(dodlDir, configFile)
}

// CalculateMetaFile calculates the path to the meta file
func CalculateMetaFile(dodlDir string) string {
	return filepath.Join(dodlDir, metaFile)
}

// CalculateUserDodlDir calculates the path to the user's .dodl directory
func CalculateUserDodlDir() (string, error) {
	userDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(userDir, dodlDirName), nil
}
