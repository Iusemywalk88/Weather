package db

import (
	"fmt"
	"github.com/Iusemywalk88/Weather/models"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"os"
)

type DB struct {
	*sqlx.DB
}

func Connect() *DB {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

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

func (d *DB) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := d.Get(&user, "SELECT id, email, password_hash FROM users WHERE email = $1", email)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
