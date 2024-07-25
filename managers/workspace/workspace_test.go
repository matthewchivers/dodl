package workspace

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWorkspaceManager_InitialiseWorkspace(t *testing.T) {
	tests := []struct {
		name          string
		entryPath     string
		incumbDodlDir string
		userDodlDir   string
		wantErr       bool
		errType       error
	}{
		{
			name:          "basic directory",
			entryPath:     "test/project1/src/",
			incumbDodlDir: "",
			userDodlDir:   "test/",
			wantErr:       false,
		},
		{
			name:          "directory with trailing slash",
			entryPath:     "test/project2/src/",
			incumbDodlDir: "",
			userDodlDir:   "test/",
			wantErr:       false,
		},
		{
			name:          "long directory path - deep nest",
			entryPath:     "test/project3/src/level1/level2/level3/level4/level5/",
			incumbDodlDir: "",
			userDodlDir:   "test/",
			wantErr:       false,
		},
		{
			name:          "directory with existing .dodl directory",
			entryPath:     "test/project4/src/",
			incumbDodlDir: "test/project4/.dodl",
			userDodlDir:   "test/",
			wantErr:       true,
			errType:       &AlreadyInWorkspaceError{},
		},
		{
			name:          "directory in user directory",
			entryPath:     "test/",
			incumbDodlDir: "",
			userDodlDir:   "test",
			wantErr:       true,
			errType:       &InUserDirectoryError{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tempDir, err := os.MkdirTemp("", "dodl")
			if err != nil {
				t.Fatalf("os.MkdirTemp() error = %v", err)
			}
			defer os.RemoveAll(tempDir)

			tt.entryPath = tempDir + tt.entryPath
			if tt.incumbDodlDir != "" {
				tt.incumbDodlDir = tempDir + tt.incumbDodlDir
			}
			tt.userDodlDir = tempDir + tt.userDodlDir

			// create directories
			err = os.MkdirAll(tt.entryPath, 0755)
			if err != nil {
				t.Fatalf("os.MkdirAll() error = %v", err)
			}
			if tt.incumbDodlDir != "" {
				err = os.MkdirAll(tt.incumbDodlDir, 0755)
				if err != nil {
					t.Fatalf("os.MkdirAll() error = %v", err)
				}
			}
			err = os.MkdirAll(tt.userDodlDir, 0755)
			if err != nil {
				t.Fatalf("os.MkdirAll() error = %v", err)
			}
			t.Setenv("HOME", tt.userDodlDir)

			// initialise workspace
			wkm, err := GetManager(tt.entryPath)
			if err != nil {
				t.Fatalf("GetWorkspaceManager() error = %v", err)
			}
			err = wkm.InitialiseWorkspace()
			if tt.wantErr {
				assert.Error(t, err)
				assert.IsType(t, tt.errType, err)
				return
			}
			assert.NoError(t, err)

			// check if .dodl directory exists where expected
			dodlPath := filepath.Join(tt.entryPath, ".dodl")
			_, err = os.Stat(dodlPath)
			assert.NoError(t, err)

			// check if config file exists where expected
			configPath := filepath.Join(dodlPath, "config.yaml")
			_, err = os.Stat(configPath)
			assert.NoError(t, err)

			// check if meta file exists where expected
			metaPath := filepath.Join(dodlPath, "dodl_meta.json")
			_, err = os.Stat(metaPath)
			assert.NoError(t, err)
		})
	}
}
