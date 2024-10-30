package workspace

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

var ErrNotInWorkspace = errors.New("not in a workspace")

// Workspace represents a workspace directory structure and provides methods to interact with it.
type Workspace struct {
	rootPath string
	dodlPath string
}

// NewWorkspace attempts to find the workspace root from a given directory.
// If found, it returns a Workspace instance; otherwise, it returns an error.
func NewWorkspace(workingDirectory string) (*Workspace, error) {
	newWsp := &Workspace{}

	err := newWsp.populatePaths(workingDirectory)
	if err != nil {
		return nil, err
	}

	return newWsp, nil
}

// RootPath returns the root directory of the workspace.
func (w *Workspace) RootPath() string {
	return w.rootPath
}

// DodlPath returns the path to the .dodl directory in the workspace.
func (w *Workspace) DodlPath() string {
	return w.dodlPath
}

// findWorkspaceRoot searches upwards from the working directory to find the workspace root.
// It returns an ErrNotInWorkspace error if the workspace root is not found.
func (w *Workspace) populatePaths(workingDirectory string) error {
	if workingDirectory == "" {
		return fmt.Errorf("supplied working directory is empty")
	}

	workingDirAbs, err := filepath.Abs(workingDirectory)
	if err != nil {
		return err
	}

	for {
		dodlDir := filepath.Join(workingDirAbs, ".dodl")
		exists, err := directoryExists(dodlDir)
		if err != nil {
			return err
		}
		if exists {
			w.rootPath = workingDirAbs
			w.dodlPath = dodlDir
			return nil
		}

		parentPath := filepath.Dir(workingDirAbs)
		if parentPath == workingDirAbs {
			return ErrNotInWorkspace
		}
		workingDirAbs = parentPath
	}
}

// directoryExists checks if a path exists and is a directory.
// It returns true if the path exists and is a directory; otherwise, it returns false and/or an error.
func directoryExists(path string) (bool, error) {
	info, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	return info.IsDir(), nil
}
