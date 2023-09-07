package controllers

import (
	"errors"
	"eventCatalogService/api"
	"eventCatalogService/database"
	"eventCatalogService/middlewares"
	"eventCatalogService/models"
	"eventCatalogService/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// Create a new event
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

		// Check if the date are valid or not
		if time.Now().After(eventInfo.StartTime) || time.Now().After(eventInfo.EndTime) {
			LogError(ctx, http.StatusBadRequest, errors.New("provide valid date"), 06)
			return
		}
		if eventInfo.StartTime.After(eventInfo.EndTime) {
			LogError(ctx, http.StatusBadRequest, errors.New("provide valid date"), 07)
			return
		}

		// Generate UUID for event
		eventInfo.ID, err = utils.GenerateUUID()
		if err != nil {
			LogError(ctx, http.StatusBadRequest, err, 04)
			return
		}

		// Fetch organizer email
		organizerPayload, _ := ctx.Get(middlewares.AuthorizationPayloadKey)
		eventInfo.OrganizerEmail = fmt.Sprintf("%s", organizerPayload)

		result := db.Create(&eventInfo)
		if result.Error != nil {
			LogError(ctx, http.StatusInternalServerError, result.Error, 05)
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"message": eventInfo.ID})

	}
}

// List all upcoming events
func DisplayUpcomingEvents(server *api.Server) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var data []models.EventInformation

		db, err := database.ConnectPostgres(server)
		if err != nil {
			LogError(ctx, http.StatusInternalServerError, err, 01)
			return
		}

		db.Find(&data)

		var result = make([]models.DisplayEventInfo, len(data))

		for i := range data {
			result[i].ID = data[i].ID
			result[i].Name = data[i].Name
			result[i].Description = data[i].Description
			result[i].Location = data[i].Location
			result[i].TimeZone = data[i].TimeZone
			result[i].URL = data[i].URL
			result[i].OrganizerEmail = data[i].OrganizerEmail
			result[i].StartTimeString = data[i].StartTime.Format("01-02-2006 Monday")
			result[i].EndTimeString = data[i].EndTime.Format("01-02-2006 Monday")

			if time.Now().Before(data[i].StartTime) {
				result[i].Status = "Upcoming"
			} else if !time.Now().After(data[i].EndTime) {
				result[i].Status = "Open"
			} else {
				result[i].Status = "Ended"
			}
		}

		ctx.JSON(http.StatusOK, gin.H{"result": result})
	}
}
