package service

import (
	"context"
	"fmt"
)

var UserService userService

type userService struct {
	UnimplementedUserServiceServer
}

func (u *userService) GetUserStock(ctx context.Context, req *UserRequest) (*UserResponse, error) {

	info := u.GetInfoById(req.UserId)

	fmt.Println(req.UserId)
	fmt.Println(info)

	return &UserResponse{ProdStock: info}, nil
}
func (u *userService) mustEmbedUnimplementedUserServiceServer() {}

func (u *userService) GetInfoById(id int32) int32 {

	return id

}
