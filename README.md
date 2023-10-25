# gRPC and Protobuf

## Protobuf

```
    $ go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
    $ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

- Update your PATH so that the protoc compiler can find the plugins:
```
    $ export PATH="$PATH:~/go/bin"
```

![protobuf](https://github.com/niuniu268/grpc/blob/master/img/Screenshot%202023-10-20%20at%2008.05.57.png?raw=true)

```
protoc --go_out=../service --go_opt=paths=source_relative --go-grpc_out=../service --go-grpc_opt=paths=source_relative UserRequest.proto
```

## gRPC

- what is gRPC
![gRPC](https://github.com/niuniu268/grpc/blob/master/img/rpc-arch.png?raw=true)

![gRPC](https://github.com/niuniu268/grpc/blob/master/img/Screenshot%202023-10-20%20at%2008.16.05.png?raw=true)

- set up a gRPC server
- set up a gRPC client under a folder, client

![gRPC](https://github.com/niuniu268/grpc/blob/master/img/Screenshot%202023-10-20%20at%2012.19.55.png?raw=true)

## Token

- configurate authinterceptor grpc.UnaryServerInterceptor
- build up an interceptor and check user name and password through Auth method in the GRPCServer
- submit token with user name is "niuniu" and password is "1234" and get an error info

![gRPC](https://github.com/niuniu268/grpc/blob/master/img/Screenshot%202023-10-20%20at%2015.25.42.png?raw=true)

- submit token with username is "admin" and password is "admin" which is correct
![gRPC](https://github.com/niuniu268/grpc/blob/master/img/Screenshot%202023-10-20%20at%2015.26.06.png?raw=true)

## Test

``` 
go install github.com/golang/mock/mockgen@latest
```
```
mockgen -destination=mocks/mock_service.go -package=mocks DemogRPC/client/service UserServiceClient
```
```go test -v ```


