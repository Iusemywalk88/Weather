package client

import (
	"encoding/json"
	"fmt"
	"github.com/Iusemywalk88/Weather/models"
	"io/ioutil"
	"net/http"
)

const (
	apiKey  = "9b4430c225c9caf186f7f5a81414f451"
	baseURL = "https://api.openweathermap.org/data/2.5/weather"
)

func GetWeather(city string) (models.WeatherResponse, error) {
	url := baseURL + "?q=" + city + "&APPID=" + apiKey + "&units=metric&lang=ru"
	resp, err := http.Get(url)
	if err != nil {
		return models.WeatherResponse{}, err
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
