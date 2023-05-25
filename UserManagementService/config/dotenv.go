package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	REDIS_DB_ADDRESS  string
	REDIS_DB_PASSWORD string
	MONGO_DB_URI      string
	SERVER_HOST       string
	SERVER_PORT       string
	JWT_SECRET        string
	RABBITMQ          string
}

func LoadConfig() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println(err.Error())
	}

	config := &Config{
		REDIS_DB_ADDRESS:  os.Getenv("REDIS_DB_ADDRESS"),
		REDIS_DB_PASSWORD: os.Getenv("REDIS_DB_PASSWORD"),
		MONGO_DB_URI:      os.Getenv("MONGO_DB_URI"),
		SERVER_HOST:       os.Getenv("SERVER_HOST"),
		SERVER_PORT:       os.Getenv("SERVER_PORT"),
		JWT_SECRET:        os.Getenv("JWT_SECRET"),
		RABBITMQ:          os.Getenv("RABBITMQ"),
	}

	return config
}
