package main

import (
	"../proto"
	"context"
	"fmt"
	"io"
	"log"
)

func main() {
	log.Println("Client for the Math Service Server is generated!")

	cc := CreateConnection()
	defer cc.Close()

	client := CreateClient(cc)
	doServerStreaming(client)

	log.Println("Client for the Math Service Server is shut down!")
}

func doServerStreaming(client proto.MathServiceClient) {
	log.Println("Starting to do a server streaming RPC...")

	request := &proto.PrimeNumberDecompositionRequest{
		Number: 638514764,
	}

	log.Printf("Sending request: %v", request)
	stream, e := client.PrimeNumberDecomposition(context.Background(), request)
	if e != nil {
		log.Fatalf("Failed to stream: %v", e)
	}

	for {
		msg, e := stream.Recv()
		if e == io.EOF {
			break
		}
		if e != nil {
			log.Fatalf("Error while reading stream: %v", e)
		}
		fmt.Printf("Response from PrimeNumberDecomposition: %v\n", msg.GetDivider())
	}

	log.Println("End of a server streaming RPC...")
}
