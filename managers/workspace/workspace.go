package workspace

import (
	"os"

	"github.com/matthewchivers/dodl/utils/pathcalculator"
	"github.com/matthewchivers/dodl/utils/pathfinder"
)

var (
	wsm *Manager
)

type (
	AlreadyInWorkspaceError struct{}
	InUserDirectoryError    = pathfinder.InUserDirectoryError
)

func (e *AlreadyInWorkspaceError) Error() string {
	return "already in a workspace"
}

// Manager (workspace) is a struct that contains the paths to the current workspace and provides methods to interact with the workspace
type Manager struct {
	// entryPath is the path to the current directory
	entryPath string

	// workspacePath is the path to the current workspace (parent of the .dodl directory)
	workspacePath string

	// dodlPath is the path to the .dodl directory
	dodlPath string

	// configPath is the path to the configuration file
	configPath string

	// metaPath is the path to the meta file
	metaPath string

	// isIncumbent is true if the current directory is in an existing workspace
	isIncumbent bool

	// isUserDirectory is true if the current directory is the user directory
	isUserDirectory bool
}

func GetManager(entryPath string) (*Manager, error) {
	if wsm != nil && entryPath == wsm.entryPath {
		return wsm, nil
	}
	wsm := &Manager{
		entryPath: entryPath,
	}
	rootDir, err := pathfinder.FindWorkspaceRootDir(entryPath)
	if err != nil {
		switch err.(type) {
		case *pathfinder.NotInWorkspaceError:
			wsm.isIncumbent = false
			wsm.workspacePath = entryPath
		case *pathfinder.InUserDirectoryError:
			wsm.isUserDirectory = true
			wsm.isIncumbent = false
			wsm.workspacePath = entryPath
		default:
			return nil, err
		}
	} else {
		wsm.isIncumbent = true
		wsm.workspacePath = rootDir
	}
	wsm.dodlPath = pathcalculator.CalculateDodlDir(wsm.workspacePath)
	wsm.configPath = pathcalculator.CalculateConfigFile(wsm.dodlPath)
	wsm.metaPath = pathcalculator.CalculateMetaFile(wsm.dodlPath)
	return wsm, nil
}

// InitialiseWorkspace creates a new dodl workspace at the current location
func (m *Manager) InitialiseWorkspace() error {
	// Check if the current directory is already in a workspace
	if m.isIncumbent {
		return &AlreadyInWorkspaceError{}
	}
	if m.isUserDirectory {
		return &InUserDirectoryError{}
	}

	dodlDir := pathcalculator.CalculateDodlDir(m.workspacePath)

	if err := os.Mkdir(dodlDir, 0755); err != nil {
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

// createConfigFile creates the configuration file in the specified directory
func (m *Manager) createConfigFile() error {
	file, err := os.Create(m.configPath)
	if err != nil {
		return err
	}
	file.Close()
	return nil
}

// createMetaFile creates the meta file in the specified directory
func (m *Manager) createMetaFile() error {
	file, err := os.Create(m.metaPath)
	if err != nil {
		return err
	}
	file.Close()
	return nil
}

// GetConfigPath returns the path to the configuration file
func (m *Manager) GetConfigPath() string {
	return m.configPath
}
