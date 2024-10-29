package initialise

import (
	"testing"

	"github.com/matthewchivers/dodl/cmd/wd"
	"github.com/stretchr/testify/assert"
)

// TestInitCommandWithMockWorkingDir verifies that the init command can be created
// and executed without error in a mock working directory environment.
func TestInitCommandWithMockWorkingDir(t *testing.T) {
	// Set up a temporary test directory and mock working directory provider
	testDir := t.TempDir()
	fakeWdProv := &wd.FakeWorkingDirProvider{Dir: testDir}

	// Initialize the init command with the mock working directory provider
	initCmd := NewInitCmd(fakeWdProv)

	// Set command arguments to defaults (no user-specified input)
	initCmd.SetArgs([]string{})

	// Execute the command and assert no error occurs
	err := initCmd.Execute()
	assert.NoError(t, err, "Expected init command to execute without error")
}
