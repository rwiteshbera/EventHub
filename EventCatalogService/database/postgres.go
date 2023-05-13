package database

import (
	"eventCatalogService/api"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectPostgres(server *api.Server) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		server.Config.POSTGRES_HOST, server.Config.POSTGRES_USER, server.Config.POSTGRES_PASSWORD, server.Config.POSTGRES_DB,
		server.Config.POSTGRES_PORT)
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: dsn, PreferSimpleProtocol: true,
	}), &gorm.Config{})

	if err != nil {
		return nil, err
	}
	return db, nil
}
