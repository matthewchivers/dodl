package wd

// FakeWorkingDirProvider is a fake implementation of WorkingDirProvider that returns a fixed directory.
type FakeWorkingDirProvider struct {
	Dir string
}

// GetWorkingDir returns the fixed directory provided to the FakeWorkingDirProvider.
func (p *FakeWorkingDirProvider) GetWorkingDir() (string, error) {
	return p.Dir, nil
}
