package config

type Config struct {
	PatternsDir string
}

func Get() *Config {
	return &Config{
		PatternsDir: "/Users/mohammad/.config/fabric/patterns",
	}
}
