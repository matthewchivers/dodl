package filesystem_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/matthewchivers/dodl/filesystem"
	"github.com/stretchr/testify/assert"
)

func TestMkDir(t *testing.T) {
	tests := []struct {
		name          string
		existingDirs  []string
		existingFiles []string
		newDirPath    string
		expectedErr   bool
	}{
		{
			name:          "create new directory",
			existingDirs:  []string{},
			existingFiles: []string{},
			newDirPath:    "newdir",
			expectedErr:   false,
		},
		{
			name:          "create directory that already exists",
			existingDirs:  []string{"existingdir"},
			existingFiles: []string{},
			newDirPath:    "existingdir",
			expectedErr:   false,
		},
		{
			name:          "fail to create directory where file exists",
			existingDirs:  []string{},
			existingFiles: []string{"existingfile"},
			newDirPath:    "existingfile",
			expectedErr:   true,
		},
		{
			name:          "create nested directories",
			existingDirs:  []string{},
			existingFiles: []string{},
			newDirPath:    filepath.Join("parent", "child"),
			expectedErr:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			baseDir, err := os.MkdirTemp("", "testmkdir")
			assert.NoError(t, err, "failed to create temp dir")

			defer os.RemoveAll(baseDir)
			for _, existingDir := range tt.existingDirs {
				path := filepath.Join(baseDir, existingDir)
				err := os.Mkdir(path, os.ModePerm)
				assert.NoError(t, err, "failed to create existing directory")
			}

			for _, existingFile := range tt.existingFiles {
				path := filepath.Join(baseDir, existingFile)
				err := os.WriteFile(path, []byte("test"), os.ModePerm)
				assert.NoError(t, err, "failed to create existing file")
			}

			targetDir := filepath.Join(baseDir, tt.newDirPath)

			err = filesystem.MkDir(targetDir)
			if tt.expectedErr {
				assert.Error(t, err, "expected an error but got none")
			} else {
				assert.NoError(t, err, "unexpected error occurred")
			}

			if !tt.expectedErr {
				info, statErr := os.Stat(targetDir)
				assert.NoError(t, statErr, "expected directory to be created, but got an error")
				assert.True(t, info.IsDir(), "expected a directory but found something else: %v", targetDir)
			}
		})
	}
}
