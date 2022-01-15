package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var connection *gorm.DB

func connect() (*gorm.DB, error) {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%d/%s", dbUser, dbPassword, dbHost, 5432, dbName)
	// dsn := fmt.Sprintf(
	// 	"host=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=America/Fortaleza",
	// 	dbHost,
	// 	dbUser,
	// 	dbPassword,
	// 	dbName,
	// )
	return gorm.Open(postgres.Open(dbURL), &gorm.Config{})
}

func GetConnection() (*gorm.DB, error) {
	if connection == nil {
		conn, err := connect()
		if err != nil {
			return nil, err
		}
		connection = conn
	}
	return connection, nil
}
