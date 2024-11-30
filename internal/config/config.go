package config

type Config struct {
	PatternsDir   string
	Model         string
	OllamaAPIBase string
}

func Current() *Config {
	return &Config{
		PatternsDir:   "/Users/mohammad/.config/fabric/patterns",
		Model:         "qwen2.5:latest",
		OllamaAPIBase: "http://localhost:11434",
	}
}

// Load the default configuration.
func Load() error {
	return nil
}
