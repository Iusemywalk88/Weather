package handlers

import (
	"github.com/Iusemywalk88/Weather/internal/services"
	"github.com/Iusemywalk88/Weather/models/requests"
	"github.com/Iusemywalk88/Weather/models/responses"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthHandler struct {
	AuthService services.AuthServiceInterface
	JWTKey      []byte
}

func NewAuthHandler(authService services.AuthServiceInterface, jwtKey []byte) *AuthHandler {
	return &AuthHandler{AuthService: authService, JWTKey: jwtKey}
}

func (a *AuthHandler) RegisterUser(c *gin.Context) {
	var req requests.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, responses.BaseResponse{Error: err.Error()})
		return
	}

	_, err := a.AuthService.Register(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.BaseResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, responses.BaseResponse{Message: "User created successfully"})
}

func (a *AuthHandler) LoginUser(c *gin.Context) {
	var req requests.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, responses.BaseResponse{Error: "Invalid request payload"})
		return
	}

	tokenString, err := a.AuthService.Login(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, responses.BaseResponse{Error: "Invalid credentials"})
		return
	}

	c.JSON(http.StatusOK, responses.BaseResponse{Data: responses.LoginResponse{Token: tokenString}})
}
