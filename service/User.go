package service

import (
	"context"
)

var UserService = &userService{}

type userService struct {
}

func (u *userService) mustEmbedUnimplementedUserServiceServer() {
	//TODO implement me
	panic("implement me")
}

func (u *userService) GetUser(ctx context.Context, request *UserRequest) (*UserResponse, error) {
	//TODO implement me
	return &UserResponse{
		Username: "niuniu",
	}, nil
}
