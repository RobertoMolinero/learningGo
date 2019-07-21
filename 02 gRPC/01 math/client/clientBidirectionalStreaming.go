package main

import (
	"../proto"
	"context"
	"fmt"
	"io"
	"log"
	"time"
)

func main() {
	log.Println("Client for the Math Service Server is generated!")

	cc := CreateConnection()
	defer cc.Close()

	client := CreateClient(cc)
	doBidirectionalStreaming(client)

	log.Println("Client for the Math Service Server is shut down!")
}

func doBidirectionalStreaming(client proto.MathServiceClient) {
	log.Println("Starting to do a bidirectional streaming RPC...")

	stream, e := client.FindMaximum(context.Background())
	if e != nil {
		log.Fatalf("Error while creating stream: %v", e)
	}

	requests := []*proto.FindMaximumRequest{
		{
			Number: 3,
		},
		{
			Number: 5,
		},
		{
			Number: 27,
		},
		{
			Number: 21,
		},
		{
			Number: 80,
		},
		{
			Number: 56,
		},
		{
			Number: 145,
		},
	}

	waitChannel := make(chan struct{})

	go func() {
		for _, req := range requests {
			log.Printf("Sending request: %v", req)
			stream.Send(req)
			time.Sleep(500 * time.Millisecond)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, e := stream.Recv()
			if e == io.EOF {
				break
			}
			if e != nil {
				log.Fatalf("Error while receiving data: %v", e)
			}
			fmt.Printf("Received: %v\n", res.Maximum)
		}
		close(waitChannel)
	}()

	<-waitChannel
	log.Println("End of the bidirectional streaming RPC...")
}
