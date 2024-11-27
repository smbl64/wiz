package config

type Config struct {
	PatternsDir string
	Model       string
}

func Current() *Config {
	return &Config{
		PatternsDir: "/Users/mohammad/.config/fabric/patterns",
		Model:       "qwen2.5:latest",
	}
}
