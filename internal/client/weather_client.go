package client

import (
	"encoding/json"
	"fmt"
	"github.com/Iusemywalk88/Weather/internal/config"
	"github.com/Iusemywalk88/Weather/models"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	baseURL = "https://api.openweathermap.org/data/2.5/weather"
)

func GetWeather(city string) (models.WeatherResponse, error) {

	cfg := config.Load()

	client := http.Client{
		Timeout: time.Second * 10,
	}

	url := baseURL + "?q=" + city + "&APPID=" + cfg.WeatherAPIKey + "&units=metric&lang=ru"
	resp, err := client.Get(url)
	if err != nil {
		return models.WeatherResponse{}, fmt.Errorf("сервер погоды не отвежает, попробуйте позже")
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return models.WeatherResponse{}, fmt.Errorf("API вернуло ошибку: %s", resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return models.WeatherResponse{}, err
	}

	var weatherResponse models.WeatherResponse
	if err := json.Unmarshal(body, &weatherResponse); err != nil {
		return models.WeatherResponse{}, err
	}

	return weatherResponse, nil
}
