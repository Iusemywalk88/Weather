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
	"testing"
)

func TestWeatherHandler_HandleWeather_EmptyCity(t *testing.T) {
	// Arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockWeatherServiceInterface(ctrl)
	handler := handlers.NewWeatherHandler(mockService)

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Request = httptest.NewRequest(http.MethodGet, "/weather/", nil)
	c.Params = []gin.Param{{Key: "city", Value: ""}}

	expectedStatus := 400
	expectedError := "Nothing was written"
	// Act
	handler.HandleWeather(c)
	// Assert
	assert.Equal(t, expectedStatus, w.Code)
	assert.Contains(t, w.Body.String(), expectedError)
}

func TestWeatherHandler_HandleWeather_Success(t *testing.T) {
	// Arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockWeatherServiceInterface(ctrl)
	mockService.EXPECT().GetWeatherAndSaveHistory("Moscow").Return(
		models.WeatherResponse{
			Name:    "Moscow",
			Main:    models.Main{Temperature: 25.5},
			Weather: []models.WeatherInfo{{Description: "clear"}},
		},
		nil,
	)

	handler := handlers.NewWeatherHandler(mockService)

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Request = httptest.NewRequest(http.MethodGet, "/weather/", nil)
	c.Params = []gin.Param{{Key: "city", Value: "Moscow"}}

	expectedStatus := 200

	// Act
	handler.HandleWeather(c)
	// Assert
	assert.Equal(t, expectedStatus, w.Code)
	assert.Contains(t, w.Body.String(), "data")
	assert.Contains(t, w.Body.String(), "Moscow")
}

func TestWeatherHandler_HandleWeather_ServiceError(t *testing.T) {
	//Arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockWeatherServiceInterface(ctrl)
	mockService.EXPECT().GetWeatherAndSaveHistory("Moscow").Return(
		models.WeatherResponse{},
		errors.New("API not available"))

	handler := handlers.NewWeatherHandler(mockService)

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Request = httptest.NewRequest(http.MethodGet, "/weather/", nil)
	c.Params = []gin.Param{{Key: "city", Value: "Moscow"}}

	expectedStatus := 500

	//Act
	handler.HandleWeather(c)

	// Assert
	assert.Equal(t, expectedStatus, w.Code)
}
