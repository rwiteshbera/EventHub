package api

import (
	"net/http"
	"userService/config"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Config config.Config
	Router *gin.Engine
}

func CreateServer() (*Server, error) {
	Config := config.LoadConfig()

	gin.SetMode(gin.ReleaseMode)

	server := &Server{
		Config: *Config,
		Router: gin.Default(),
	}
	server.Router.Use(gin.Recovery())

	server.Router.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"service": "UserManagementService", "status": http.StatusOK})
	})

	return server, nil
}

func (server *Server) Start() error {
	return server.Router.Run(":" + server.Config.SERVER_PORT)
}
