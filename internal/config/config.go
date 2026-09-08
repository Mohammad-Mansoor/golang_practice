package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PORT string
	ENV  string
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

	return Config{
		PORT: PORT,
		ENV:  ENV,
	}

}
