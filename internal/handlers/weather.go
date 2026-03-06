package handlers

import (
	"github.com/Iusemywalk88/Weather/internal/services"
	"github.com/Iusemywalk88/Weather/models/responses"
	"github.com/gin-gonic/gin"
	"net/http"
)

type weatherHandler struct {
	weatherService services.WeatherService
}

func NewWeatherHandler(service services.WeatherService) *weatherHandler {
	return &weatherHandler{
		weatherService: service,
	}
}

func (w *weatherHandler) HandleWeather(c *gin.Context) {
	city := c.Param("city")

	if city == "" {
		c.JSON(http.StatusBadRequest, responses.BaseResponse{Error: "Nothing was written"})
		return
	}

	weather, err := w.weatherService.GetWeatherAndSaveHistory(city)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.BaseResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, responses.BaseResponse{Data: weather})
}
