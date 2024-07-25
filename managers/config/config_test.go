package config

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/matthewchivers/dodl/managers/workspace"
	"github.com/stretchr/testify/assert"
)

// getTempWorkspaceDir creates a temporary workspace directory for testing.
func getTempWorkspaceDir(t *testing.T) string {
	// generate a randomised temporary workspace directory for testing
	dir, err := os.MkdirTemp("", "dodl-test-workspace*")
	if err != nil {
		t.Fatalf("failed to crate temo dir: %v", err)
	}
	return dir
}

// setupTempWorkspaceDir creates a temporary workspace directory for testing.
func getTempUserDir(t *testing.T) string {
	// generate a randomised temporary workspace directory for testing
	dir, err := os.MkdirTemp("", "dodl-test-user*")
	if err != nil {
		t.Fatalf("failed to crate temp dir: %v", err)
	}
	return dir
}

// injectConfig injects a configuration file into the user's home directory.
func injectConfig(t *testing.T, directory, dataFile string) {
	// create a temporary configuration directory only if it doesn't exist
	configDir := filepath.Join(directory, ".dodl")
	if _, err := os.Stat(configDir); os.IsNotExist(err) {
		err := os.Mkdir(configDir, 0755)
		if err != nil {
			t.Fatalf("failed to create config directory: %v", err)
		}
	}
	// create a temporary configuration file only if it doesn't exist
	configFile := filepath.Join(configDir, "config.yaml")
	// read content from dataFile and write it to the configuration file
	content, err := os.ReadFile(dataFile)
	if err != nil {
		t.Fatalf("failed to read data file: %v", err)
	}
	// write content into the configuration file (overwrite if it already exists)
	if err = os.WriteFile(configFile, []byte(content), 0600); err != nil {
		t.Fatalf("failed to write content to config file: %v", err)
	}
}

func TestConfigManager_GetConfig(t *testing.T) {
	homeDir := getTempUserDir(t)
	defer os.RemoveAll(homeDir)
	t.Setenv("HOME", homeDir)
	workspaceDir := getTempWorkspaceDir(t)

	tests := []struct {
		name        string
		testDataDir string
		want        *Config
		wantErr     bool
	}{
		{
			name:        "test full user config",
			testDataDir: "fulluser",
			want: &Config{
				DocTypes: []DocType{
					{
						ID:               "fooID",
						DirectoryPattern: "{{.Year}}/{{.Month}}/{{.Day}}",
						FileNamePattern:  "{{.DocID}}.md",
						Topics:           []string{"fooTopic"},
						Editor:           "code",
					},
				},
			},
		},
		{
			name:        "test full workspace config",
			testDataDir: "fullworkspace",
			want: &Config{
				DocTypes: []DocType{
					{
						ID:               "fooID",
						DirectoryPattern: "{{.Year}}/{{.Month}}/{{.Day}}",
						FileNamePattern:  "{{.DocID}}.md",
						Topics:           []string{"fooTopic"},
						Editor:           "code",
					},
				},
			},
		},
		{
			name:        "test overwrite user config",
			testDataDir: "overwrite",
			want: &Config{
				DocTypes: []DocType{
					{
						ID:               "fooID",
						DirectoryPattern: "{{.Year}}/{{.Month}}/{{.Day}}",
						FileNamePattern:  "{{.DocID}}.md",
						Topics:           []string{"fooTopic"},
						Editor:           "code",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userConfigFile := filepath.Join("testdata", tt.testDataDir, "user.yaml")
			injectConfig(t, homeDir, userConfigFile)
			workspaceConfigFile := filepath.Join("testdata", tt.testDataDir, "workspace.yaml")
			injectConfig(t, workspaceDir, workspaceConfigFile)

			wsp, err := workspace.GetManager(workspaceDir)
			if err != nil {
				t.Fatalf("failed to get workspace manager: %v", err)
			}

			c := &Manager{
				workspace: wsp,
			}
			got, err := c.GetConfig()
			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetConfig() got = %v, want %v", got, tt.want)
			}
		})
	}
}
