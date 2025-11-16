package handlers

import (
	"github.com/Iusemywalk88/Weather/internal/client"
	"github.com/Iusemywalk88/Weather/models/handlers"
	"github.com/gin-gonic/gin"
)

type weatherHandler struct {
	weatherClient client.WeatherClient
}

func NewWeatherHandler(client client.WeatherClient) *weatherHandler {
	return &weatherHandler{
		weatherClient: client,
	}
}

func (w *weatherHandler) HandleWeather(c *gin.Context) {
	city := c.Param("city")

	if city == "" {
		c.JSON(400, handlers.BaseResponse{Error: "город не указан"})
		return
	}

	weather, err := w.weatherClient.GetWeather(city)
	if err != nil {
		c.JSON(500, handlers.BaseResponse{Error: err.Error()})
		return
	}

	c.JSON(200, handlers.BaseResponse{Data: weather})
}
