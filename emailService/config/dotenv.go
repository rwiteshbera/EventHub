package config

import (
	"os"
)

type Config struct {
	SENDER_GMAIL     string
	SENDER_PASSWORD  string
	SERVER_HOST      string
	SERVER_PORT      string
	MEMPHIS_HOST     string
	MEMPHIS_PASSWORD string
	MEMPHIS_USERNAME string
}

func LoadConfig() (*Config, error) {
	config := &Config{
		SENDER_GMAIL:     os.Getenv("SENDER_GMAIL"),
		SENDER_PASSWORD:  os.Getenv("SENDER_PASSWORD"),
		SERVER_HOST:      os.Getenv("SERVER_HOST"),
		SERVER_PORT:      os.Getenv("SERVER_PORT"),
		MEMPHIS_HOST:     os.Getenv("MEMPHIS_HOST"),
		MEMPHIS_USERNAME: os.Getenv("MEMPHIS_USERNAME"),
		MEMPHIS_PASSWORD: os.Getenv("MEMPHIS_PASSWORD"),
	}

	return config, nil
}
