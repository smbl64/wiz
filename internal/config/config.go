package config

import (
	"fmt"
	"os"
	"path"
)

type Config struct {
	PatternsDir   string
	Model         string
	OllamaAPIBase string
}

var configDir string

// Init creates the config folder if it doesn't already exist.
func Init() error {
	home, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("cannot detect home directory: %v", err)
	}

	configDir = path.Join(home, ".config", "wiz")
	err = os.MkdirAll(configDir, 0700)
	if err != nil {
		return fmt.Errorf("cannot create config directory: %v", err)
	}

	return nil
}

func defaultConfig() *Config {

	return &Config{
		PatternsDir:   path.Join(configDir, "patterns"),
		Model:         "qwen2.5:latest",
		OllamaAPIBase: "http://localhost:11434",
	}
}

func Current() *Config {
	return defaultConfig()
}

// Load the default configuration.
func Load() error {
	return nil
}
