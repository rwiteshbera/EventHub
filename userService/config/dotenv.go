package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DB_ADDRESS  string
	DB_PASSWORD string
	SERVER_HOST string
	SERVER_PORT string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return nil, err
	}

	config := &Config{
		DB_ADDRESS:  os.Getenv("DB_ADDRESS"),
		DB_PASSWORD: os.Getenv("DB_PASSWORD"),
		SERVER_HOST: os.Getenv("SERVER_HOST"),
		SERVER_PORT: os.Getenv("SERVER_PORT"),
	}

	return config, nil
}
