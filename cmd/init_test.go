package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/matthewchivers/dodl/cmd/wd"
)

// TestinitCommandWithMockWorkingDir is the simplest test case for the init command
// Just check that the command can be created and executed without error
func TestInitCommandWithMockWorkingDir(t *testing.T) {
	testDir := t.TempDir()
	fakeWdProv := &wd.FakeWorkingDirProvider{
		Dir: testDir,
	}

	initCmd := NewInitCmd(fakeWdProv)

	args := []string{} // no user-specified input (use defaults)

	initCmd.SetArgs(args)

	err := initCmd.Execute()

	assert.NoError(t, err)
}
