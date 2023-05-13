package controllers

import (
	"eventCatalogService/api"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
)

func CreateEvent(server *api.Server) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
			server.Config.POSTGRES_HOST, server.Config.POSTGRES_USER, server.Config.POSTGRES_PASSWORD, server.Config.POSTGRES_DB,
			server.Config.POSTGRES_PORT)
		_, err := gorm.Open(postgres.New(postgres.Config{
			DSN: dsn, PreferSimpleProtocol: true,
		}), &gorm.Config{})

		if err != nil {
			LogError(ctx, http.StatusInternalServerError, err)
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"message": "success"})
	}
}
