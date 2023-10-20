package main

import (
	"DemogRPC/service"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"log"
	"net"
)

func main() {

	var authInterceptor grpc.UnaryServerInterceptor

	authInterceptor = func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		//	intercept requests and check Token
		err = Auth(ctx)
		if err != nil {
			return
		}

		//	continue handling
		return handler(ctx, req)
	}

	listener, err := net.Listen("tcp", ":8800")

	if err != nil {
		log.Fatal("monitor error: ", err)
	}

	rpcServer := grpc.NewServer(grpc.UnaryInterceptor(authInterceptor))
	service.RegisterUserServiceServer(rpcServer, &service.UserService)

	err2 := rpcServer.Serve(listener)

	if err2 != nil {
		log.Fatal("launch service error", err2)
	}
	fmt.Println("the gRPC server is launching")

}

func Auth(ctx context.Context) error {

	md, ok := metadata.FromIncomingContext(ctx)

	if !ok {
		return fmt.Errorf("missing credentials")
	}

	var user string
	var password string

	if val, ok := md["user"]; ok {
		user = val[0]
	}
	if val, ok := md["password"]; ok {
		password = val[0]
	}
	if user != "admin" || password != "admin" {
		return status.Errorf(codes.Unauthenticated, "token is illegal")
	}
	return nil

}
