package main

import (
	"github.com/Iusemywalk88/Weather/internal/handlers"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	r.GET("/:city", handlers.HandleWeather)

	r.Run(":8080")
}
