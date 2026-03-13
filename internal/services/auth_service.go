package services

import (
	"github.com/Iusemywalk88/Weather/db"
	"github.com/Iusemywalk88/Weather/models"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"time"
)

//go:generate mockgen -destination=mocks/mock_auth_service.go -package=mocks github.com/Iusemywalk88/Weather/internal/services AuthServiceInterface

type AuthServiceInterface interface {
	Register(email, password string) (*models.User, error)
	Login(email, password string) (string, error)
}

const TokenTTLInHours = 24

type AuthService struct {
	DB     *db.DB
	JWTKey []byte
}

func NewAuthService(db *db.DB, jwtKey []byte) *AuthService {
	return &AuthService{DB: db, JWTKey: jwtKey}
}

func (s AuthService) Register(email, password string) (*models.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Email:        email,
		PasswordHash: string(hashedPassword),
		CreatedAt:    time.Now(),
	}

	if err := s.DB.CreateUser(user); err != nil {
		return nil, err
	}
	return user, nil
}

func (s AuthService) Login(email, password string) (string, error) {
	user, err := s.DB.GetUserByEmail(email)
	if err != nil {
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * TokenTTLInHours).Unix(),
	})

	tokenString, err := token.SignedString(s.JWTKey)
	if err != nil {

		return "", err
	}

	return tokenString, nil
}
