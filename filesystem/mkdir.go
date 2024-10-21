package filesystem

import (
	"fmt"
	"os"
)

// MkDir creates a directory at the specified path if it does not already exist.
func MkDir(path string) error {
	info, err := os.Stat(path)
	if err == nil {
		if !info.IsDir() {
			return fmt.Errorf("tried to create directory %q but a file with the same name already exists", path)
		}
	} else if !os.IsNotExist(err) {
		return err
	}
	err = os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}
