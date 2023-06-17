package grpc

import (
	"context"
	"eventCatalogService/pb"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Server struct {
	pb.AuthorizationServer
}

func (s *Server) AuthorizeUser(ctx context.Context, req *pb.UserPayload) (*pb.AuthToken, error) {
	fmt.Println(req.GetUserEmail())
	res := &pb.AuthToken{
		UserToken: "abcd",
	}
	return res, nil
}

func GRPCServe() {
	listener, err := net.Listen("tcp", ":9051")
	if err != nil {
		log.Fatalln(err.Error())
		return
	}

	rpcServer := grpc.NewServer()
	pb.RegisterAuthorizationServer(rpcServer, &Server{})

	if err = rpcServer.Serve(listener); err != nil {
		log.Fatalln(err.Error())
		return
	}
}
