package services

import (
	"github.com/Iusemywalk88/Weather/db"
	"github.com/Iusemywalk88/Weather/internal/client"
	"github.com/Iusemywalk88/Weather/models"
	"time"
)

type WeatherService struct {
	weatherClient client.WeatherClient
	DB            *db.DB
}

func NewWeatherService(client client.WeatherClient, db *db.DB) WeatherService {
	return WeatherService{
		weatherClient: client,
		DB:            db,
	}
}

func (s *WeatherService) GetWeatherAndSaveHistory(city string) (models.WeatherResponse, error) {
	weather, err := s.weatherClient.GetWeather(city)
	if err != nil {
		return models.WeatherResponse{}, err
	}

	err = s.DB.CreateHistory(city, weather.Main.Temperature, weather.Weather[0].Description, time.Now())
	if err != nil {
		return models.WeatherResponse{}, err
	}

	return weather, nil
}
