package config

import (
	"os"
)

type Config struct {
	Port          string
	WeatherAPIKey string
}

func Load() *Config {
	return &Config{
		Port:          getEnv("PORT", "8080"),
		WeatherAPIKey: getEnv("WEATHER_API_KEY", "9b4430c225c9caf186f7f5a81414f451"),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
