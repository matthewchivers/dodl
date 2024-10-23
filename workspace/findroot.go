package workspace

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

var ErrNotInWorkspace = errors.New("not in a workspace")

// checkForDodlDir checks if a .dodl directory exists at the given path.
func checkForDodlDir(path string) (bool, error) {
	dodlPath := filepath.Join(path, ".dodl")
	info, err := os.Stat(dodlPath)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil // .dodl does not exist at this path
		}
		return false, err // some other error occurred
	}

	return info.IsDir(), nil // .dodl exists
}

// FindWorkspaceRoot identifies the root of the workspace that contains the given entry point.
// It returns the root path if in a workspace, ErrNotInWorkspace if not in a workspace, or an error.
func FindWorkspaceRoot(entryPoint string) (string, error) {
	if entryPoint == "" {
		return "", fmt.Errorf("supplied entry point is empty")
	}

	currentPath, err := filepath.Abs(entryPoint)
	if err != nil {
		return "", err // Return the error if entryPoint cannot be resolved
	}

	for {
		exists, err := checkForDodlDir(currentPath)
		if err != nil {
			return "", err
		}
		if exists {
			return currentPath, nil
		}

		parentPath := filepath.Dir(currentPath)
		if parentPath == currentPath {
			// reached the root of the filesystem without finding .dodl
			return "", ErrNotInWorkspace
		}
		currentPath = parentPath
	}
}
