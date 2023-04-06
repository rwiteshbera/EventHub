package routes

import (
	"userService/api"
	"userService/controllers"
)

func AuthenticationRoutes(server *api.Server) {
	server.Router.POST("/user/login", controllers.Signup(server))
}
