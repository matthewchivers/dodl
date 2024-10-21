package workspace

import (
	"os"
	"path/filepath"
	"testing"
)

func TestFindWorkspaceRootWithDirs(t *testing.T) {
	// Helper function to create a temporary directory structure
	createTempDirWithStructure := func(t *testing.T, dirs []string) string {
		root, err := os.MkdirTemp("", "workspace_test")
		if err != nil {
			t.Fatalf("Failed to create temp directory: %v", err)
		}

		for _, dir := range dirs {
			fullPath := filepath.Join(root, dir)
			err := os.MkdirAll(fullPath, 0755)
			if err != nil {
				t.Fatalf("Failed to create directory structure: %v", err)
			}
		}

		return root
	}

	// Clean up temporary directories
	deferCleanup := func(t *testing.T, path string) {
		err := os.RemoveAll(path)
		if err != nil {
			t.Fatalf("Failed to clean up temp directory: %v", err)
		}
	}

	tests := []struct {
		name        string
		dirs        []string
		entryPoint  string
		expectRoot  string
		expectError error
	}{
		{
			name: "finds workspace root in nested directories",
			dirs: []string{
				"workspace/.dodl/nested/child",
			},
			entryPoint:  "workspace/.dodl/nested/child",
			expectRoot:  "workspace",
			expectError: nil,
		},
		{
			name: "returns error when not in a workspace",
			dirs: []string{
				"dir1/dir2",
			},
			entryPoint:  "dir1/dir2",
			expectRoot:  "",
			expectError: ErrNotInWorkspace,
		},
		{
			name: "detects workspace at the same level",
			dirs: []string{
				"workspace/.dodl",
			},
			entryPoint:  "workspace",
			expectRoot:  "workspace",
			expectError: nil,
		},
		{
			name:        "handles empty entry point",
			dirs:        nil,
			entryPoint:  "",
			expectRoot:  "",
			expectError: ErrNotInWorkspace,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup the test environment
			tempDir := createTempDirWithStructure(t, tt.dirs)
			defer deferCleanup(t, tempDir)
			entryPoint := filepath.Join(tempDir, tt.entryPoint)

			// Run the function under test
			root, err := FindWorkspaceRoot(entryPoint)

			if tt.expectError != nil {
				if err != tt.expectError {
					t.Fatalf("Expected error: %v, got: %v", tt.expectError, err)
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
