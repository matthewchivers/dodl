package core

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/matthewchivers/dodl/pkg/config"
	"github.com/matthewchivers/dodl/pkg/workspace"
)

// setupTestEnvironment creates a test directory structure and writes a template file to it
func setupTestEnvironment(t *testing.T, testDir string) (templateFilePath string, dodlDir string) {
	t.Helper()

	dodlDir = filepath.Join(testDir, ".dodl")
	templateDir := filepath.Join(dodlDir, "templates")
	require.NoError(t, os.MkdirAll(templateDir, os.ModePerm), "Failed to create template directory")

	templateFilePath = filepath.Join(templateDir, "test_template.md")
	templateContent := `Document Title: {{ .DocName }}
Date: {{ .Today | date "02-01-2006" }}
Topic: {{ .Topic }}
CustomField: {{ .CustomField }}
AnotherField: {{ .AnotherField }}`

	require.NoError(t, os.WriteFile(templateFilePath, []byte(templateContent), 0644), "Failed to write template file")

	return templateFilePath, dodlDir
}

// createMockAppContext creates a mock application context for testing
func createMockAppContext(testDir string, mockReferenceTime time.Time) *AppContext {
	return &AppContext{
		WorkingDir:    testDir,
		StartTime:     time.Now(),
		ReferenceTime: mockReferenceTime,
	}
}

// createMockDocumentType creates a mock document type for testing
func createMockDocumentType() config.DocumentType {
	return config.DocumentType{
		FileNamePattern: "{{ .Today | date \"2006-01-02\" }}-{{ .DocName }}.md",
		DirectoryPattern: []string{
			"docs",
			"{{ .Today | date \"2006\" }}",
			"{{ .Today | date \"January\" }}",
		},
		TemplateFile: "test_template.md",
		CustomFields: map[string]interface{}{"CustomField": "Example Custom Field"},
	}
}

// createCreateCommand creates a "create" command for testing
func createCreateCommand(appCtx *AppContext, docType config.DocumentType, wsp *workspace.Workspace) *CreateCommand {
	return &CreateCommand{
		DocName:      "TestDocument",
		DocType:      docType,
		CustomFields: map[string]interface{}{"AnotherField": "Additional Field"},
		Topic:        "Test Topic",
		AppCtx:       appCtx,
		Workspace:    wsp,
	}
}

// verifyGeneratedDocument checks that the expected document was generated
func verifyGeneratedDocument(t *testing.T, expectedDirPath, expectedFilePath, expectedContent string) {
	t.Helper()

	require.DirExists(t, expectedDirPath, "Expected directory does not exist")
	require.FileExists(t, expectedFilePath, "Expected file does not exist")

	content, err := os.ReadFile(expectedFilePath)
	require.NoError(t, err, "Failed to read the expected file")

	require.Equal(t, expectedContent, string(content), "File content does not match expected content")
}

// TestCreateCommand_Execute validates that the create command can be executed without errors
func TestCreateCommand_Execute(t *testing.T) {
	testDir := t.TempDir()
	setupTestEnvironment(t, testDir)

	mockDate := "2024-10-29"
	mockReferenceTime, err := time.Parse("2006-01-02", mockDate)
	require.NoError(t, err)

	appCtx := createMockAppContext(testDir, mockReferenceTime)
	docType := createMockDocumentType()
	wsp, err := workspace.NewWorkspace(testDir)
	require.NoError(t, err)
	cmd := createCreateCommand(appCtx, docType, wsp)

	require.NoError(t, cmd.Execute(), "Failed to execute create command")

	expectedDirPath := filepath.Join(testDir, "docs", mockReferenceTime.Format("2006"), mockReferenceTime.Format("January"))
	expectedFilePath := filepath.Join(expectedDirPath, fmt.Sprintf("%s-TestDocument.md", mockReferenceTime.Format("2006-01-02")))

	expectedContent := fmt.Sprintf(`Document Title: TestDocument
Date: %s
Topic: Test Topic
CustomField: Example Custom Field
AnotherField: Additional Field`, mockReferenceTime.Format("02-01-2006"))

	verifyGeneratedDocument(t, expectedDirPath, expectedFilePath, expectedContent)
}
