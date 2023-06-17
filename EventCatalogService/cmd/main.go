package main

import (
	"eventCatalogService/api"
	"eventCatalogService/grpc"
	"eventCatalogService/routes"
	"log"
)

func main() {
	server, err := api.CreateServer()
	if err != nil {
		log.Fatalln("unable to create server: ", err.Error())
	}

	routes.Routes(server)

	go grpc.GRPCServe()
	err = server.Start()
	if err != nil {
		log.Fatalln("unable to start the server: ", err.Error())
		return
	}
}
