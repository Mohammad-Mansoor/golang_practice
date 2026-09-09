package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PORT   string
	ENV    string
	DB_URL string
}

func MustLoad() Config {
	godotenv.Load()

	PORT := os.Getenv("PORT")
	if PORT == "" {
		panic("PORT is Required")
	}

	ENV := os.Getenv("ENV")
	if ENV == "" {
		panic("ENV is Required")
	}

	DB_URL := os.Getenv("DB_URL")
	if DB_URL == "" {
		panic("DB_URL is Required")
	}

	return Config{
		PORT:   PORT,
		ENV:    ENV,
		DB_URL: DB_URL,
	}

}
