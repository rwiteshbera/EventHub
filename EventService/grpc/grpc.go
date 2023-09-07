package grpc

import (
	"context"
	"errors"
	"eventCatalogService/pb"
	"google.golang.org/grpc"
)

func GRPCServe(token string) (*pb.UserPayload, error) {
	var rpcConn *grpc.ClientConn
	rpcConn, err := grpc.Dial(":9051", grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	defer rpcConn.Close()
	client := pb.NewAuthorizationClient(rpcConn)

	res, err := client.AuthorizeUser(context.Background(), &pb.AuthToken{
		UserToken: token,
	})
	if err != nil {
		return nil, errors.New("authorization failed")
	}
	return res, nil
}
