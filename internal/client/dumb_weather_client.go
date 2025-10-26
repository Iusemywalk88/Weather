package client

import "github.com/Iusemywalk88/Weather/models"

type dumbWeatherClient struct {
}

func NewDumbWeatherClient() WeatherClient {
	return dumbWeatherClient{}
}

func (dumbWeatherClient) GetWeather(city string) (models.WeatherResponse, error) {
	return models.WeatherResponse{}, nil
}
