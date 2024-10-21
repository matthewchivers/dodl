package wd

type FakeWorkingDirProvider struct {
	Dir string
}

func (p *FakeWorkingDirProvider) GetWorkingDir() (string, error) {
	return p.Dir, nil
}
