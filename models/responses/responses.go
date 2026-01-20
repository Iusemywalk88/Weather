package responses

import "github.com/Iusemywalk88/Weather/models"

type BaseResponse struct {
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type FavouriteWeatherResponse struct {
	City    string                  `json:"city"`
	Weather interface{}             `json:"weather"`
	History []models.WeatherHistory `json:"history,omitempty"`
}
