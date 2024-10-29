package create

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/matthewchivers/dodl/cmd/wd"
	"github.com/stretchr/testify/assert"
)

// TestCreateCommandWithMockWorkingDir validates that the create command can be initialized and executed without errors.
func TestCreateCommandWithMockWorkingDir(t *testing.T) {
	testDir := t.TempDir()
	fakeWdProv := &wd.FakeWorkingDirProvider{Dir: testDir}

	// Setup the .dodl directory structure and configuration file
	setupDodlDirectory(t, testDir)
	setupConfigFile(t, testDir)
	setupTemplateFile(t, testDir, "journal.md", "This is a journal template")

	// Initialize create command with mock working directory provider
	createCmd := NewCreateCmd(fakeWdProv)
	createCmd.SetArgs([]string{}) // No user-specified document type, rely on defaults

	// Configure the command to use the mock config file
	configFilePath := filepath.Join(testDir, ".dodl", "config.yaml")
	assert.NoError(t, createCmd.Flags().Set("config", configFilePath))

	// Execute the command and assert no error occurred
	assert.NoError(t, createCmd.Execute())
}

// setupDodlDirectory creates the .dodl directory in the test environment.
func setupDodlDirectory(t *testing.T, rootDir string) {
	t.Helper()
	err := os.Mkdir(filepath.Join(rootDir, ".dodl"), 0755)
	assert.NoError(t, err)
}

// setupConfigFile creates a mock config.yaml file in the .dodl directory.
func setupConfigFile(t *testing.T, rootDir string) {
	t.Helper()

	configContent := `default_document_type: journal
document_types:
  journal:
    template_file: "journal.md"
    file_name_pattern: "{{.Year}}-{{.Month}}-{{.Day}}.md"
    directory_pattern: "{{.Year}}/{{.Month}}"
`

	configFilePath := filepath.Join(rootDir, ".dodl", "config.yaml")
	err := os.WriteFile(configFilePath, []byte(configContent), 0644)
	assert.NoError(t, err)
}

// setupTemplateFile creates a specified template file in the .dodl/templates directory.
func setupTemplateFile(t *testing.T, rootDir, filename, content string) {
	t.Helper()

	templatesDir := filepath.Join(rootDir, ".dodl", "templates")
	err := os.MkdirAll(templatesDir, 0755)
	assert.NoError(t, err)

	templateFilePath := filepath.Join(templatesDir, filename)
	err = os.WriteFile(templateFilePath, []byte(content), 0644)
	assert.NoError(t, err)
}
