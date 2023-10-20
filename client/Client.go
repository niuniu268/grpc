package main

import (
	"DemogRPC/client/service"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {

	conn, err := grpc.Dial(":8800", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("server has problem: ", err)
	}

	fmt.Println("gRPC is connecting", conn.GetState())

	defer conn.Close()

	UserServiceClient := service.NewUserServiceClient(conn)

	req := &service.UserRequest{UserId: 22}

	resp, err := UserServiceClient.GetUserStock(context.Background(), req)
	if err != nil {
		log.Fatal("gRPC error: ", err)
	}

	fmt.Println("gRPC success，ProdStock = ", resp.ProdStock)

}
