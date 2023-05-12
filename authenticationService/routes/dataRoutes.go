package routes

import (
	"authenticationService/api"
	"authenticationService/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DataRoutes(server *api.Server) {
	server.Router.Use(middlewares.Authenticate(server.Config.JWT_SECRET))
	server.Router.GET("/data", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"authorization": "success"})
	})
}