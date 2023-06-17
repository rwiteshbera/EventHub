package routes

import (
	"eventCatalogService/api"
	"eventCatalogService/controllers"
	"eventCatalogService/middlewares"
)

func Routes(server *api.Server) {
	server.Router.GET("/event", controllers.DisplayUpcomingEvents(server))

	server.Router.Use(middlewares.Authorization())
	server.Router.POST("/event", controllers.CreateEvent(server))
}
