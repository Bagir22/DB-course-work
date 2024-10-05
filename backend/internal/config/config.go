package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	PgPort     string
	PgUser     string
	PgPassword string
	PgDatabase string

	ServerPort string
	ServerHost string
}

func InitConfig() *Config {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Can't loading .env")
	}

	cfg := &Config{
		PgPort:     os.Getenv("DB_PORT"),
		PgUser:     os.Getenv("DB_USER"),
		PgPassword: os.Getenv("DB_PASSWORD"),
		PgDatabase: os.Getenv("DB_DATABASE"),

		ServerPort: os.Getenv("SERVER_PORT"),
		ServerHost: os.Getenv("SERVER_HOST"),
	}

	return cfg
}
