package config

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfigurationsYamlConfig(t *testing.T) {
	tempDir := t.TempDir()
	yamlContent := `
default_document_type: markdown
custom_values:
  author: "Test User"
document_types:
  markdown:
    template_file: "default.md"
    file_name_pattern: "{{.Title}}.md"
    directory_pattern: "{{.Year}}/{{.Month}}"
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
	assert.Equal(t, "Test User", config.CustomValues["author"])
	assert.Equal(t, "default.md", config.DocumentTypes["markdown"].TemplateFile)
	assert.Equal(t, "{{.Title}}.md", config.DocumentTypes["markdown"].FileNamePattern)
	assert.Equal(t, "{{.Year}}/{{.Month}}", config.DocumentTypes["markdown"].DirectoryPattern)
}

func TestLoadConfigurationsJsonConfig(t *testing.T) {
	tempDir := t.TempDir()
	jsonContent := `{
        "default_document_type": "text",
        "custom_values": {
            "author": "Another User"
        },
        "document_types": {
            "text": {
                "template_file": "default.txt",
                "file_name_pattern": "{{.Title}}.txt",
                "directory_pattern": "{{.Year}}/{{.Month}}"
            }
        }
    }`

	filePath := filepath.Join(tempDir, "config.json")
	err := os.WriteFile(filePath, []byte(jsonContent), 0644)
	assert.NoError(t, err)

	options := ConfigOptions{
		CustomConfigFilePath: filePath,
	}

	config, err := LoadConfigurations(options)
	assert.NoError(t, err)
	assert.Equal(t, "text", config.DefaultDocumentType)
	assert.Equal(t, "Another User", config.CustomValues["author"])
	assert.Equal(t, "default.txt", config.DocumentTypes["text"].TemplateFile)
	assert.Equal(t, "{{.Title}}.txt", config.DocumentTypes["text"].FileNamePattern)
	assert.Equal(t, "{{.Year}}/{{.Month}}", config.DocumentTypes["text"].DirectoryPattern)
}

func TestLoadConfigurationsMultipleLayers(t *testing.T) {
	userTempDir := t.TempDir()
	workspaceTempDir := t.TempDir()

	userConfigContent := `
default_document_type: journal
custom_values:
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
		UserConfigDir:    userTempDir,
		WorkspaceDodlDir: workspaceTempDir,
	}

	config, err := LoadConfigurations(options)
	assert.NoError(t, err)
	assert.Equal(t, "journal", config.DefaultDocumentType)
	assert.Equal(t, "John Doe", config.CustomValues["author"])
	assert.Equal(t, "workspace.md", config.DocumentTypes["journal"].TemplateFile)
}

func TestLoadConfigurationsMergeDocuments(t *testing.T) {
	userTempDir := t.TempDir()
	workspaceTempDir := t.TempDir()

	userConfigContent := `
default_document_type: journal
custom_values:
  author: "John Doe"
document_types:
  journal:
    template_file: "user.md"
    directory_pattern: "{{.Year}}/{{.Month}}"
    file_name_pattern: "{{.Title}}.md"
    custom_values:
      mood: "sad"
`

	workspaceConfigContent := `
document_types:
  journal:
    template_file: "workspace.md"
    file_name_pattern: "workspace-{{.Title}}.md"
    directory_pattern: "{{.Year}}/workspaceoverride"
    custom_values:
      mood: "happy"
  planning:
    template_file: "planning.md"
    file_name_pattern: "plan-{{.Year}}-{{.Month}}-{{.Day}}.md"
`

	userConfigPath := filepath.Join(userTempDir, "config.yaml")
	err := os.WriteFile(userConfigPath, []byte(userConfigContent), 0644)
	assert.NoError(t, err)

	workspaceConfigPath := filepath.Join(workspaceTempDir, "config.yaml")
	err = os.WriteFile(workspaceConfigPath, []byte(workspaceConfigContent), 0644)
	assert.NoError(t, err)

	options := ConfigOptions{
		UserConfigDir:    userTempDir,
		WorkspaceDodlDir: workspaceTempDir,
	}

	config, err := LoadConfigurations(options)
	assert.NoError(t, err)
	assert.Equal(t, "journal", config.DefaultDocumentType)
	assert.Equal(t, "John Doe", config.CustomValues["author"])

	journalDocType := config.DocumentTypes["journal"]
	assert.NotNil(t, journalDocType)
	assert.Equal(t, "workspace.md", journalDocType.TemplateFile)
	assert.Equal(t, "{{.Year}}/workspaceoverride", journalDocType.DirectoryPattern)
	assert.Equal(t, "workspace-{{.Title}}.md", journalDocType.FileNamePattern)
	assert.Equal(t, "happy", journalDocType.CustomValues["mood"])

	planningDocType := config.DocumentTypes["planning"]
	assert.NotNil(t, planningDocType)
	assert.Equal(t, "planning.md", planningDocType.TemplateFile)
	assert.Equal(t, "plan-{{.Year}}-{{.Month}}-{{.Day}}.md", planningDocType.FileNamePattern)
}

func TestLoadConfigurationsMissingFile(t *testing.T) {
	options := ConfigOptions{
		CustomConfigFilePath: "non_existent.yaml",
	}

	_, err := LoadConfigurations(options)
	assert.Error(t, err)
}

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
