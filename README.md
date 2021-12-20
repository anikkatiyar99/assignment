# Assignment

## Structure
```
service.go # service interface definition
 \
  - core/
    service.go # implementation of the service interface that DOES NOT use gRPC.
 \
  - grpc/
    service.proto # protobuf definition file for the gRPC service
    service.pb.go # compiled protobufs
     \
      - client
        client.go # gRPC client wrapper that implements the top-level service interface
     \  
      - server
        controller.go # controller that implements the gRPC service interface in service.pb.go
        main.go       # gRPC server executable that exposes the controller
```

## Run the Example

cd mysvctest
export GRPC_ADDR=":9000"
go build -o ./mysvctest .
# sample commands:
./mysvctest 1
./mysvctest 1 2 3
./mysvctest 5
./mysvctest 1 2 5
```


