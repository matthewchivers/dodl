package core

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/matthewchivers/dodl/config"
)

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

func createMockAppContext(testDir string, mockStartTime time.Time) *AppContext {
	return &AppContext{
		WorkingDir:    testDir,
		WorkspaceRoot: testDir,
		StartTime:     mockStartTime,
	}
}

func createMockDocumentType() config.DocumentType {
	return config.DocumentType{
		FileNamePattern:  "notes/{{ .Today | date \"2006-01-02\" }}-{{ .DocName }}.md",
		DirectoryPattern: "docs/{{ .Today | date \"2006\" }}/{{ .Today | date \"January\" }}",
		TemplateFile:     "test_template.md",
		CustomValues:     map[string]interface{}{"CustomField": "Example Custom Value"},
	}
}

func createCreateCommand(appCtx *AppContext, docType config.DocumentType) *CreateCommand {
	return &CreateCommand{
		DocName:      "TestDocument",
		DocType:      docType,
		CustomFields: map[string]interface{}{"AnotherField": "Additional Value"},
		Topic:        "Test Topic",
		AppCtx:       appCtx,
	}
}

func verifyGeneratedDocument(t *testing.T, expectedDirPath, expectedFilePath, expectedContent string) {
	t.Helper()

	require.DirExists(t, expectedDirPath, "Expected directory does not exist")
	require.FileExists(t, expectedFilePath, "Expected file does not exist")

	content, err := os.ReadFile(expectedFilePath)
	require.NoError(t, err, "Failed to read the expected file")

	require.Equal(t, expectedContent, string(content), "File content does not match expected content")
}

func TestCreateCommand_Execute(t *testing.T) {
	testDir := t.TempDir()
	setupTestEnvironment(t, testDir)

	mockDate := "2024-10-29"
	mockStartTime, err := time.Parse("2006-01-02", mockDate)
	require.NoError(t, err)

	appCtx := createMockAppContext(testDir, mockStartTime)
	docType := createMockDocumentType()
	cmd := createCreateCommand(appCtx, docType)

	require.NoError(t, cmd.Execute(), "Failed to execute create command")

	expectedDirPath := filepath.Join(testDir, "docs", mockStartTime.Format("2006"), mockStartTime.Format("January"))
	expectedFilePath := filepath.Join(expectedDirPath, "notes", fmt.Sprintf("%s-TestDocument.md", mockStartTime.Format("2006-01-02")))

	expectedContent := fmt.Sprintf(`Document Title: TestDocument
Date: %s
Topic: Test Topic
CustomField: Example Custom Value
AnotherField: Additional Value`, mockStartTime.Format("02-01-2006"))

	verifyGeneratedDocument(t, expectedDirPath, expectedFilePath, expectedContent)
}
