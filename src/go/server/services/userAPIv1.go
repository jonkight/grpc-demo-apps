package services

import (
	"context"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"

	api "github.com/jonkight/grpc-demo-clients/gen/go/demo/user/v1"
)

type userAPIv1 struct{}

func RegisterUserAPIv1(server *grpc.Server) {
	api.RegisterUserAPIServer(server, &userAPIv1{})
}

func (s *userAPIv1) CreateUser(c context.Context, r *api.CreateUserRequest) (*api.CreateUserResponse, error) {
	return &api.CreateUserResponse{}, nil
}

func (s *userAPIv1) GetUser(c context.Context, r *api.GetUserRequest) (*api.GetUserResponse, error) {
	timestamp := timestamppb.Timestamp{Seconds: time.Now().Unix()}
	user := api.User{UserId: r.GetUserId(), CreatedTimestamp: &timestamp}
	return &api.GetUserResponse{User: &user}, nil
}
