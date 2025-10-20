package models

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
