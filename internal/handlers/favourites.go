package handlers

import (
	"github.com/Iusemywalk88/Weather/db"
	"github.com/Iusemywalk88/Weather/internal/client"
	"github.com/Iusemywalk88/Weather/internal/constants"
	"github.com/Iusemywalk88/Weather/models/requests"
	"github.com/Iusemywalk88/Weather/models/responses"
	"github.com/gin-gonic/gin"
	"net/http"
)

type FavouritesHandler struct {
	DB            *db.DB
	WeatherClient client.WeatherClient
}

func NewFavouritesHandler(db *db.DB, client client.WeatherClient) *FavouritesHandler {
	return &FavouritesHandler{
		DB:            db,
		WeatherClient: client,
	}
}

func (f *FavouritesHandler) AddFavourites(c *gin.Context) {
	var fav requests.FavouriteRequest

	if err := c.ShouldBindJSON(&fav); err != nil {
		c.JSON(http.StatusBadRequest, responses.BaseResponse{Error: "Invalid request body: " + err.Error()})
		return
	}

	userID := c.GetInt(constants.ContextKeyUserID)

	cityID, err := f.DB.GetOrCreateCity(fav.City)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.BaseResponse{Error: "Problem creating city: " + err.Error()})
		return
	}

	isAlreadyFavorite, err := f.DB.CheckAlreadyFavorite(userID, cityID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.BaseResponse{Error: "Failed to check if city is already favourite: " + err.Error()})
		return
	}
	if isAlreadyFavorite {
		c.JSON(http.StatusConflict, responses.BaseResponse{Error: "City already in favourites"})
		return
	}

	err = f.DB.AddFavourite(userID, cityID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.BaseResponse{Error: "Failed to add favourite: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, responses.BaseResponse{Message: "City added to favourites successfully"})
}

func (f *FavouritesHandler) GetFavourites(c *gin.Context) {
	userID := c.GetInt(constants.ContextKeyUserID)
	cities, err := f.DB.GetAllCities(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.BaseResponse{Error: "Failed to get favourites: " + err.Error()})
		return
	}

	var favouriteWeather []responses.FavouriteWeatherResponse

	for _, city := range cities {
		weather, err := f.WeatherClient.GetWeather(city.Name)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.BaseResponse{Error: "Failed to get weather: " + city.Name + err.Error()})
			continue
		}
		favouriteWeather = append(favouriteWeather, responses.FavouriteWeatherResponse{City: city.Name, Weather: weather})
	}

	c.JSON(http.StatusOK, responses.BaseResponse{Data: favouriteWeather})
}
