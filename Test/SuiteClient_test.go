package main

import (
	"DemogRPC/client/service"
	"context"
	"github.com/golang/mock/gomock"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

func TestClientTestSuite(t *testing.T) {
	suite.Run(t, new(ClientTestSuite))
}

func (suite *ClientTestSuite) TestGetUserStock() {
	req := &service.UserRequest{UserId: 22}

	// Set your expectations and behavior for the mocked client method.
	suite.mockUserService.EXPECT().GetUserStock(gomock.Any(), req).Return(&service.UserResponse{ProdStock: 42}, nil)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	stock, err := suite.mockUserService.GetUserStock(ctx, &service.UserRequest{UserId: 22})

	logrus.WithFields(logrus.Fields{
		"expect": 42,
		"real":   stock.String(),
	}).Info("logging")
	suite.Nil(err, "Unexpected error: %v", err)

	suite.Equal(int32(42), stock.ProdStock, "Expect prodstock is 42, but real prodstock is %d", stock.ProdStock)

}
