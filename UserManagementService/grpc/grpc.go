package grpc

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/grpc"
	"log"
	"net"
	api2 "userService/api"
	"userService/database"
	"userService/pb"
	"userService/utils"
)

type AuthServer struct {
	jwt       string
	mongo_uri string
	pb.AuthorizationServer
}

func (s *AuthServer) AuthorizeUser(ctx context.Context, req *pb.AuthToken) (*pb.UserPayload, error) {
	fmt.Println(req.GetUserToken())
	// Read the auth token sent by Event Catalog Service
	authToken := req.GetUserToken()

	// Validate the token
	claims, err := utils.ValidateToken(authToken, s.jwt)
	if err != nil {
		return nil, errors.New("invalid token")
	}

	// Check the user existence in MONGODB
	// Create MongoDB Instance
	mongoClient, err2 := database.CreateMongoInstance(s.mongo_uri)
	if err2 != nil {
		return nil, err2
	}
	userCollection := database.OpenMongoCollection(mongoClient, "user")
	filter := bson.D{{Key: "email", Value: claims.Email}}
	count, err := userCollection.CountDocuments(ctx, filter)
	if err != nil {
		return nil, err
	}
	if count == 0 {
		return nil, errors.New("no user found")
	}

	// Send the payload data to Event Catalog Service
	res := &pb.UserPayload{
		UserEmail: claims.Email,
	}
	return res, nil
}

func GRPCListen(serverConfig *api2.Server) {
	listener, err := net.Listen("tcp", ":9051")
	if err != nil {
		log.Fatalln(err.Error())
		return
	}

	rpcServer := grpc.NewServer()
	pb.RegisterAuthorizationServer(rpcServer, &AuthServer{jwt: serverConfig.Config.JWT_SECRET, mongo_uri: serverConfig.Config.MONGO_DB_URI})

	if err = rpcServer.Serve(listener); err != nil {
		log.Fatalln(err.Error())
		return
	}
}
