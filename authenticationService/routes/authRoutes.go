package routes

import (
	"authenticationService/api"
	"authenticationService/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthenticationRoutes(server *api.Server) {
	server.Router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"authenticationService": "success"})
	})
	server.Router.POST("/user/login", controllers.Login(server))
	server.Router.POST("/user/verify", controllers.VerifyOTP(server))
}
