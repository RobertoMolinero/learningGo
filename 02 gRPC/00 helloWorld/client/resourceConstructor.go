package main

import (
	"../proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
)

//CreateConnection Function for establishing the connection to the server.
func CreateConnection() *grpc.ClientConn {
	log.Println("Client will be started next")

	certFile := "../../02 ssl/ssl/ca.crt"
	transportCredentials, e := credentials.NewClientTLSFromFile(certFile, "")
	if e != nil {
		log.Fatalf("Failed loading CA trust certificate: %v", e)
	}

	dialOption := grpc.WithTransportCredentials(transportCredentials)
	cc, e := grpc.Dial("localhost:50051", dialOption)
	if e != nil {
		log.Fatalf("Failed to dial: %v", e)
	}

	log.Println("Client is started")
	return cc
}

//CreateClient Function for constructing a client.
func CreateClient(cc *grpc.ClientConn) proto.HelloWorldServiceClient {
	log.Println("Connection will be established soon")

	client := proto.NewHelloWorldServiceClient(cc)
	log.Printf("Created client: %f", client)

	log.Println("Connection is established")
	return client
}
