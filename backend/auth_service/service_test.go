package main

import (
	"context"
	"log"
	"testing"

	"github.com/ttnsgn/grpc-blog-app/global"
	"github.com/ttnsgn/grpc-blog-app/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func Test_authServer_Login(t *testing.T) {
	global.ConnectToTestDB()
	pw, _ := bcrypt.GenerateFromPassword([]byte("example"), bcrypt.DefaultCost)
	global.DB.Collection("user").InsertOne(context.Background(), global.User{ID: primitive.NewObjectID(), Email: "test@gmail.com", Username: "Carl", Password: string(pw)})
	server := authServer{}
	_, err := server.Login(context.Background(), &proto.LogInRequest{Login: "test@gmail.com", Password: "example"})
	if err != nil {
		t.Errorf("1. An error was returned: %v\n", err)
	}
	_, err = server.Login(context.Background(), &proto.LogInRequest{Login: "test@gmail.com", Password: "not-example"})
	if err == nil {
		t.Error("2. Error was nil, expected not nil")
	}
	log.Printf("Error: %v\n", err)
	_, err = server.Login(context.Background(), &proto.LogInRequest{Login: "Carl", Password: "example"})
	if err != nil {
		t.Errorf("3. An error was returned: %v\n", err)
	}
}
