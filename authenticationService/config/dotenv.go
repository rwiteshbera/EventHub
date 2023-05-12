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
	MEMPHIS_HOST      string
	MEMPHIS_PASSWORD  string
	MEMPHIS_USERNAME  string
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
		MEMPHIS_HOST:      os.Getenv("MEMPHIS_HOST"),
		MEMPHIS_USERNAME:  os.Getenv("MEMPHIS_USERNAME"),
		MEMPHIS_PASSWORD:  os.Getenv("MEMPHIS_PASSWORD"),
	}

	return config
}
