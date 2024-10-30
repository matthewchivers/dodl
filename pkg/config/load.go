package config

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

const (
	DefaultConfigFileName = "config.yaml"
)

// ConfigOptions provides options for loading configurations.
type ConfigOptions struct {
	// CustomConfigFilePath is the path to a custom configuration file (optional / use for overriding from CLI flags).
	CustomConfigFilePath string

	// WorkspaceDodlDir is the path to the workspace .dodl directory.
	WorkspaceDodlDir string

	// UserDir is the path to the user's home directory
	UserDir string
}

// ConfigFileMissingError is an error type for when a config file is missing.
type ConfigFileMissingError struct {
	FilePath string
}

// Error returns the error message for a ConfigFileMissingError.
func (e *ConfigFileMissingError) Error() string {
	return fmt.Sprintf("config file %s does not exist", e.FilePath)
}

// LoadConfiguration loads a configuration from a single file.
// Returns the loaded configuration or an error if the configuration could not be loaded.
func LoadConfiguration(filePath string) (*Config, error) {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return nil, &ConfigFileMissingError{FilePath: filePath}
	}

	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("error opening config file %s: %w", filePath, err)
	}
	defer file.Close()

	var newConfig Config
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&newConfig); err != nil {
		return nil, fmt.Errorf("error decoding config file %s: %w", filePath, err)
	}
	return &newConfig, nil
}

// LoadConfigurations loads configurations from the user and workspace directories (or a custom file).
// Returns the merged configuration or an error if the configurations could not be loaded.
func LoadConfigurations(options ConfigOptions) (*Config, error) {
	if options.CustomConfigFilePath != "" {
		customConfig, err := LoadConfiguration(options.CustomConfigFilePath)
		if err != nil {
			return nil, err
		}
		return customConfig, nil
	}

	configsToMerge := []*Config{}

	userConfigDir := options.UserDir
	if userConfigDir == "" {
		confDir, err := getUserConfigDir(userConfigDir)
		if err != nil {
			return nil, err
		}
		userConfigDir = confDir
	}
	userConfig, err := LoadConfiguration(filepath.Join(userConfigDir, DefaultConfigFileName))
	if err != nil {
		if _, ok := err.(*ConfigFileMissingError); !ok {
			return nil, err
		}
	}
	if userConfig != nil {
		configsToMerge = append(configsToMerge, userConfig)
	}

	workspaceConfigDir := options.WorkspaceDodlDir
	if workspaceConfigDir == "" {
		return nil, fmt.Errorf("workspace config directory is required")
	}
	workspaceConfig, err := LoadConfiguration(filepath.Join(workspaceConfigDir, DefaultConfigFileName))
	if err != nil {
		if _, ok := err.(*ConfigFileMissingError); !ok {
			return nil, err
		}
	}
	if workspaceConfig != nil {
		configsToMerge = append(configsToMerge, workspaceConfig)
	}

	finalConfig := &Config{}
	for _, config := range configsToMerge {
		deepMergeConfig(finalConfig, config)
	}

	return finalConfig, nil
}

// deepMergeConfig merges the source configuration into the destination configuration.
func deepMergeConfig(dst, src *Config) {
	if src.DefaultDocumentType != "" {
		dst.DefaultDocumentType = src.DefaultDocumentType
	}

	if dst.CustomFields == nil {
		dst.CustomFields = map[string]interface{}{}
	}
	for k, v := range src.CustomFields {
		dst.CustomFields[k] = v
	}

	if dst.DocumentTypes == nil {
		dst.DocumentTypes = map[string]DocumentType{}
	}
	for k, v := range src.DocumentTypes {
		if existing, ok := dst.DocumentTypes[k]; ok {
			if v.TemplateFile != "" {
				existing.TemplateFile = v.TemplateFile
			}
			if v.FileNamePattern != "" {
				existing.FileNamePattern = v.FileNamePattern
			}
			if v.DirectoryPattern != "" {
				existing.DirectoryPattern = v.DirectoryPattern
			}
			if existing.CustomFields == nil {
				existing.CustomFields = map[string]interface{}{}
			}
			for k, v := range v.CustomFields {
				existing.CustomFields[k] = v
			}
			dst.DocumentTypes[k] = existing
		} else {
			dst.DocumentTypes[k] = v
		}
	}
}

// getUserConfigDir returns the path to the user's configuration directory
// based on the XDG_CONFIG_HOME environment variable or the user's home directory (or an override).
func getUserConfigDir(userDirOverride string) (string, error) {
	if userDirOverride != "" {
		return filepath.Join(userDirOverride, ".config", "dodl"), nil
	}

	configHome := os.Getenv("XDG_CONFIG_HOME")
	if configHome == "" {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		configHome = filepath.Join(homeDir, ".config")
	}

	return filepath.Join(configHome, "dodl"), nil
}
