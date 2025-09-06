package config

import (
	"os"
)

type Config struct {
	Port   string
	Debug  bool
	APIKey string
}

func Load() *Config {
	return &Config{
		Port:   getEnv("PORT", "8080"),
		Debug:  getEnv("DEBUG", "false") == "true",
		APIKey: getEnv("API_KEY", ""),
	}
}

func (c *Config) GetPort() string {
	if c.Port == "" {
		return "8080"
	}
	return c.Port
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
