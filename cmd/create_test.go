package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/matthewchivers/dodl/cmd/wd"
)

// TestCreateCommandWithMockWorkingDir is the simplest test case for the create command
// Just check that the command can be created and executed without error
func TestCreateCommandWithMockWorkingDir(t *testing.T) {
	fakeWdProv := &wd.FakeWorkingDirProvider{
		Dir: "/fake/test/dir",
	}

	createCmd := NewInitCmd(fakeWdProv)

	args := []string{} // no user-specified input (use defaults)

	createCmd.SetArgs(args)

	err := createCmd.Execute()

	assert.NoError(t, err)
}
