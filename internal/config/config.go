package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port                 string
	DBUser               string
	DBPassword           string
	DBName               string
	DBHost               string
	DBPort               string
	FirebaseCredBase64   string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Println(".env file not found")
	}

	config := &Config{
		Port:               os.Getenv("PORT"),
		DBUser:             os.Getenv("DB_USER"),
		DBPassword:         os.Getenv("DB_PASSWORD"),
		DBName:             os.Getenv("DB_NAME"),
		DBHost:             os.Getenv("DB_HOST"),
		DBPort:             os.Getenv("DB_PORT"),
		FirebaseCredBase64: os.Getenv("FIREBASE_CREDENTIALS_BASE64"),
	}

	return config, nil
}
