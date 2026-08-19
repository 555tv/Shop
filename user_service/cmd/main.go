package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/555tv/databaze/user_service/internal/server"
	"github.com/555tv/databaze/user_service/proto_contracts"
)

func main() {

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()

	proto_contracts.RegisterUserServiceServer(
		grpcServer,
		&server.UserServer{},
	)

	log.Println("User Service running on :50051")

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatal(err)
	}

}
