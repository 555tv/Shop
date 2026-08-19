package clients

import (
	"context"

	"github.com/555tv/databaze/user_service/proto_contracts"
	"google.golang.org/grpc"
)

type UserClient struct {
	client proto_contracts.UserServiceClient
}

func NewUserClient(conn *grpc.ClientConn) *UserClient {
	return &UserClient{
		client: proto_contracts.NewUserServiceClient(conn),
	}
}

func (c *UserClient) GetUser(
	ctx context.Context,
	id string,
) (*proto_contracts.User, error) {

	return c.client.GetUser(
		ctx,
		&proto_contracts.GetUserRequest{
			Id: id,
		},
	)
}
