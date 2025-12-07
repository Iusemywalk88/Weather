package main

import (
	"github.com/Iusemywalk88/Weather/db"
	"github.com/Iusemywalk88/Weather/internal/client"
	"github.com/Iusemywalk88/Weather/internal/config"
	"github.com/Iusemywalk88/Weather/internal/handlers"
	"github.com/Iusemywalk88/Weather/internal/middleware"
	"github.com/Iusemywalk88/Weather/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Error loading config")
	}
	database := db.New(*cfg)

	r := gin.Default()
	weatherClient := client.NewWeatherClient(cfg.WeatherAPIURL, cfg.WeatherAPIKey)
	weatherHandler := handlers.NewWeatherHandler(weatherClient)
	authService := services.NewAuthService(database, []byte(cfg.JWTKey))
	authHandler := handlers.NewAuthHandler(authService, []byte(cfg.JWTKey))
	favoriteHandler := handlers.NewFavouritesHandler(database, weatherClient)
	authMiddleware := middleware.NewMiddleware([]byte(cfg.JWTKey))

	authorized := r.Group("/")
	authorized.Use(authMiddleware.AuthMiddleware())
	{
		authorized.POST("/favourites", favoriteHandler.AddFavourites)
		authorized.GET("/favourites", favoriteHandler.GetFavourites)
		authorized.DELETE("/favourites", favoriteHandler.DeleteFavourites)
	}

	r.GET("/weather/:city", weatherHandler.HandleWeather)
	r.POST("/register", authHandler.RegisterUser)
	r.POST("/login", authHandler.LoginUser)

	log.Printf("Сервер запущен на http://localhost:%s", cfg.Port)
	r.Run(":" + cfg.Port)
}
