package routes

import (
	"eventCatalogService/api"
	"eventCatalogService/controllers"
)

func Routes(server *api.Server) {
	server.Router.POST("/event", controllers.CreateEvent(server))
	server.Router.GET("/event", controllers.DisplayUpcomingEvents(server))
}
