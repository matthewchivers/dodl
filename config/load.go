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

type ConfigOptions struct {
	CustomConfigFilePath string
	WorkspaceDodlDir     string
	UserConfigDir        string
}

type ConfigFileMissingError struct {
	FilePath string
}

func (e *ConfigFileMissingError) Error() string {
	return fmt.Sprintf("config file %s does not exist", e.FilePath)
}

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

func LoadConfigurations(options ConfigOptions) (*Config, error) {
	if options.CustomConfigFilePath != "" {
		customConfig, err := LoadConfiguration(options.CustomConfigFilePath)
		if err != nil {
			return nil, err
		}
		return customConfig, nil
	}

	configsToMerge := []*Config{}

	userConfigDir := options.UserConfigDir
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

func deepMergeConfig(dst, src *Config) {
	if src.DefaultDocumentType != "" {
		dst.DefaultDocumentType = src.DefaultDocumentType
	}

	if dst.CustomValues == nil {
		dst.CustomValues = map[string]interface{}{}
	}
	for k, v := range src.CustomValues {
		dst.CustomValues[k] = v
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
			if existing.CustomValues == nil {
				existing.CustomValues = map[string]interface{}{}
			}
			for key, value := range v.CustomValues {
				existing.CustomValues[key] = value
			}
			dst.DocumentTypes[k] = existing
		} else {
			dst.DocumentTypes[k] = v
		}
	}
}

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
