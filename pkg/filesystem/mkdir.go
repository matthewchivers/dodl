package filesystem

import (
	"fmt"
	"os"
	"path/filepath"
)

// MkDir creates a directory at the specified path if it does not already exist.
// An error is returned if a file with the same name already exists (or if there is an error creating the directory).
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
// An error is returned if a file with the same name already exists (or if there is an error creating the directory).
func EnsureDirExists(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		// Create the directory if it doesn't exist
		return MkDir(path)
	}

	info, err := os.Stat(path)
	if err == nil && !info.IsDir() {
		return fmt.Errorf("tried to create directory %q but a file with the same name already exists", path)
	}
	return nil
}

// EnsureFileExists creates a file at the specified path if it does not already exist.
// An error is returned if there is an error creating the file.
func EnsureFileExists(path string, data []byte) error {
	dir := filepath.Dir(path)
	if err := EnsureDirExists(dir); err != nil {
		return err
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		return WriteFile(path, data)
	}

	return nil
}

// WriteFile writes the specified data to a file at the specified path.
// An error is returned if there is an error writing the file.
func WriteFile(path string, data []byte) error {
	err := os.WriteFile(path, data, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}
