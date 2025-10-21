package main

import (
	"github.com/Iusemywalk88/Weather/internal/config"
	"github.com/Iusemywalk88/Weather/internal/handlers"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {

	cfg := config.Load()

	r := gin.Default()

	r.GET("/weather/:city", handlers.HandleWeather)

	log.Printf("🚀 Сервер запущен на http://localhost:%s", cfg.Port)
	r.Run(":" + cfg.Port)
}
