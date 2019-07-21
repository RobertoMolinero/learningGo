package main

import (
	"../proto"
	"context"
	"fmt"
	"log"
	"time"
)

func main() {
	log.Println("Client for the Math Service Server is generated!")

	cc := CreateConnection()
	defer cc.Close()

	client := CreateClient(cc)
	doClientStreaming(client)

	log.Println("Client for the Math Service Server is shut down!")
}

func doClientStreaming(client proto.MathServiceClient) {
	log.Println("Starting to do a client streaming RPC...")

	stream, e := client.ComputeAverage(context.Background())
	if e != nil {
		log.Fatalf("Error while calling ComputeAverage: %v", e)
	}

	requests := []*proto.ComputeAverageRequest{
		{
			Number: 1,
		},
		{
			Number: 2,
		},
		{
			Number: 3,
		},
		{
			Number: 4,
		},
	}

	for _, req := range requests {
		log.Printf("Sending request: %v", req)
		stream.Send(req)
		time.Sleep(500 * time.Millisecond)
	}

	response, e := stream.CloseAndRecv()
	if e != nil {
		log.Fatalf("Error while receiving from ComputeAverage: %v", e)
	}

	fmt.Printf("ComputeAverage Response: %v\n", response.GetAverage())
	log.Println("End of a client streaming RPC...")
}
