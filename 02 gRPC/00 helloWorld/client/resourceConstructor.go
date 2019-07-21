package main

import (
	"../proto"
	"google.golang.org/grpc"
	"log"
)

//CreateConnection Function for establishing the connection to the server.
func CreateConnection() *grpc.ClientConn {
	cc, e := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if e != nil {
		log.Fatalf("Failed to dial: %v", e)
	}
	return cc
}

//CreateClient Function for constructing a client.
func CreateClient(cc *grpc.ClientConn) proto.HelloWorldServiceClient {
	client := proto.NewHelloWorldServiceClient(cc)
	log.Printf("Created client: %f", client)
	return client
}
