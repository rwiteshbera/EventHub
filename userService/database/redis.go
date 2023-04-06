package database

import (
	"context"
	"userService/config"

	"github.com/redis/go-redis/v9"
)

var Ctx = context.Background()

func CreateClient(config *config.Config, databaseNumber int) *redis.Client {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     config.DB_ADDRESS,
		Password: config.DB_PASSWORD,
		DB:       databaseNumber,
	})

	return redisClient
}
