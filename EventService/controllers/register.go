package controllers

import (
	"errors"
	"eventCatalogService/api"
	"eventCatalogService/database"
	"eventCatalogService/middlewares"
	"eventCatalogService/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterEvent(server *api.Server) gin.HandlerFunc {
	return func(context *gin.Context) {
		event_id := context.Query("event_id")
		participantEmail, isExists := context.Get(middlewares.AuthorizationPayloadKey)
		if !isExists {
			LogError(context, http.StatusInternalServerError, errors.New("unauthorized"), 01)
			return
		}

		db, err := database.ConnectPostgres(server)
		if err != nil {
			LogError(context, http.StatusInternalServerError, err, 02)
			return
		}

		err = db.AutoMigrate(&models.RegisteredParticipantInformation{})
		if err != nil {
			LogError(context, http.StatusInternalServerError, err, 03)
			return
		}

		var event_informations models.DisplayEventInfo
		// db.Where("id = ?", event_id).First(&event_informations)
		db.Table("event_informations").First(&event_informations, "id = ?", event_id)

		if event_informations.OrganizerEmail == participantEmail {
			LogError(context, http.StatusOK, errors.New("organizer cannot join as participant. use different account"), 04)
			return
		}

		context.JSON(http.StatusOK, "awesome")
	}
}
