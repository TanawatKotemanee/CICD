package config

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"os"
)

var DB *sqlx.DB

func InitDB() {
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_SSLMODE"),
	)

	// Connect to the database
	var err error
	DB, err = sqlx.Connect("postgres", connStr)
	if err != nil {
		if os.Getenv("CONNECT_DB") == "true" {
			log.Fatalf("Failed to connect to the database: %v", err)
		}
		log.Printf("Failed to connect to the database: %v", err)
		DB = nil
		return
	}

	log.Println("Database connection established successfully")
}
