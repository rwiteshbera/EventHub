package controllers

import (
	"eventCatalogService/api"
	"eventCatalogService/database"
	"eventCatalogService/models"
	"eventCatalogService/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateEvent(server *api.Server) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var eventInfo models.EventInformation

		db, err := database.ConnectPostgres(server)
		if err != nil {
			LogError(ctx, http.StatusInternalServerError, err, 01)
			return
		}

		err = db.AutoMigrate(&models.EventInformation{})
		if err != nil {
			LogError(ctx, http.StatusInternalServerError, err, 02)
			return
		}

		err = ctx.ShouldBindJSON(&eventInfo)
		if err != nil {
			LogError(ctx, http.StatusBadRequest, err, 03)
			return
		}

		eventInfo.ID, err = utils.GenerateUUID()
		if err != nil {
			LogError(ctx, http.StatusBadRequest, err, 04)
			return
		}

		result := db.Create(&eventInfo)
		if result.Error != nil {
			LogError(ctx, http.StatusInternalServerError, result.Error, 05)
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"message": eventInfo.ID})

	}
}
