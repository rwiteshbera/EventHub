package routes

import (
	"net/http"
	"userService/api"
	"userService/utils"

	"github.com/gin-gonic/gin"
)

func DataRoutes(server *api.Server) {
	server.Router.Use(utils.Authenticate(server.Config.JWT_SECRET))
	server.Router.GET("/data", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"authorization": "success"})
	})
}
