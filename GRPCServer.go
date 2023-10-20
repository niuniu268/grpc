package main

import (
	"DemogRPC/service"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	rpcServer := grpc.NewServer()
	service.RegisterUserServiceServer(rpcServer, service.UserService)

	listener, err := net.Listen("tcp", ":8802")

	if err != nil {
		log.Fatal("moniter error: ", err)
	}

	err2 := rpcServer.Serve(listener)

	if err2 != nil {
		log.Fatal("lauch service error", err2)
	}

}
