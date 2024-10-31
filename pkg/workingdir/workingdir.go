package wd

import "os"

// WorkingDirProvider is an interface for providing the current working directory.
type WorkingDirProvider interface {
	GetWorkingDir() (string, error)
}

// DefaultWorkingDirProvider is a default implementation of WorkingDirProvider that uses os.Getwd.
type DefaultWorkingDirProvider struct{}

// GetWorkingDir returns the current working directory using os.Getwd.
func (p *DefaultWorkingDirProvider) GetWorkingDir() (string, error) {
	return os.Getwd()
}
