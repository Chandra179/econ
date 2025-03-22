package api

import (
	"os"
	"sync"
)

// Config holds configuration values for the API
type Config struct {
	AlphaVantageAPIKey string
	DefaultVersion     string
}

var (
	config     *Config
	configOnce sync.Once
)

// GetConfig returns the singleton config instance
func GetConfig() *Config {
	configOnce.Do(func() {
		config = &Config{
			AlphaVantageAPIKey: getEnvWithDefault("ALPHAVANTAGE_API_KEY", "demo"),
			DefaultVersion:     getEnvWithDefault("API_DEFAULT_VERSION", "1.0"),
		}
	})
	return config
}

// getEnvWithDefault returns the value of an environment variable or a default value
func getEnvWithDefault(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
