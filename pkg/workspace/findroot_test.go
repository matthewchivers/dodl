package workspace

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestNewWorkspace verifies that the NewWorkspace function correctly identifies the workspace root (or errors when not in a workspace).
func TestNewWorkspace(t *testing.T) {
	createTempDirWithStructure := func(t *testing.T, dirs []string) string {
		tempDir := t.TempDir()
		for _, dir := range dirs {
			fullPath := filepath.Join(tempDir, dir)
			err := os.MkdirAll(fullPath, 0755)
			if err != nil {
				t.Fatalf("Failed to create directory structure: %v", err)
			}
		}
		return tempDir
	}

	tests := []struct {
		name             string
		dirs             []string
		workingDirectory string
		expectRoot       string
		expectError      bool
	}{
		{
			name: "finds workspace root in nested directories",
			dirs: []string{
				"workspace/.dodl/nested/child",
			},
			workingDirectory: "workspace/.dodl/nested/child",
			expectRoot:       "workspace",
			expectError:      false,
		},
		{
			name: "returns error when not in a workspace",
			dirs: []string{
				"dir1/dir2",
			},
			workingDirectory: "dir1/dir2",
			expectRoot:       "",
			expectError:      true,
		},
		{
			name: "detects workspace at the same level",
			dirs: []string{
				"workspace/.dodl",
			},
			workingDirectory: "workspace",
			expectRoot:       "workspace",
			expectError:      false,
		},
		{
			name:             "handles empty working directory",
			dirs:             nil,
			workingDirectory: "",
			expectRoot:       "",
			expectError:      true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tempDir := createTempDirWithStructure(t, tt.dirs)
			workingDirectory := tt.workingDirectory
			if workingDirectory != "" {
				workingDirectory = filepath.Join(tempDir, tt.workingDirectory)
			}

			wsp, err := NewWorkspace(workingDirectory)

			if tt.expectError {
				assert.Error(t, err)
				assert.Nil(t, wsp, "expected nil workspace on error")
				return
			}

			assert.NoError(t, err, "unexpected error occurred")

			expectedRoot := filepath.Join(tempDir, tt.expectRoot)
			assert.Equal(t, expectedRoot, wsp.RootPath(), "root path mismatch")

			expectedDodlPath := filepath.Join(expectedRoot, ".dodl")
			assert.Equal(t, expectedDodlPath, wsp.DodlPath(), "dodl path mismatch")
		})
	}
}
