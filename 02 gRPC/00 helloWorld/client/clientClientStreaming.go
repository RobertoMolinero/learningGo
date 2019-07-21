package main

import (
	"../proto"
	"context"
	"fmt"
	"log"
	"time"
)

func main() {
	log.Println("Client for the Hello World Service Server is generated!")

	cc := CreateConnection()
	defer cc.Close()

	client := CreateClient(cc)
	doClientStreaming(client)

	log.Println("Client for the Hello World Service Server is shut down!")
}

func doClientStreaming(client proto.HelloWorldServiceClient) {
	log.Println("Starting to do a client streaming RPC...")

	stream, e := client.HelloWorldLong(context.Background())
	if e != nil {
		log.Fatalf("Error while calling HelloWorldLong: %v", e)
	}

	requests := []*proto.HelloWorldLongRequest{
		{
			Greeting: &proto.Greeting{
				FirstName: "Homer",
				LastName:  "Simpson",
			},
		},
		{
			Greeting: &proto.Greeting{
				FirstName: "Ned",
				LastName:  "Flanders",
			},
		},
		{
			Greeting: &proto.Greeting{
				FirstName: "Lisa",
				LastName:  "Simpson",
			},
		},
	}

	for _, req := range requests {
		log.Printf("Sending request: %v", req)
		stream.Send(req)
		time.Sleep(500 * time.Millisecond)
	}

	response, e := stream.CloseAndRecv()
	if e != nil {
		log.Fatalf("Error while receiving from HelloWorldLong: %v", e)
	}

	fmt.Printf("HelloWorldLong Response: %v\n", response.GetResult())
	log.Println("End of a client streaming RPC...")
}
