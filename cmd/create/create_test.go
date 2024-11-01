package create

import (
	"os"
	"path/filepath"
	"testing"

	wd "github.com/matthewchivers/dodl/pkg/workingdir"
	"github.com/stretchr/testify/assert"
)

// TestCreateCommandWithMockWorkingDir validates that the create command can be initialized and executed without errors.
func TestCreateCommandWithMockWorkingDir(t *testing.T) {
	testDir := t.TempDir()
	fakeWdProv := &wd.FakeWorkingDirProvider{Dir: testDir}

	setupDodlDirectory(t, testDir)
	setupConfigFile(t, testDir)
	setupTemplateFile(t, testDir, "journal.md", "This is a journal template")

	createCmd := NewCreateCmd(fakeWdProv)
	createCmd.SetArgs([]string{}) // No user-specified document type, rely on defaults

	configFilePath := filepath.Join(testDir, ".dodl", "config.yaml")
	assert.NoError(t, createCmd.Flags().Set("config", configFilePath))

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
    file_name_pattern: "{{.Today}}-journal.md"
    directory_pattern: ["{{.Today}}-dir"]
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
