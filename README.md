# gRPC and Protobuf

## Protobuf

```
    $ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
    $ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```

- Update your PATH so that the protoc compiler can find the plugins:
```

    $ export PATH="$PATH:~/go/bin"
```

![protobuf](https://github.com/niuniu268/grpc/blob/master/img/Screenshot%202023-10-20%20at%2008.05.57.png?raw=true)

## gRPC

![gRPC](https://github.com/niuniu268/grpc/blob/master/img/rpc-arch.png?raw=true)

![gRPC](https://github.com/niuniu268/grpc/blob/master/img/Screenshot%202023-10-20%20at%2008.16.05.png?raw=true)

![gRPC](https://github.com/niuniu268/grpc/blob/master/img/Screenshot%202023-10-20%20at%2012.19.55.png?raw=true)

## Token