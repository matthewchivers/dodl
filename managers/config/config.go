package config

import (
	"io"
	"os"

	"github.com/matthewchivers/dodl/managers/workspace"
	"github.com/matthewchivers/dodl/utils/pathcalculator"
	"gopkg.in/yaml.v2"
)

type Manager struct {
	workspace    *workspace.Manager
	mergedConfig *Config
}

func NewManager(entryPath string) (*Manager, error) {
	workspaceManager, err := workspace.GetManager(entryPath)
	if err != nil {
		return nil, err
	}
	return &Manager{
		workspace: workspaceManager,
	}, nil
}

// GetConfig loads the configuration.
func (c *Manager) GetConfig() (*Config, error) {
	if c.mergedConfig != nil {
		return c.mergedConfig, nil
	}
	workspaceConfig := &Config{}

	workspaceConfigPath := c.workspace.GetConfigPath()
	workspaceConfigData, err := c.getDataFromPath(workspaceConfigPath)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(workspaceConfigData, workspaceConfig)
	if err != nil {
		return nil, err
	}

	userConfig := &Config{}

	userConfigPath, err := pathcalculator.CalculateUserConfigPath()
	if err != nil {
		return nil, err
	}
	userConfigData, err := c.getDataFromPath(userConfigPath)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(userConfigData, userConfig)
	if err != nil {
		return nil, err
	}

	userConfig.Apply(workspaceConfig)
	c.mergedConfig = userConfig

	return c.mergedConfig, nil
}

// getDataFromPath loads the configuration from a given path.
func (c *Manager) getDataFromPath(configPath string) ([]byte, error) {
	configFile, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer configFile.Close()
	configData, err := io.ReadAll(configFile)
	if err != nil {
		return nil, err
	}
	return configData, nil
}
