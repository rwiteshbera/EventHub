package routes

import (
	"eventCatalogService/api"
	"eventCatalogService/controllers"
)

func Routes(server *api.Server) {
	server.Router.POST("/event", controllers.CreateEvent(server))
}
