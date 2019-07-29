# gRPC

gRPC (gRPC Remote Procedure Calls) is an open source remote procedure call (RPC) system initially developed at Google. It uses HTTP/2 for transport, Protocol Buffers as the interface description language, and provides features such as authentication, bidirectional streaming and flow control, blocking or nonblocking bindings, and cancellation and timeouts. It generates cross-platform client and server bindings for many languages. Most common usage scenarios include connecting services in microservices style architecture and connect mobile devices, browser clients to backend services.

[Source Wikipedia](https://en.wikipedia.org/wiki/GRPC)

## The Examples

### Hello World

The first project is a modification of the famous Hello World. After one or more calls by the client, one or more messages are returned.

### Math

In the second project, a mathematical function is offered as an example for each call type.

### Evans

The tool is used to control an implemented service via a command line.

Thus it is possible to test the server even if no client software is available yet.

It is possible to display the provided packages and to navigate between them. The existing methods and their parameters can be read. The methods can be called with and their results can be displayed.

To use the most common commands, there is a separate README.md in the subfolder.

### SSL

When communicating over a public network, interception of the sent data cannot be prevented. Therefore the encryption of this data is indispensable.

The library provides general cryptographic functions for encryption and decryption. Certificates are generated to control who is authorized to access the data. 

These certificates are then assigned to the program parts.

### Blog

In the last example the functions for a CRUD application are implemented. A MongoDb instance serves as data storage.

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

...

```
protoc *.proto --go_out=plugins=grpc:.
```

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