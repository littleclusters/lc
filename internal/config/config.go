package config

import (
	"fmt"
	"os"

	"github.com/goccy/go-yaml"
)

const configPath = "lc.yaml"

// Stages tracks the current and completed challenge stages.
type Stages struct {
	Current   string   `yaml:"current"`
	Completed []string `yaml:"completed"`
}

// Config represents the lc.yaml configuration file structure.
type Config struct {
	Challenge string `yaml:"challenge"`
	Stages    Stages `yaml:"stages"`
}

// Load reads and parses the lc.yaml configuration file.
func Load() (*Config, error) {
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return nil, fmt.Errorf("Not in a challenge directory\nRun this command from a directory created with 'lc init <challenge>'")
	}

	bytes, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("Failed to read config file: %w", err)
	}

	var cfg Config
	err = yaml.Unmarshal(bytes, &cfg)
	if err != nil {
		return nil, fmt.Errorf("Failed to parse config file: %w", err)
	}

	if cfg.Stages.Completed == nil {
		cfg.Stages.Completed = []string{}
	}

	return &cfg, nil
}

// Save writes the configuration to the default lc.yaml file.
func Save(cfg *Config) error {
	return SaveTo(cfg, configPath)
}

// SaveTo writes the configuration to the specified path.
func SaveTo(cfg *Config, path string) error {
	bytes, err := yaml.Marshal(cfg)
	if err != nil {
		return fmt.Errorf("Failed to serialize config: %w", err)
	}

	err = os.WriteFile(path, bytes, 0644)
	if err != nil {
		return fmt.Errorf("Failed to write config file: %w", err)
	}

	return nil
}
