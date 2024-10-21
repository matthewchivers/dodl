package wd

import "os"

type WorkingDirProvider interface {
	GetWorkingDir() (string, error)
}

type DefaultWorkingDirProvider struct{}

func (p *DefaultWorkingDirProvider) GetWorkingDir() (string, error) {
	return os.Getwd()
}
