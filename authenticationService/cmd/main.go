package main

import (
	"authenticationService/api"
	"authenticationService/routes"
	"log"
)

func main() {
	server, err := api.CreateServer()
	if err != nil {
		log.Fatalln("unable to create server: ", err.Error())
	}

	routes.AuthenticationRoutes(server)
	routes.DataRoutes(server)

	err = server.Start()
	if err != nil {
		log.Fatalln("unable to start the server: ", err.Error())
	}
}
