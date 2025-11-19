package handlers

import (
	"github.com/Iusemywalk88/Weather/internal/services"
	"github.com/Iusemywalk88/Weather/models/requests"
	"github.com/Iusemywalk88/Weather/models/responses"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"strings"
)

type AuthHandler struct {
	AuthService *services.AuthService
	JWTKey      []byte
}

func NewAuthHandler(authService *services.AuthService, jwtKey []byte) *AuthHandler {
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

func (a *AuthHandler) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, responses.BaseResponse{Error: "Invalid authorization header"})
			return
		}

		authHeader = strings.TrimPrefix(authHeader, "Bearer ")

		token, err := jwt.Parse(authHeader, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, http.ErrAbortHandler
			}
			return a.JWTKey, nil
		})

		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, responses.BaseResponse{Error: "Invalid token"})
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			if sub, ok := claims["sub"].(float64); ok {
				c.Set("userID", int(sub))
				c.Next()
			} else {
				c.AbortWithStatusJSON(http.StatusUnauthorized, responses.BaseResponse{Error: "Invalid token claims"})
			}
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, responses.BaseResponse{Error: "Invalid token claims"})
		}
	}
}
