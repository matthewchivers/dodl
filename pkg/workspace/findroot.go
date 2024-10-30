package workspace

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

var ErrNotInWorkspace = errors.New("not in a workspace")

var (
	workspaceRoot map[string]string
)

// isDodlDirPresent checks if the supplied path contains a .dodl directory.
func isDodlDirPresent(path string) (bool, error) {
	info, err := os.Stat(filepath.Join(path, ".dodl"))
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil // .dodl does not exist at this path
		}
		return false, err // some other error occurred
	}
	return info.IsDir(), nil // .dodl exists
}

// FindWorkspaceRoot searches upwards from the working directory to find the workspace root.
// It returns the root path if found, or ErrNotInWorkspace if not.
func FindWorkspaceRoot(workingDirectory string) (string, error) {
	if workingDirectory == "" {
		return "", fmt.Errorf("supplied working directory is empty")
	}

	if workspaceRoot != nil {
		if root, ok := workspaceRoot[workingDirectory]; ok {
			return root, nil
		}
	}

	currentPath, err := filepath.Abs(workingDirectory)
	if err != nil {
		return "", err // Return the error if workingDirectory cannot be resolved
	}

	for {
		exists, err := isDodlDirPresent(currentPath)
		if err != nil {
			return "", err
		}
		if exists {
			if workspaceRoot == nil {
				workspaceRoot = make(map[string]string)
			}
			workspaceRoot[workingDirectory] = currentPath
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

// GetDodlDirectoryPath returns the path to the .dodl directory within the workspace root
func GetDodlDirectoryPath(workingDirectory string) (string, error) {
	workspaceRoot, err := FindWorkspaceRoot(workingDirectory)
	if err != nil {
		return "", err
	}
	// check if the .dodl directory exists
	if exists, err := isDodlDirPresent(workspaceRoot); err != nil {
		return "", err
	} else if !exists {
		return "", fmt.Errorf("workspace root does not contain a .dodl directory")
	}
	return filepath.Join(workspaceRoot, ".dodl"), nil
}
