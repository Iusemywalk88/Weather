package handlers

import (
	"github.com/Iusemywalk88/Weather/internal/services"
	"github.com/Iusemywalk88/Weather/models/handlers"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthHandler struct {
	AuthService *services.AuthService
}

func NewAuthHandler(authService *services.AuthService) *AuthHandler {
	return &AuthHandler{AuthService: authService}
}

func (a *AuthHandler) RegisterUser(c *gin.Context) {
	var req handlers.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, handlers.BaseResponse{Error: err.Error()})
		return
	}

	_, err := a.AuthService.Register(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, handlers.BaseResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, handlers.BaseResponse{Message: "User created successfully"})
}

func (a *AuthHandler) LoginUser(c *gin.Context) {
	var req handlers.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, handlers.BaseResponse{Error: "Invalid request payload"})
		return
	}

	tokenString, err := a.AuthService.Login(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, handlers.BaseResponse{Error: "Invalid credentials"})
		return
	}

	c.JSON(http.StatusOK, handlers.BaseResponse{Data: handlers.LoginResponse{Token: tokenString}})
}
