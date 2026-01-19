package config

import (
	"os"
)

type Config struct {
	DataPath string // Path to data directory
}

// Load loads configuration from environment variables
func Load() (*Config, error) {
	cfg := &Config{
		DataPath: getEnv("DATA_PATH", "data"),
	}

	return cfg, nil
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
