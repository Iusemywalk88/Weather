package middleware

import (
	"github.com/Iusemywalk88/Weather/internal/constants"
	"github.com/Iusemywalk88/Weather/models/responses"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"strings"
)

type Middleware struct {
	JWTKey []byte
}

func NewMiddleware(jwtKey []byte) *Middleware {
	return &Middleware{JWTKey: jwtKey}
}

func (m *Middleware) AuthMiddleware() gin.HandlerFunc {
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
			return m.JWTKey, nil
		})

		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, responses.BaseResponse{Error: "Invalid token"})
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			if sub, ok := claims["sub"].(float64); ok {
				c.Set(constants.ContextKeyUserID, int(sub))
				c.Next()
			} else {
				c.AbortWithStatusJSON(http.StatusUnauthorized, responses.BaseResponse{Error: "Invalid token claims"})
			}
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, responses.BaseResponse{Error: "Invalid token claims"})
		}
	}
}
