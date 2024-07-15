package pathfinder

import (
	"os"
	"path/filepath"

	"github.com/matthewchivers/dodl/utils/pathcalculator"
)

var (
	// workspaceDir is the path to the current workspace (parent of the .dodl directory)
	workspaceDir string
)

type NotInWorkspaceError struct{}

func (e *NotInWorkspaceError) Error() string {
	return "not in a workspace"
}

// FindWorkspaceRootDir performs a search based on the provided "path".
// If the path is in a workspace, it will return the path to the workspace root directory (parent of .dodl)
func FindWorkspaceRootDir(path string) (string, error) {
	if workspaceDir != "" {
		return workspaceDir, nil
	}

	userDodlDir, err := pathcalculator.CalculateUserDodlDir()
	if err != nil {
		return "", err
	}

	dir := path
	for {
		dirToCheck := pathcalculator.CalculateDodlDir(dir)

		if dirToCheck == userDodlDir {
			return "", &NotInWorkspaceError{}
		}

		if _, err := os.Stat(dirToCheck); err != nil {
			parent := filepath.Dir(dir)
			if parent == dir {
				return "", &NotInWorkspaceError{}
			}
			dir = parent
		} else {
			workspaceDir = dir
			return workspaceDir, nil // Found the .dodl directory
		}
	}
}
