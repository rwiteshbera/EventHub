package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	SENDER_GMAIL    string
	SENDER_PASSWORD string
	SERVER_HOST     string
	SERVER_PORT     string
	RABBITMQ        string
}

func LoadConfig() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println(err.Error())
	}

	config := &Config{
		SENDER_GMAIL:    os.Getenv("SENDER_GMAIL"),
		SENDER_PASSWORD: os.Getenv("SENDER_PASSWORD"),
		SERVER_HOST:     os.Getenv("SERVER_HOST"),
		SERVER_PORT:     os.Getenv("SERVER_PORT"),
		RABBITMQ:        os.Getenv("RABBITMQ"),
	}

	return config
}
