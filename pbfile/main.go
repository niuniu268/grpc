package main

//
//import (
//	"DemogRPC/service"
//	"fmt"
//	"github.com/golang/protobuf/proto"
//)
//
//func main() {
//
//	s := &service.User{
//		Username: "niuniu",
//		Age:      12,
//	}
//
//	b, err := proto.Marshal(s)
//	if err != nil {
//		fmt.Println("error info:", err)
//	}
//
//	fmt.Println(b)
//
//	newUser := &service.User{}
//	err1 := proto.Unmarshal(b, newUser)
//	if err1 != nil {
//		fmt.Println("unmarshal error is ", err1)
//	}
//
//}
