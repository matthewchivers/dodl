package workspace_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/matthewchivers/dodl/pkg/workspace"
	"github.com/stretchr/testify/assert"
)

func TestInitialise(t *testing.T) {
	tests := []struct {
		name          string
		existingDirs  []string
		existingFiles []string
		targetDir     string
		expectedErr   bool
	}{
		{
			name:          "initialise new workspace",
			existingDirs:  []string{},
			existingFiles: []string{},
			targetDir:     "workspace1",
			expectedErr:   false,
		},
		{
			name:          "re-initialise existing workspace",
			existingDirs:  []string{filepath.Join("workspace2", ".dodl"), filepath.Join("workspace2", ".dodl", "templates")},
			existingFiles: []string{},
			targetDir:     "workspace2",
			expectedErr:   false,
		},
		{
			name:          "fail to initialise where file exists",
			existingDirs:  []string{"workspace3"},
			existingFiles: []string{filepath.Join("workspace3", ".dodl")},
			targetDir:     "workspace3",
			expectedErr:   true,
		},
		{
			name:          "initialise nested workspace",
			existingDirs:  []string{"parent_workspace", ".dodl"},
			existingFiles: []string{},
			targetDir:     filepath.Join("parent_workspace", "nested_workspace"),
			expectedErr:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tempDir := t.TempDir()

			for _, existingDir := range tt.existingDirs {
				path := filepath.Join(tempDir, existingDir)
				err := os.MkdirAll(path, os.ModePerm)
				assert.NoError(t, err, "failed to create existing directory")
			}

			for _, existingFile := range tt.existingFiles {
				path := filepath.Join(tempDir, existingFile)
				err := os.WriteFile(path, []byte("test"), os.ModePerm)
				assert.NoError(t, err, "failed to create existing file")
			}

			targetDir := filepath.Join(tempDir, tt.targetDir)

			err := workspace.Initialise(targetDir)
			if tt.expectedErr {
				assert.Error(t, err, "expected an error but got none")
			} else {
				assert.NoError(t, err, "unexpected error occurred")
			}

			if !tt.expectedErr {
				workspaceRoot := filepath.Join(targetDir, ".dodl")
				info, statErr := os.Stat(workspaceRoot)
				assert.NoError(t, statErr, "expected workspace root to be created, but got an error")
				assert.True(t, info.IsDir(), "expected a directory but found something else: %v", workspaceRoot)

				templatesDir := filepath.Join(workspaceRoot, "templates")
				templatesInfo, templatesErr := os.Stat(templatesDir)
				assert.NoError(t, templatesErr, "expected templates directory to be created, but got an error")
				assert.True(t, templatesInfo.IsDir(), "expected a directory but found something else: %v", templatesDir)
			}
		})
	}
}
