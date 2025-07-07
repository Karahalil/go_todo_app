//This file will include the configuration for the application

package config

import (
	"os"

	"github.com/joho/godotenv"
)

var (
	DBUser     string
	DBPassword string
	DBName     string
)

func Load() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file")
	}
	DBUser = os.Getenv("DB_USER")
	DBPassword = os.Getenv("DB_PASSWORD")
	DBName = os.Getenv("DB_NAME")
	if DBUser == "" || DBPassword == "" || DBName == "" {
		panic("Database configuration is not set. Please set DB_USER, DB_PASSWORD, and DB_NAME environment variables.")
	}
}
