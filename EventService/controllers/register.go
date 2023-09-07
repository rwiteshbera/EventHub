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

		// Get the event_id from URL query
		event_id := context.Query("event_id")

		var registrationDetails models.RegisteredParticipantInformation

		// Get participant Email from context
		participantEmail, isExists := context.Get(middlewares.AuthorizationPayloadKey)
		if !isExists {
			LogError(context, http.StatusInternalServerError, errors.New("unauthorized"), 01)
			return
		}

		// Connect with Postgres
		db, err := database.ConnectPostgres(server)
		if err != nil {
			LogError(context, http.StatusInternalServerError, err, 02)
			return
		}

		// Migrate Model
		err = db.AutoMigrate(&models.RegisteredParticipantInformation{})
		if err != nil {
			LogError(context, http.StatusInternalServerError, err, 03)
			return
		}

		// Check if the organizer is trying to register or not?
		// Organizer cannot join as participant
		var event_informations models.DisplayEventInfo
		db.Table("event_informations").First(&event_informations, "id = ?", event_id)

		if event_informations.OrganizerEmail == participantEmail {
			// // Organizer cannot join as participant
			LogError(context, http.StatusOK, errors.New("organizer cannot join as participant. use different account"), 04)
			return
		}

		// Check if the user already registerd on the event or not
		// No duplicate registration is allowed
		rowsAffected := db.Table("registered_participant_informations").Find(&models.RegisteredParticipantInformation{}, "user_email = ? AND event_id = ?", participantEmail, event_id).RowsAffected

		if rowsAffected > 0 {
			// If already registered
			context.JSON(http.StatusOK, gin.H{"message": "already registered"})
			return
		}

		// Add details
		registrationDetails.UserEmail = participantEmail.(string)
		registrationDetails.EventId = event_id

		// Registartion for event
		// Store registration details in postgres
		result := db.Create(&registrationDetails)
		if result.Error != nil {
			LogError(context, http.StatusInternalServerError, result.Error, 06)
			return
		}

		// Send response
		context.JSON(http.StatusOK, registrationDetails)
	}
}
