package services

import (
	"github.com/Iusemywalk88/Weather/internal/client"
	"github.com/Iusemywalk88/Weather/models"
	"time"
)

//go:generate mockgen -destination=mocks/mock_weather_client.go -package=mocks github.com/Iusemywalk88/Weather/internal/client WeatherClient
//go:generate mockgen -destination=mocks/mock_weather_repo.go -package=mocks github.com/Iusemywalk88/Weather/internal/services WeatherRepo
//go:generate mockgen -destination=mocks/mock_weather_service.go -package=mocks github.com/Iusemywalk88/Weather/internal/services WeatherServiceInterface

type WeatherServiceInterface interface {
	GetWeatherAndSaveHistory(city string) (models.WeatherResponse, error)
}

type WeatherRepo interface {
	CreateHistory(cityName string, temperature float64, description string, createdAt time.Time) error
}

type WeatherService struct {
	weatherClient client.WeatherClient
	repo          WeatherRepo
}

func NewWeatherService(client client.WeatherClient, repo WeatherRepo) *WeatherService {
	return &WeatherService{
		weatherClient: client,
		repo:          repo,
	}
}

func (s *WeatherService) GetWeatherAndSaveHistory(city string) (models.WeatherResponse, error) {
	weather, err := s.weatherClient.GetWeather(city)
	if err != nil {
		return models.WeatherResponse{}, err
	}

	err = s.repo.CreateHistory(city, weather.Main.Temperature, weather.Weather[0].Description, time.Now())
	if err != nil {
		return models.WeatherResponse{}, err
	}

	return weather, nil
}
