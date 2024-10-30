package initialise

import (
	"testing"

	"github.com/matthewchivers/dodl/cmd/wd"
	"github.com/stretchr/testify/assert"
)

// TestInitCommandWithMockWorkingDir verifies that the init command can be created
// and executed without error in a mock working directory environment.
func TestInitCommandWithMockWorkingDir(t *testing.T) {
	testDir := t.TempDir()
	fakeWdProv := &wd.FakeWorkingDirProvider{Dir: testDir}

	initCmd := NewInitCmd(fakeWdProv)

	initCmd.SetArgs([]string{})

	err := initCmd.Execute()
	assert.NoError(t, err, "Expected init command to execute without error")
}
