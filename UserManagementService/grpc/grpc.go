package grpc

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"userService/pb"
)

func GRPCServe() {
	var rpcConn *grpc.ClientConn
	rpcConn, err := grpc.Dial(":9051", grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err.Error())
		return
	}
	defer rpcConn.Close()
	client := pb.NewAuthorizationClient(rpcConn)

	res, err := client.AuthorizeUser(context.Background(), &pb.UserPayload{
		UserEmail: "rwiteshbera@gmail.com",
	})
	if err != nil {
		log.Fatalln(err.Error())
		return
	}
	fmt.Println(res)
}
