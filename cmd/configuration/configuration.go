// SPDX-FileCopyrightText: 2020 SAP SE or an SAP affiliate company and Gardener contributors
//
// SPDX-License-Identifier: Apache-2.0

package configuration

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

const (
	defaultConfigFileName = "config"
	// DocforgeHomeDir defines the docforge home location
	DocforgeHomeDir = ".docforge"
)

// Loader loads the configuration
type Loader interface {
	Load() (*Config, error)
}

// DefaultConfigurationLoader default implementation of Loader
type DefaultConfigurationLoader func() (*Config, error)

// Load returns docforge configuration
func (d *DefaultConfigurationLoader) Load() (*Config, error) {
	if configFilePath, found := os.LookupEnv("DOCFORGECONFIG"); found {
		if configFilePath == "" {
			return nil, fmt.Errorf("the provided environment variable DOCFORGECONFIG is set to empty string")
		}
		return load(configFilePath)
	}

	userHomerDir, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("failed to get user home directory: %v", err)
	}

	configFilePath := filepath.Join(userHomerDir, DocforgeHomeDir, defaultConfigFileName)
	return load(configFilePath)
}

func load(configFilePath string) (*Config, error) {
	if len(configFilePath) <= 0 {
		return &Config{}, nil
	}

	stat, err := os.Stat(configFilePath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return &Config{}, nil
		}
		return nil, fmt.Errorf("failed to get file info for configuration file path %s: %v", configFilePath, err)
	}
	if stat.IsDir() {
		panic(fmt.Errorf("the config file path %s is directory, instead of file", configFilePath))
	}
	configFile, err := os.ReadFile(configFilePath)
	if err != nil {
		return nil, err
	}

	config := &Config{}
	if err := yaml.Unmarshal(configFile, config); err != nil {
		return nil, err
	}
	return config, nil
}
