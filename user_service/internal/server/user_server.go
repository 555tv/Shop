package server

import (
	"context"
	"github.com/555tv/databaze/user_service/proto_contracts"
	"log"
)

type UserServer struct {
	proto_contracts.UnimplementedUserServiceServer
}

func (s *UserServer) GetUser(
	ctx context.Context,
	req *proto_contracts.GetUserRequest,
) (*proto_contracts.User, error) {

	log.Println("GetUser called via gRPC, ID:", req.Id)

	return &proto_contracts.User{
		Id:        req.Id,
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john@example.com",
	}, nil
}
