package pathcalculator

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateDodlDir(t *testing.T) {
	tests := []struct {
		name         string
		workspaceDir string
		want         string
	}{
		{
			name:         "basic directory",
			workspaceDir: "test",
			want:         "test/.dodl",
		},
		{
			name:         "directory with trailing slash",
			workspaceDir: "test/",
			want:         "test/.dodl",
		},
		{
			name:         "directory with multiple trailing slashes",
			workspaceDir: "test///",
			want:         "test/.dodl",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CalculateDodlDir(tt.workspaceDir)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestCalculateConfigFile(t *testing.T) {
	tests := []struct {
		name         string
		workspaceDir string
		want         string
	}{
		{
			name:         "basic directory",
			workspaceDir: "test",
			want:         "test/config.yaml",
		},
		{
			name:         "directory with trailing slash",
			workspaceDir: "test/",
			want:         "test/config.yaml",
		},
		{
			name:         "directory with multiple trailing slashes",
			workspaceDir: "test///",
			want:         "test/config.yaml",
		},
		{
			name:         "directory with .dodl",
			workspaceDir: "test/.dodl",
			want:         "test/.dodl/config.yaml",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CalculateConfigFile(tt.workspaceDir)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestCalculateMetaFile(t *testing.T) {
	tests := []struct {
		name         string
		workspaceDir string
		want         string
	}{
		{
			name:         "basic directory",
			workspaceDir: "test",
			want:         "test/dodl_meta.json",
		},
		{
			name:         "directory with trailing slash",
			workspaceDir: "test/",
			want:         "test/dodl_meta.json",
		},
		{
			name:         "directory with multiple trailing slashes",
			workspaceDir: "test///",
			want:         "test/dodl_meta.json",
		},
		{
			name:         "directory with .dodl",
			workspaceDir: "test/.dodl",
			want:         "test/.dodl/dodl_meta.json",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CalculateMetaFile(tt.workspaceDir)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestCalculateUserDodlDir(t *testing.T) {
	tempHomeDir, err := os.MkdirTemp("", "dodlhome")
	assert.Nil(t, err)
	defer os.RemoveAll(tempHomeDir)
	t.Setenv("HOME", tempHomeDir)

	want := filepath.Join(tempHomeDir, dodlDirName)
	got, err := CalculateUserDodlDir()

	assert.NoError(t, err)
	assert.Equal(t, want, got)
}

func TestCalculateUserConfigPath(t *testing.T) {
	tempHomeDir, err := os.MkdirTemp("", "dodlhome")
	assert.Nil(t, err)
	defer os.RemoveAll(tempHomeDir)
	t.Setenv("HOME", tempHomeDir)

	want := filepath.Join(tempHomeDir, dodlDirName, configFile)
	got, err := CalculateUserConfigPath()

	assert.NoError(t, err)
	assert.Equal(t, want, got)
}
