package controllers

import (
	"net/http"
	"time"
	"userService/api"
	"userService/database"
	"userService/models"

	"github.com/gin-gonic/gin"
)

var tokenExpiry time.Duration = 1 * time.Minute // Token Expiry

func Signup(server *api.Server) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var userLoginRequest models.UserLogin

		err1 := ctx.ShouldBindJSON(&userLoginRequest)
		if err1 != nil {
			LogError(ctx, http.StatusBadRequest, err1)
			return
		}

		db := database.CreateClient(&server.Config, 0)
		defer db.Close()

		err2 := db.Set(database.Ctx, userLoginRequest.Email, "Hello", tokenExpiry).Err()
		if err2 != nil {
			LogError(ctx, http.StatusInternalServerError, err2)
			return
		}

		LogMessage(ctx, userLoginRequest)
	}
}
