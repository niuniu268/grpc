package main

import (
	"DemogRPC/mocks"
	"github.com/golang/mock/gomock"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"
	"os"
)

type ClientTestSuite struct {
	suite.Suite
	ctrl            *gomock.Controller
	mockUserService *mocks.MockUserServiceClient
}

func (suite *ClientTestSuite) SetupSuite() {
	suite.ctrl = gomock.NewController(suite.T())
	suite.mockUserService = mocks.NewMockUserServiceClient(suite.ctrl)

	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stdout)
	logrus.SetFormatter(&logrus.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})
}

func (suite *ClientTestSuite) TearDownSuite() {
	suite.ctrl.Finish()
}
