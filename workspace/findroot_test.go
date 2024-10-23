package workspace

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindWorkspaceRootWithDirs(t *testing.T) {
	// Helper function to create a temporary directory structure
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
		name        string
		dirs        []string
		entryPoint  string
		expectRoot  string
		expectError bool
	}{
		{
			name: "finds workspace root in nested directories",
			dirs: []string{
				"workspace/.dodl/nested/child",
			},
			entryPoint:  "workspace/.dodl/nested/child",
			expectRoot:  "workspace",
			expectError: false,
		},
		{
			name: "returns error when not in a workspace",
			dirs: []string{
				"dir1/dir2",
			},
			entryPoint:  "dir1/dir2",
			expectRoot:  "",
			expectError: true,
		},
		{
			name: "detects workspace at the same level",
			dirs: []string{
				"workspace/.dodl",
			},
			entryPoint:  "workspace",
			expectRoot:  "workspace",
			expectError: false,
		},
		{
			name:        "handles empty entry point",
			dirs:        nil,
			entryPoint:  "",
			expectRoot:  "",
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup the test environment
			tempDir := createTempDirWithStructure(t, tt.dirs)
			entryPoint := tt.entryPoint
			if entryPoint != "" {
				entryPoint = filepath.Join(tempDir, tt.entryPoint)
			}

			// Run the function under test
			root, err := FindWorkspaceRoot(entryPoint)

			if tt.expectError {
				assert.Error(t, err)
				if err == nil {
					t.Logf("Expected error, but got workspace root: %q", root)
				}
				return
			}

			expectedRoot := filepath.Join(tempDir, tt.expectRoot)
			if root != expectedRoot {
				t.Errorf("Expected root: %s, got: %s", expectedRoot, root)
			}
		})
	}
}
