package handlers

import (
	"github.com/Iusemywalk88/Weather/internal/client"
	"github.com/Iusemywalk88/Weather/models/handlers"
	"github.com/gin-gonic/gin"
	"net/http"
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
		c.JSON(http.StatusBadRequest, handlers.BaseResponse{Error: "город не указан"})
		return
	}

	weather, err := w.weatherClient.GetWeather(city)
	if err != nil {
		c.JSON(http.StatusInternalServerError, handlers.BaseResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, handlers.BaseResponse{Data: weather})
}
