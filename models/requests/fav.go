package requests

type FavouriteRequest struct {
	City string `json:"city" binding:"required"`
}

type DeleteCityRequest struct {
	CityId int `json:"cityId" binding:"required"`
}
