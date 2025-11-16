package handlers

import (
	"github.com/Iusemywalk88/Weather/models/handlers"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Favourite struct {
	Name string `json:"city"`
}

func addFavourites(c *gin.Context) {
	var fav Favourite
	if err := c.BindJSON(&fav); err != nil {
		c.JSON(http.StatusBadRequest, handlers.BaseResponse{Error: err.Error()})
	}

}

func getFavourites(c *gin.Context) {

}
