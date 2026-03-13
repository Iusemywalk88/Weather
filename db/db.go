package db

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/Iusemywalk88/Weather/internal/config"
	"github.com/Iusemywalk88/Weather/models"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"time"
)

type DB struct {
	*sqlx.DB
}

type City struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}

func New(cfg config.Config) *DB {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPass, cfg.DBName)

	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	log.Println("Connected to PostgreSQL")
	return &DB{db}
}

func (db *DB) CreateUser(user *models.User) error {
	_, err := db.Exec("INSERT INTO public.users (email, password_hash) VALUES ($1, $2)",
		user.Email, user.PasswordHash)
	if err != nil {
		return err
	}
	return nil
}

func (db *DB) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := db.Get(&user, "SELECT id, email, password_hash FROM users WHERE email = $1", email)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (db *DB) GetCity(cityName string) (int, error) {
	var cityID int
	err := db.Get(&cityID, "SELECT id FROM cities WHERE name = $1", cityName)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			db.createCity(cityName)
		}
		return 0, err
	}
	return cityID, nil
}

func (db *DB) createCity(cityName string) (int, error) {
	var cityID int
	createErr := db.QueryRow("INSERT INTO cities (name) VALUES ($1) RETURNING id", cityName).Scan(&cityID)
	if createErr != nil {
		return 0, createErr
	}
	return cityID, nil
}

func (db *DB) AddFavourite(userID, cityID int) error {
	_, err := db.Exec("INSERT INTO favorite_cities (user_, city_id) VALUES ($1, $2)", userID, cityID)
	return err
}

func (db *DB) CheckAlreadyFavorite(userID, cityID int) (bool, error) {
	var exists bool
	err := db.Get(&exists, "SELECT EXISTS(SELECT 1 FROM favorite_cities WHERE user_id = $1 AND city_id = $2)", userID, cityID)
	if err != nil {
		return false, err
	}
	return exists, nil
}

func (db *DB) GetAllCities(userID int) ([]City, error) {
	var cityNames []City

	query := `
        SELECT c.name, c.id
        FROM cities c
		JOIN favorite_cities fc ON c.id = fc.city_id
		WHERE fc.user_id = $1`

	err := db.Select(&cityNames, query, userID)
	if err != nil {
		return nil, err
	}

	return cityNames, nil
}

func (db *DB) DeleteCity(userID int, cityID int) error {

	_, err := db.Exec("DELETE FROM favorite_cities WHERE user_id = $1 AND city_id = $2", userID, cityID)
	if err != nil {
		return err
	}
	return nil
}

func (db *DB) HistoryExistsForToday(cityID int) (bool, error) {
	var exists bool
	err := db.Get(&exists, "SELECT EXISTS(SELECT 1 FROM weather_history WHERE city_id = $1 AND created_at::date = now()::date)", cityID)
	if err != nil {
		return false, err
	}
	return exists, nil
}

func (db *DB) CreateHistory(
	cityName string,
	temperature float64,
	description string,
	createdAt time.Time) error {

	cityID, err := db.GetCity(cityName)
	if err != nil {
		return err
	}

	exists, err := db.HistoryExistsForToday(cityID)
	if err != nil {
		return err
	}
	if exists {
		return nil
	}

	_, err = db.Exec("INSERT INTO weather_history (city_id, temperature, description, created_at) VALUES ($1,$2,$3,$4)", cityID, temperature, description, createdAt)
	if err != nil {
		return err
	}
	return nil
}

func (db *DB) GetHistory(cityName string) ([]models.WeatherHistory, error) {
	var cityID int
	err := db.Get(&cityID, "SELECT id FROM cities WHERE name = $1", cityName)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	var history []models.WeatherHistory
	query := `
		SELECT temperature, description, created_at
		FROM weather_history
		WHERE city_id = $1
		ORDER BY created_at DESC`

	err = db.Select(&history, query, cityID)
	if err != nil {
		return nil, err
	}
	return history, nil
}
