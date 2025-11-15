package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"log"
)

type Config struct {
	Port          string `envconfig:"PORT" default:"8080"`
	WeatherAPIKey string `envconfig:"WEATHER_API_KEY"`
	WeatherAPIURL string `envconfig:"WEATHER_API_HOST"`
	JWTKey        string `envconfig:"JWT_KEY"`
	DBHost        string `envconfig:"DB_HOST"`
	DBPort        string `envconfig:"DB_PORT"`
	DBUser        string `envconfig:"DB_USER"`
	DBPass        string `envconfig:"DB_PASSWORD"`
	DBName        string `envconfig:"DB_NAME"`
}

func Load() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
		return nil, err
	}

	var c Config
	err := envconfig.Process("", &c)
	if err != nil {
		panic(err)
	}

	return &c, nil
}
