package main

import (
	"context"
	"log"
	"net"
	"time"

	"github.com/ttnsgn/grpc-blog-app/global"
	"github.com/ttnsgn/grpc-blog-app/proto"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type authServer struct {
	proto.AuthServiceServer
}

func (*authServer) Login(_ context.Context, req *proto.LogInRequest) (*proto.AuthResponse, error) {
	log.Println("Login was invoked")
	login, password := req.GetLogin(), req.GetPassword()
	ctx, cancel := global.NewDBContext(5 * time.Second)
	defer cancel()
	var user global.User
	global.DB.Collection("user").FindOne(ctx, bson.M{"$or": []bson.M{{"username": login}, {"email": login}}}).Decode(&user)
	if user == global.NilUser {
		return &proto.AuthResponse{}, status.Error(codes.NotFound, "Wrong Login Credentials provided - Username/Email")
	}
	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
		return &proto.AuthResponse{}, status.Error(codes.Unauthenticated, "Wrong Login Credentials provided - Password")
	}
	return &proto.AuthResponse{Token: user.GetToken()}, nil
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
