package db

import (
	"fmt"
	"github.com/Iusemywalk88/Weather/internal/config"
	"github.com/Iusemywalk88/Weather/models"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
)

type DB struct {
	*sqlx.DB
}

func New() *DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Error loading config")
	}

	host := cfg.DBHost
	port := cfg.DBPort
	user := cfg.DBUser
	password := cfg.DBPass
	dbname := cfg.DBName

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	log.Println("Connected to PostgreSQL")
	return &DB{db}
}

func (db *DB) CreateUser(user *models.User) error {
	_, err := db.Exec("INSERT INTO public.users (email, password_hash) VALUES ($1, $2)", user.Email, user.PasswordHash)
	return err
}

func (db *DB) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := db.Get(&user, "SELECT id, email, password_hash FROM users WHERE email = $1", email)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
