package main

import (
	"log"
	"userService/api"
	"userService/grpc"
	"userService/routes"
)

func main() {
	server, err := api.CreateServer()
	if err != nil {
		log.Fatalln("unable to create server: ", err.Error())
	}

	routes.AuthenticationRoutes(server)

	go grpc.GRPCListen(server.Config.JWT_SECRET)
	err = server.Start()
	if err != nil {
		log.Fatalln("unable to start the server: ", err.Error())
	}
}
