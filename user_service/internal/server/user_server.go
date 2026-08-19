package server

import (
	"context"

	"github.com/555tv/databaze/user_service/proto_contracts"
)

type UserServer struct {
	proto_contracts.UnimplementedUserServiceServer
}

func (s *UserServer) GetUser(
	ctx context.Context,
	req *proto_contracts.GetUserRequest,
) (*proto_contracts.User, error) {

	// Пока просто тестовый пользователь
	return &proto_contracts.User{
		Id:        req.Id,
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john@example.com",
	}, nil
}
