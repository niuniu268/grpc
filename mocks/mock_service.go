// Code generated by MockGen. DO NOT EDIT.
// Source: DemogRPC/client/service (interfaces: UserServiceClient)

// Package mocks is a generated GoMock package.
package mocks

import (
	service "DemogRPC/client/service"
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
)

// MockUserServiceClient is a mock of UserServiceClient interface.
type MockUserServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockUserServiceClientMockRecorder
}

// MockUserServiceClientMockRecorder is the mock recorder for MockUserServiceClient.
type MockUserServiceClientMockRecorder struct {
	mock *MockUserServiceClient
}

// NewMockUserServiceClient creates a new mock instance.
func NewMockUserServiceClient(ctrl *gomock.Controller) *MockUserServiceClient {
	mock := &MockUserServiceClient{ctrl: ctrl}
	mock.recorder = &MockUserServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserServiceClient) EXPECT() *MockUserServiceClientMockRecorder {
	return m.recorder
}

// GetUserStock mocks base method.
func (m *MockUserServiceClient) GetUserStock(arg0 context.Context, arg1 *service.UserRequest, arg2 ...grpc.CallOption) (*service.UserResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetUserStock", varargs...)
	ret0, _ := ret[0].(*service.UserResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserStock indicates an expected call of GetUserStock.
func (mr *MockUserServiceClientMockRecorder) GetUserStock(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserStock", reflect.TypeOf((*MockUserServiceClient)(nil).GetUserStock), varargs...)
}
