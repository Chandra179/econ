package config

import (
	"os"
	"sync"
)

// Config holds all configuration settings for the application
type Config struct {
	// API settings
	AlphaVantageAPIKey  string
	AlphaVantageBaseURL string
	DefaultAPIVersion   string
	Port                string
}

var (
	config     *Config
	configOnce sync.Once
)

// GetConfig returns the singleton config instance
func GetConfig() *Config {
	configOnce.Do(func() {
		config = &Config{
			AlphaVantageAPIKey:  getEnvWithDefault("ALPHAVANTAGE_API_KEY", "demo"),
			AlphaVantageBaseURL: "https://www.alphavantage.co/query",
			DefaultAPIVersion:   getEnvWithDefault("API_DEFAULT_VERSION", "1.0"),
			Port:                getEnvWithDefault("PORT", "8080"),
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
