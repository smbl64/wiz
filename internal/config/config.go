package config

import (
	"errors"
	"fmt"
	"os"
	"path"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Model         string `yaml:"model"`
	OllamaAPIBase string `yaml:"ollamaApiBase"`
}

var (
	configDir string
	config    *Config
)

// Init creates the config folder if it doesn't already exist.
// Afterwards it will load the configuration file if one exists.
func Init() error {
	home, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("cannot detect home directory: %v", err)
	}

	// TODO: Windows
	configDir = path.Join(home, ".config", "wiz")
	err = os.MkdirAll(configDir, 0700)
	if err != nil {
		return fmt.Errorf("cannot create config directory: %v", err)
	}

	return loadConfig()
}

func ConfigDir() string {
	return configDir
}

func defaultConfig() *Config {
	return &Config{
		Model:         "qwen2.5:latest",
		OllamaAPIBase: "http://localhost:11434",
	}
}

func configFile() string {
	return path.Join(ConfigDir(), "config.yaml")
}

func Current() *Config {
	return config
}

func loadConfig() error {
	// Load the file
	bytes, err := os.ReadFile(configFile())
	if err != nil {
		// Does not exist? Return default config
		if errors.Is(err, os.ErrNotExist) {
			config = defaultConfig()
			return nil
		}

		return err
	}

	// Unmarshal it
	var c Config
	err = yaml.Unmarshal(bytes, &c)
	if err != nil {
		return fmt.Errorf("loadConfig: %v", err)
	}

	config = &c

	return nil
}

func saveConfig() error {
	c := Current()
	bytes, err := yaml.Marshal(c)
	if err != nil {
		return fmt.Errorf("saveConfig: %v", err)
	}

	return os.WriteFile(configFile(), bytes, 0600)
}
