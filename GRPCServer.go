package main

import (
	"DemogRPC/service"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {

	listener, err := net.Listen("tcp", ":8800")

	if err != nil {
		log.Fatal("monitor error: ", err)
	}

	rpcServer := grpc.NewServer()
	service.RegisterUserServiceServer(rpcServer, &service.UserService)

	err2 := rpcServer.Serve(listener)

	if err2 != nil {
		log.Fatal("launch service error", err2)
	}

}
