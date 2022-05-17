package main

import (
	"context"
	"log"
	"net"

	"github.com/ttnsgn/grpc-blog-app/proto"
	"google.golang.org/grpc"
)

type authServer struct {
	proto.AuthServiceServer
}

func (*authServer) Login(ctx context.Context, req *proto.LogInRequest) (*proto.AuthResponse, error) {
	log.Println("Login was invoked")
	return &proto.AuthResponse{}, nil
}

func main() {
	server := grpc.NewServer()
	proto.RegisterAuthServiceServer(server, &authServer{})
	listener, err := net.Listen("tcp", ":5000")
	if err != nil {
		log.Fatalf("Error creating listener: %v\n", err)
	}
	if err := server.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}
