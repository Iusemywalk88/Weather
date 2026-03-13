package handlers_test

import (
	"errors"
	"github.com/Iusemywalk88/Weather/internal/handlers"
	"github.com/Iusemywalk88/Weather/internal/services/mocks"
	"github.com/Iusemywalk88/Weather/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestAuthHandler_RegisterUser_InvalidJSON(t *testing.T) {
	//Arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockAuthServiceInterface(ctrl)
	handler := handlers.NewAuthHandler(mockService, []byte("test-key"))

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Request = httptest.NewRequest("POST", "/api/user/register", nil)
	c.Request.Header.Set("Content-Type", "application/json")

	//Act
	handler.RegisterUser(c)

	//Assert
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestAuthHandler_RegisterUser_Success(t *testing.T) {
	//Arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockAuthServiceInterface(ctrl)
	mockService.EXPECT().Register("test@example.com", "123123").Return(
		&models.User{
			ID:    1,
			Email: "test@example.com",
		},
		nil)

	handler := handlers.NewAuthHandler(mockService, []byte("test-key"))

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	body := `{"email": "test@example.com", "password": "123123"}`
	c.Request = httptest.NewRequest(http.MethodPost, "/register", strings.NewReader(body))
	c.Request.Header.Set("Content-Type", "application/json")

	//Act
	handler.RegisterUser(c)

	//Assert
	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Contains(t, w.Body.String(), "User created successfully")
}

func TestAuthHandler_LoginUser_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockAuthServiceInterface(ctrl)
	mockService.EXPECT().Login("test@example.com", "123123").Return(
		"some-JWT-token",
		nil)

	handler := handlers.NewAuthHandler(mockService, []byte("some-JWT-token"))

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	body := `{"email": "test@example.com", "password": "123123"}`
	c.Request = httptest.NewRequest("POST", "/api/user/login", strings.NewReader(body))
	c.Request.Header.Set("Content-Type", "application/json")

	//Act
	handler.LoginUser(c)

	//Assert
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "token")
}

func TestAuthHandler_LoginUser_WrongPassword(t *testing.T) {
	//Arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockAuthServiceInterface(ctrl)
	mockService.EXPECT().Login("test@example.com", "Wrong-password").Return("", errors.New("wrong password"))

	handler := handlers.NewAuthHandler(mockService, []byte("test-key"))

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	body := `{"email": "test@example.com", "password": "Wrong-password"}`
	c.Request = httptest.NewRequest("POST", "/api/user/login", strings.NewReader(body))
	c.Request.Header.Set("Content-Type", "application/json")

	//Act
	handler.LoginUser(c)

	//Assert
	assert.Equal(t, http.StatusUnauthorized, w.Code)
	assert.Contains(t, w.Body.String(), "Invalid credentials")

}
