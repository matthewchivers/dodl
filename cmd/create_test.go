package cmd

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/matthewchivers/dodl/cmd/wd"
)

// TestCreateCommandWithMockWorkingDir is the simplest test case for the create command
// Just check that the command can be created and executed without error
func TestCreateCommandWithMockWorkingDir(t *testing.T) {
	testDir := t.TempDir()
	fakeWdProv := &wd.FakeWorkingDirProvider{
		Dir: testDir,
	}

	// make a .dodl directory in the test directory
	err := os.Mkdir(filepath.Join(testDir, ".dodl"), 0755)
	assert.NoError(t, err)

	// multiline string literal for the config file
	data := `default_document_type: journal
document_types:
  journal:
    template_file: "workspace.md"
    file_name_pattern: "{{.Year}}-{{.Month}}-{{.Day}}.md"
    directory_pattern: "{{.Year}}/{{.Month}}"
`

	// write a config file to the .dodl directory
	err = os.WriteFile(filepath.Join(testDir, ".dodl", "config.yaml"), []byte(data), 0644)
	assert.NoError(t, err)

	createCmd := NewCreateCmd(fakeWdProv)

	args := []string{} // no user-specified input (use defaults)

	createCmd.SetArgs(args)

	// create a config file flag
	err = createCmd.Flags().Set("config", filepath.Join(testDir, ".dodl", "config.yaml"))
	assert.NoError(t, err)

	err = createCmd.Execute()

	assert.NoError(t, err)
}
