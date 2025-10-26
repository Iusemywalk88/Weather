package client

import (
	"encoding/json"
	"fmt"
	"github.com/Iusemywalk88/Weather/models"
	"net/http"
	"time"
)

type WeatherClient interface {
	GetWeather(city string) (models.WeatherResponse, error)
}

type weatherClient struct {
	baseUrl    string
	token      string
	httpClient http.Client
}

func NewWeatherClient(baseUrl, token string) WeatherClient {
	return &weatherClient{
		baseUrl: baseUrl,
		token:   token,
		httpClient: http.Client{
			Timeout: time.Second * 10,
		},
	}
}

func (w *weatherClient) GetWeather(city string) (models.WeatherResponse, error) {
	req, err := http.NewRequest(http.MethodGet, w.baseUrl, nil)
	if err != nil {
		return models.WeatherResponse{}, err
	}

	values := req.URL.Query()
	values.Add("q", city)
	values.Add("APPID", w.token)
	values.Add("units", "metric")
	values.Add("lang", "ru")

	req.URL.RawQuery = values.Encode()

	resp, err := w.httpClient.Do(req)

	if err != nil {
		return models.WeatherResponse{}, fmt.Errorf("сервер погоды не отвежает, попробуйте позже")
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode != 200 {
		return models.WeatherResponse{}, fmt.Errorf("API вернуло ошибку: %s", resp.Status)
	}

	var weatherResponse models.WeatherResponse

	if err := json.NewDecoder(resp.Body).Decode(&weatherResponse); err != nil {
		return models.WeatherResponse{}, err
	}

	return weatherResponse, nil
}
