package config

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestLoadConfigurationsEmptyConfig verifies that the default configuration is loaded
func TestLoadConfigurationsYamlConfig(t *testing.T) {
	tempDir := t.TempDir()
	yamlContent := `
default_document_type: markdown
custom_fields:
  author: "Test User"
document_types:
  markdown:
    template_file: "default.md"
    file_name_pattern: "{{ .DocType }}.md"
    directory_pattern:
        - "{{ .Today | date \"2006\"}}"
        - "{{ .Today | date \"Jan\" }}"
`

	filePath := filepath.Join(tempDir, "config.yaml")
	err := os.WriteFile(filePath, []byte(yamlContent), 0644)
	assert.NoError(t, err)

	options := ConfigOptions{
		CustomConfigFilePath: filePath,
	}

	config, err := LoadConfigurations(options)
	assert.NoError(t, err)
	assert.Equal(t, "markdown", config.DefaultDocumentType)
	assert.Equal(t, "Test User", config.CustomFields["author"])
	assert.Equal(t, "default.md", config.DocumentTypes["markdown"].TemplateFile)
	assert.Equal(t, "{{ .DocType }}.md", config.DocumentTypes["markdown"].FileNamePattern)
	assert.Equal(t, []string{"{{ .Today | date \"2006\"}}", "{{ .Today | date \"Jan\" }}"}, config.DocumentTypes["markdown"].DirectoryPattern)
}

// TestLoadConfigurationsMultipleLayers verifies that configurations are loaded from multiple layers (and merged correctly)
func TestLoadConfigurationsMultipleLayers(t *testing.T) {
	userTempDir := t.TempDir()
	workspaceTempDir := t.TempDir()

	userConfigContent := `
default_document_type: journal
custom_fields:
  author: "John Doe"
`

	workspaceConfigContent := `
document_types:
  journal:
    template_file: "workspace.md"
`

	userConfigPath := filepath.Join(userTempDir, "config.yaml")
	err := os.WriteFile(userConfigPath, []byte(userConfigContent), 0644)
	assert.NoError(t, err)

	workspaceConfigPath := filepath.Join(workspaceTempDir, "config.yaml")
	err = os.WriteFile(workspaceConfigPath, []byte(workspaceConfigContent), 0644)
	assert.NoError(t, err)

	options := ConfigOptions{
		UserDir:          userTempDir,
		WorkspaceDodlDir: workspaceTempDir,
	}

	config, err := LoadConfigurations(options)
	assert.NoError(t, err)
	assert.Equal(t, "journal", config.DefaultDocumentType)
	assert.Equal(t, "John Doe", config.CustomFields["author"])
	assert.Equal(t, "workspace.md", config.DocumentTypes["journal"].TemplateFile)
}

// TestLoadConfigurationsMergeDocuments verifies that document types are merged correctly (more in depth than TestLoadConfigurationsMultipleLayers)
func TestLoadConfigurationsMergeDocuments(t *testing.T) {
	userTempDir := t.TempDir()
	workspaceTempDir := t.TempDir()

	userConfigContent := `
default_document_type: journal
custom_fields:
  author: "John Doe"
document_types:
  journal:
    template_file: "user.md"
    directory_pattern: [
        "{{ .Today | date \"2006\" }}",
        "{{ .Today | date \"Jan\" }}"]
    file_name_pattern: ".{{ .DocType }}.md"
    custom_fields:
      mood: "sad"
`

	workspaceConfigContent := `
document_types:
  journal:
    template_file: "workspace.md"
    file_name_pattern: "workspace-{{ .DocType }}.md"
    directory_pattern: [ "{{ .Today | date \"2006\" }}", "workspaceoverride" ]
    custom_fields:
      mood: "happy"
  planning:
    template_file: "planning.md"
    file_name_pattern: "plan-{{ .Today | date \"2006\" }}-{{ .Today | date \"Jan\" }}-{{ .Today | date \"02\" }}.md"
`

	userConfigPath := filepath.Join(userTempDir, "config.yaml")
	err := os.WriteFile(userConfigPath, []byte(userConfigContent), 0644)
	assert.NoError(t, err)

	workspaceConfigPath := filepath.Join(workspaceTempDir, "config.yaml")
	err = os.WriteFile(workspaceConfigPath, []byte(workspaceConfigContent), 0644)
	assert.NoError(t, err)

	options := ConfigOptions{
		UserDir:          userTempDir,
		WorkspaceDodlDir: workspaceTempDir,
	}

	config, err := LoadConfigurations(options)
	assert.NoError(t, err)
	assert.Equal(t, "journal", config.DefaultDocumentType)
	assert.Equal(t, "John Doe", config.CustomFields["author"])

	journalDocType := config.DocumentTypes["journal"]
	assert.NotNil(t, journalDocType)
	assert.Equal(t, "workspace.md", journalDocType.TemplateFile)
	assert.Equal(t, []string{"{{ .Today | date \"2006\" }}", "workspaceoverride"}, journalDocType.DirectoryPattern)
	assert.Equal(t, "workspace-{{ .DocType }}.md", journalDocType.FileNamePattern)
	assert.Equal(t, "happy", journalDocType.CustomFields["mood"])

	planningDocType := config.DocumentTypes["planning"]
	assert.NotNil(t, planningDocType)
	assert.Equal(t, "planning.md", planningDocType.TemplateFile)
	assert.Equal(t, "plan-{{ .Today | date \"2006\" }}-{{ .Today | date \"Jan\" }}-{{ .Today | date \"02\" }}.md", planningDocType.FileNamePattern)
}

// TestLoadConfigurationsMissingFile verifies that an error is returned when the configuration file is missing
func TestLoadConfigurationsMissingFile(t *testing.T) {
	options := ConfigOptions{
		CustomConfigFilePath: "non_existent.yaml",
	}

	_, err := LoadConfigurations(options)
	assert.Error(t, err)
}

// TestLoadConfigurationsWithOverride verifies that the custom configuration file overrides the default configuration
func TestLoadConfigurationsWithOverride(t *testing.T) {
	tempDir := t.TempDir()
	yamlContent := `
default_document_type: text
`

	filePath := filepath.Join(tempDir, "config.yaml")
	err := os.WriteFile(filePath, []byte(yamlContent), 0644)
	assert.NoError(t, err)

	options := ConfigOptions{
		CustomConfigFilePath: filePath,
	}

	config, err := LoadConfigurations(options)
	assert.NoError(t, err)
	assert.Equal(t, "text", config.DefaultDocumentType)
}
