package client

import "context"
import "DemogRPC/client/service"

// UserServiceClientInterface is an interface for the gRPC client.
type UserServiceClientInterface interface {
	GetUserStock(ctx context.Context, in *service.UserRequest) (*service.UserResponse, error)
}
