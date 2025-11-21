package requests

type FavouriteRequest struct {
	City string `json:"city" binding:"required"`
}
