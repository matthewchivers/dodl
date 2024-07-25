package pathfinder

import (
	"os"
	"path/filepath"

	"github.com/matthewchivers/dodl/utils/pathcalculator"
)

type NotInWorkspaceError struct{}

func (e *NotInWorkspaceError) Error() string {
	return "not in a workspace"
}

type InUserDirectoryError struct{}

func (e *InUserDirectoryError) Error() string {
	return "in the user directory"
}

// FindWorkspaceRootDir performs a search based on the provided "path".
// If the path is in a workspace, it will return the path to the workspace root directory (parent of .dodl)
func FindWorkspaceRootDir(path string) (string, error) {
	depth := 0
	userDodlDir, err := pathcalculator.CalculateUserDodlDir()
	if err != nil {
		return "", err
	}

	dir := path
	for {
		dirToCheck := pathcalculator.CalculateDodlDir(dir)

		if dirToCheck == userDodlDir && depth == 0 { // If the first directory checked is the user directory, return an error
			return "", &InUserDirectoryError{}
		} else if dirToCheck == userDodlDir && depth > 0 {
			return "", &NotInWorkspaceError{}
		}

		if _, err := os.Stat(dirToCheck); err != nil {
			parent := filepath.Dir(dir)
			depth++
			if parent == dir {
				return "", &NotInWorkspaceError{}
			}
			dir = parent
		} else {
			return dir, nil // Found the .dodl directory
		}
	}
}
