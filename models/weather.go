package models

import "time"

type WeatherResponse struct {
	Name    string        `json:"name"`
	Weather []WeatherInfo `json:"weather"`
	Main    Main          `json:"main"`
}

type WeatherInfo struct {
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type Main struct {
	Temperature float64 `json:"temp"`
	Humidity    int     `json:"humidity"`
	Pressure    int     `json:"pressure"`
}

type WeatherHistory struct {
	Temperature float64   `db:"temperature" json:"temperature"`
	Description string    `db:"description" json:"description"`
	CreatedAt   time.Time `db:"created_at"  json:"created_at"`
}
