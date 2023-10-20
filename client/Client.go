package client

import (
	"DemogRPC/service"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {

	conn, err := grpc.Dial(":8802", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	UserServiceClient := service.NewUserServiceClient(conn)

	resp, err := UserServiceClient.GetUser(context.Background(), &service.UserRequest{Username: "niuniu"})
	if err != nil {
		log.Fatal("gRPC error: ", err)
	}

	fmt.Println("gRPC success，ProdStock = ", resp.Age)

}
