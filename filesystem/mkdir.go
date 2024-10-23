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

// EnsureDirExists creates a directory at the specified path if it does not already exist.
func EnsureDirExists(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		// Create the directory if it doesn't exist
		return MkDir(path)
	}
	// check if path is an existing file (rather than directory)
	info, err := os.Stat(path)
	if err == nil && !info.IsDir() {
		return fmt.Errorf("tried to create directory %q but a file with the same name already exists", path)
	}
	return nil
}

// EnsureFileExists creates a file at the specified path if it does not already exist.
func EnsureFileExists(path string, data []byte) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		// Create the file if it doesn't exist
		return WriteFile(path, data)
	}
	return nil
}

// WriteFile writes the specified data to a file at the specified path.
func WriteFile(path string, data []byte) error {
	err := os.WriteFile(path, data, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}
