package pathfinder

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindWorkspaceRootDir(t *testing.T) {
	tests := []struct {
		name          string
		workspacePath string
		entryPath     string
		wantedPath    string
		wantErr       bool
		errType       error
	}{
		{
			name:          "basic directory",
			workspacePath: "test/project1/.dodl",
			entryPath:     "test/project1/src/",
			wantedPath:    "test/project1",
			wantErr:       false,
		},
		{
			name:          "directory with trailing slash",
			workspacePath: "test/project2/.dodl",
			entryPath:     "test/project2/src/",
			wantedPath:    "test/project2",
			wantErr:       false,
		},
		{
			name:          "long directory path - deep nest",
			workspacePath: "test/project3/.dodl",
			entryPath:     "test/project3/src/level1/level2/level3/level4/level5/",
			wantedPath:    "test/project3",
			wantErr:       false,
		},
		{
			name:          "directory with multiple trailing slashes",
			workspacePath: "test/project4/.dodl",
			entryPath:     "test/project4/src///",
			wantedPath:    "test/project4",
			wantErr:       false,
		},
		{
			name:          "directory not in workspace",
			workspacePath: "test/project5/.dodl",
			entryPath:     "test/project6/src/",
			wantedPath:    "",
			wantErr:       true,
			errType:       &NotInWorkspaceError{},
		},
		{
			name:          "directory in user directory",
			workspacePath: "test/project7/.dodl",
			entryPath:     "test/project7/src/",
			wantedPath:    "test/project7",
			wantErr:       true,
			errType:       &InUserDirectoryError{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// get temporary directory
			tempDir, err := os.MkdirTemp("", "dodl")
			if err != nil {
				t.Errorf("error creating temporary directory: %v", err)
			}
			defer os.RemoveAll(tempDir)

			// setup "home" if necessary
			if tt.errType != nil {
				if _, ok := tt.errType.(*InUserDirectoryError); ok {
					userHome := filepath.Join(tempDir, tt.wantedPath)
					t.Setenv("HOME", userHome)
				}
			}

			// create dodl directory using os
			wrkPath := filepath.Join(tempDir, tt.workspacePath)
			err = os.MkdirAll(wrkPath, 0755)
			if err != nil {
				t.Errorf("error creating directory: %v", err)
			}
			// create entry directory using os
			entPath := filepath.Join(tempDir, tt.entryPath)
			err = os.MkdirAll(entPath, 0755)
			if err != nil {
				t.Errorf("error creating directory: %v", err)
			}

			// finalise expected path
			tt.wantedPath = filepath.Join(tempDir, tt.wantedPath)

			got, err := FindWorkspaceRootDir(entPath)
			if tt.wantErr {
				assert.Error(t, err)
				assert.IsType(t, tt.errType, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tt.wantedPath, got)
		})
	}
}
