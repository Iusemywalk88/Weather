package handlers

import (
	"github.com/Iusemywalk88/Weather/internal/client"
	"github.com/gin-gonic/gin"
)

type weatherHandler struct {
	weatherCLient client.WeatherClient
}

func NewWeatherHandler(client client.WeatherClient) *weatherHandler {
	return &weatherHandler{
		weatherCLient: client,
	}
}

func (w *weatherHandler) HandleWeather(c *gin.Context) {
	city := c.Param("city")

	if city == "" {
		c.JSON(400, gin.H{"error": "город не указан"})
		return
	}

	weather, err := w.weatherCLient.GetWeather(city)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, weather)
}
