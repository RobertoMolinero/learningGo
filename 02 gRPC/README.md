# gRPC

## Writing the protocol buffer file

The first part defines the syntax used and the surrounding package.

```
syntax = "proto3";

package math;
option go_package = "proto";
```

Input and output values are defined for each function. The function itself is defined within a service structure.

```
message RequestStruct {
    <Type> parameter 1 = 1;
    <Type> parameter 2 = 2;
    ...
}

message ResponseStruct {
     <Type> parameter 1 = 1;
     <Type> parameter 2 = 2;
     ...
}

service MathService {
    rpc NameOfTheFunction (RequestStruct) returns (ResponseStruct) {
    };
}
```

## Generation

protoc *.proto --go_out=plugins=grpc:.

## Execution

Simply start the server and the client in 2 separate terminals.

### Server

```
cd server
go run server.go
```

### Client

```
cd client
go run clientUnary.go resourceConstructor.go
```