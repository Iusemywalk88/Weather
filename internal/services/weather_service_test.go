package services_test

import (
	"errors"
	"testing"

	"github.com/Iusemywalk88/Weather/internal/services"
	"github.com/Iusemywalk88/Weather/internal/services/mocks"
	"github.com/Iusemywalk88/Weather/models"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestWeatherService_GetWeatherAndSaveHistory_Success(t *testing.T) {
	// Arrange (подготовка)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Создаём моки
	mockClient := mocks.NewMockWeatherClient(ctrl)
	mockRepo := mocks.NewMockWeatherRepo(ctrl)

	// Ожидаемые данные
	expectedCity := "Moscow"
	expectedWeather := models.WeatherResponse{
		Main: models.Main{
			Temperature: 25.5,
		},
		Weather: []models.WeatherInfo{
			{Description: "clear sky"},
		},
	}

	// Настраиваем ожидания для мока клиента
	mockClient.EXPECT().GetWeather(expectedCity).Return(expectedWeather, nil)

	// Настраиваем ожидания для мока репозитория (БД)
	mockRepo.EXPECT().CreateHistory(
		expectedCity,
		expectedWeather.Main.Temperature,
		expectedWeather.Weather[0].Description,
		gomock.Any(), // время не проверяем точно
	).Return(nil)

	// Создаём сервис с моками
	service := services.NewWeatherService(mockClient, mockRepo)

	// Act (действие)
	result, err := service.GetWeatherAndSaveHistory(expectedCity)

	// Assert (проверка)
	assert.NoError(t, err)
	assert.Equal(t, expectedWeather.Main.Temperature, result.Main.Temperature)
	assert.Equal(t, expectedWeather.Weather[0].Description, result.Weather[0].Description)
}

func TestWeatherService_GetWeatherAndSaveHistory_ClientError(t *testing.T) {
	// Arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mocks.NewMockWeatherClient(ctrl)
	mockRepo := mocks.NewMockWeatherRepo(ctrl)

	expectedCity := "Moscow"
	expectedError := errors.New("API not available")

	// Ожидаем ошибку от клиента, БД не должна вызываться
	mockClient.EXPECT().GetWeather(expectedCity).Return(models.WeatherResponse{}, expectedError)
	// mockRepo.EXPECT().CreateHistory(...) — не вызывается, поэтому не указываем

	service := services.NewWeatherService(mockClient, mockRepo)

	// Act
	result, err := service.GetWeatherAndSaveHistory(expectedCity)

	// Assert
	assert.Error(t, err)
	assert.Equal(t, expectedError, err)
	assert.Empty(t, result)
}

func TestWeatherService_GetWeatherAndSaveHistory_RepoError(t *testing.T) {
	// Arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := mocks.NewMockWeatherClient(ctrl)
	mockRepo := mocks.NewMockWeatherRepo(ctrl)

	expectedCity := "Moscow"
	weatherData := models.WeatherResponse{
		Main: models.Main{
			Temperature: 20.0,
		},
		Weather: []models.WeatherInfo{
			{Description: "rain"},
		},
	}
	repoError := errors.New("database connection failed")

	// Клиент успешно возвращает данные
	mockClient.EXPECT().GetWeather(expectedCity).Return(weatherData, nil)

	// Репозиторий возвращает ошибку
	mockRepo.EXPECT().CreateHistory(
		expectedCity,
		weatherData.Main.Temperature,
		weatherData.Weather[0].Description,
		gomock.Any(),
	).Return(repoError)

	service := services.NewWeatherService(mockClient, mockRepo)

	// Act
	result, err := service.GetWeatherAndSaveHistory(expectedCity)

	// Assert
	assert.Error(t, err)
	assert.Equal(t, repoError, err)
	assert.Empty(t, result) // Возвращаем пустой результат при ошибке
}
