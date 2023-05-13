package controllers

import (
	"database/sql"
	"eventCatalogService/api"
	"eventCatalogService/database"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateEvent(server *api.Server) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		db, err := database.ConnectPostgres(server)
		d, _ := db.DB()
		defer func(d *sql.DB) {
			err := d.Close()
			if err != nil {
				LogError(ctx, http.StatusInternalServerError, err)
				return
			}
		}(d)
		
		if err != nil {
			LogError(ctx, http.StatusInternalServerError, err)
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "success"})
	}
}
