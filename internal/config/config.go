package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Port          string `envconfig:"PORT" default:"8080"`
	WeatherAPIKey string `envconfig:"WEATHER_API_KEY" default:""`
	WeatherAPIURL string `envconfig:"WEATHER_API_HOST" default:""`
}

func Load() *Config {
	var c Config
	err := envconfig.Process("", &c)
	if err != nil {
		panic(err)
	}

	return &c
}
