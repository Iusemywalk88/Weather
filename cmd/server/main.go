package main

import (
	"github.com/Iusemywalk88/Weather/db"
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
	db := db.Connect()

	r := gin.Default()
	weatherClient := client.NewWeatherClient(cfg.WeatherAPIURL, cfg.WeatherAPIKey)
	weatherHandler := handlers.NewWeatherHandler(weatherClient)
	authHandler := handlers.NewAuthHandler(db)

	r.GET("/weather/:city", weatherHandler.HandleWeather)
	r.POST("/register", authHandler.RegisterUser)
	r.POST("/login", authHandler.Login)

	log.Printf("Сервер запущен на http://localhost:%s", cfg.Port)
	r.Run(":" + cfg.Port)
}
