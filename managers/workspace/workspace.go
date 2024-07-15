package workspace

import (
	"os"

	"github.com/matthewchivers/dodl/utils/pathcalculator"
	"github.com/matthewchivers/dodl/utils/pathfinder"
)

type WorkspaceManager struct {
	// workspaceDir is the path to the current workspace (parent of the .dodl directory)
	workspaceDir string

	// dodlDir is the path to the .dodl directory
	dodlDir string
}

func NewWorkspaceManager() *WorkspaceManager {
	return &WorkspaceManager{}
}

type AlreadyInWorkspaceError struct{}

func (e *AlreadyInWorkspaceError) Error() string {
	return "already in a workspace"
}

// InitialiseWorkspace creates a new dodl workspace at the current location
func (m *WorkspaceManager) InitialiseWorkspace(path string) error {
	if _, err := pathfinder.FindWorkspaceRootDir(path); err == nil {
		return &AlreadyInWorkspaceError{}
	}
	m.workspaceDir = path

	m.dodlDir = pathcalculator.CalculateDodlDir(path)

	if err := os.Mkdir(m.dodlDir, 0755); err != nil {
		return err
	}

	if err := m.createConfigFile(m.dodlDir); err != nil {
		return err
	}

	if err := m.createMetaFile(m.dodlDir); err != nil {
		return err
	}
	return nil
}

// createConfigFile creates the configuration file in the specified directory
func (m *WorkspaceManager) createConfigFile(dodlDir string) error {
	configPath := pathcalculator.CalculateConfigFile(dodlDir)
	file, err := os.Create(configPath)
	if err != nil {
		return err
	}
	file.Close()
	return nil
}

// createMetaFile creates the meta file in the specified directory
func (m *WorkspaceManager) createMetaFile(dodlDir string) error {
	metaPath := pathcalculator.CalculateMetaFile(dodlDir)
	file, err := os.Create(metaPath)
	if err != nil {
		return err
	}
	file.Close()
	return nil
}
