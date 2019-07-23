# gRPC

...

## The Examples

There are two examples which demonstrate each form of communication once.

### Hello World

The first project is a modification of the famous Hello World. After one or more calls by the client, one or more messages are returned.

### Math

In the second project, a mathematical function is offered as an example for each call type.

## Structure

The examples are divided into 3 folders. The folder 'proto' contains the definitions of the services in the form of a protocol buffer file.

> Note:
> Normally I would not check in the generated file.
> Here I did it anyway for illustrative purposes.

The offered functions are located in the folder server in the file server.go.

The client folder contains a file for each individual access type. There is also an additional file that can be used to create the resources required for each access type.

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

## Implemenation of the server

...

## Implementation of the client

...

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