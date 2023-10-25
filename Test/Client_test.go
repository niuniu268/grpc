package main

import (
	"context"
	"testing"
	"time"

	"DemogRPC/client/service"
	"DemogRPC/mocks"
	"github.com/golang/mock/gomock"
)

func TestClient(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserService := mocks.NewMockUserServiceClient(ctrl)

	req := &service.UserRequest{UserId: 22}

	// Set your expectations and behavior for the mocked client method.
	mockUserService.EXPECT().GetUserStock(gomock.Any(), req).Return(&service.UserResponse{ProdStock: 42}, nil)

	testGetUserStock(t, mockUserService)

}

func testGetUserStock(t *testing.T, client service.UserServiceClient) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	stock, err := client.GetUserStock(ctx, &service.UserRequest{
		UserId: 22,
	})
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if stock.ProdStock != 42 {
		t.Errorf("expect prodstock is %d, but real prodstock is %d", 42, stock.ProdStock)
	}

}
