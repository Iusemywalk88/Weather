package main

import (
	"github.com/Iusemywalk88/Weather/internal/client"
	"github.com/Iusemywalk88/Weather/internal/config"
	"github.com/Iusemywalk88/Weather/internal/handlers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg := config.Load()

	r := gin.Default()
	weatherClient := client.NewWeatherClient(cfg.WeatherAPIURL, cfg.WeatherAPIKey)
	handler := handlers.NewWeatherHandler(weatherClient)

	r.GET("/weather/:city", handler.HandleWeather)

	log.Printf("🚀 Сервер запущен на http://localhost:%s", cfg.Port)
	r.Run(":" + cfg.Port)
}
