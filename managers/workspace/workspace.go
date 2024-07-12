package workspace

import (
	"errors"
	"os"
	"path/filepath"
)

const (
	// DodlDir is the name of the directory for the workspace
	DodlDir = ".dodl"

	// ConfigFile is the name of the configuration file for the workspace
	ConfigFile = "config.yaml"

	// MetaFile is the name of the meta file for the workspace
	MetaFile = "dodl_meta.json"
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

// GetWorkspaceRootPath retrieves the path to the current workspace (parent of the .dodl directory)
func (m *WorkspaceManager) GetWorkspaceRootPath(path string) (string, error) {
	if m.workspaceDir != "" {
		return m.workspaceDir, nil
	}
	dir := path
	for {
		if _, err := os.Stat(filepath.Join(dir, DodlDir)); err != nil {
			parent := filepath.Dir(dir)
			if parent == dir {
				return "", errors.New("not within a dodl workspace")
			}
			dir = parent
		} else {
			m.workspaceDir = dir
			return dir, nil // Found the .dodl directory
		}
	}
}

// InitialiseWorkspace creates a new dodl workspace at the current location
func (m *WorkspaceManager) InitialiseWorkspace() error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}
	if _, err := m.GetWorkspaceRootPath(dir); err == nil {
		return errors.New("already in a dodl workspace")
	}

	m.workspaceDir = dir

	dodlDir := filepath.Join(m.workspaceDir, DodlDir)
	m.dodlDir = dodlDir

	if err := os.Mkdir(m.dodlDir, 0755); err != nil {
		return err
	}

	if err := m.createConfigFile(); err != nil {
		return err
	}

	if err := m.createMetaFile(); err != nil {
		return err
	}
	return nil
}

func (m *WorkspaceManager) createConfigFile() error {
	configPath := filepath.Join(m.dodlDir, "config.yaml")
	file, err := os.Create(configPath)
	if err != nil {
		return err
	}
	file.Close()
	return nil
}

func (m *WorkspaceManager) createMetaFile() error {
	metaPath := filepath.Join(m.dodlDir, "dodl.meta")
	file, err := os.Create(metaPath)
	if err != nil {
		return err
	}
	file.Close()
	return nil
}
