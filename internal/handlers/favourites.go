package handlers

import (
	"github.com/Iusemywalk88/Weather/db"
	"github.com/Iusemywalk88/Weather/models/requests"
	"github.com/Iusemywalk88/Weather/models/responses"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"net/http"
)

type FavouritesHandler struct {
	DB *db.DB
}

func NewFavouritesHandler(db *db.DB) *FavouritesHandler {
	return &FavouritesHandler{DB: db}
}

func (f *FavouritesHandler) AddFavourites(c *gin.Context) {
	var fav requests.FavouriteRequest
	if err := c.ShouldBindJSON(&fav); err != nil {
		c.JSON(http.StatusBadRequest, responses.BaseResponse{Error: "Invalid request body: " + err.Error()})
		return
	}

	userIDUntyped, ok := c.Get("userID")
	if !ok {
		c.JSON(http.StatusUnauthorized, responses.BaseResponse{Error: "Unauthorized"})
		return
	}

	userID, ok := userIDUntyped.(int)
	if !ok {
		c.JSON(http.StatusUnauthorized, responses.BaseResponse{Error: "Problem at parcing user id"})
		return
	}

	cityID, err := f.DB.GetOrCreateCity(fav.City)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.BaseResponse{Error: "Problem creating city: " + err.Error()})
		return
	}

	err = f.DB.AddFavourite(userID, cityID)
	if err != nil {
		if pgErr, ok := err.(*pq.Error); ok && pgErr.Code == "23505" {
			c.JSON(http.StatusConflict, responses.BaseResponse{Error: "City already in favourites"})
			return
		}
		c.JSON(http.StatusInternalServerError, responses.BaseResponse{Error: "Failed to add favourite: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, responses.BaseResponse{Message: "City added to favourites successfully"})
}

func getFavourites(c *gin.Context) {

}
