package client

import (
	"encoding/json"
	"fmt"
	"github.com/Iusemywalk88/Weather/db"
	"github.com/Iusemywalk88/Weather/models"
	"net/http"
	"time"
)

type WeatherClient interface {
	GetWeather(city string) (models.WeatherResponse, error)
}

type weatherClient struct {
	database   *db.DB
	baseUrl    string
	token      string
	httpClient http.Client
}

func NewWeatherClient(db *db.DB, baseUrl, token string) WeatherClient {
	return &weatherClient{
		database: db,
		baseUrl:  baseUrl,
		token:    token,
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
		return models.WeatherResponse{}, fmt.Errorf("Server doesn`t respond")
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode != 200 {
		return models.WeatherResponse{}, fmt.Errorf("API returns error: %s", resp.Status)
	}

	var weatherResponse models.WeatherResponse

	if err := json.NewDecoder(resp.Body).Decode(&weatherResponse); err != nil {
		return models.WeatherResponse{}, err
	}

	w.database.CreateHistory(city, weatherResponse.Main.Temperature, weatherResponse.Weather[0].Description, time.Now())

	return weatherResponse, nil
}
