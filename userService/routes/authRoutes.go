package routes

import (
	"net/http"
	"userService/api"
	"userService/controllers"

	"github.com/gin-gonic/gin"
)

func AuthenticationRoutes(server *api.Server) {
	server.Router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"userService": "success"})
	})
	server.Router.POST("/user/login", controllers.Login(server))
	server.Router.POST("/user/verify", controllers.VerifyOTP(server))
}
