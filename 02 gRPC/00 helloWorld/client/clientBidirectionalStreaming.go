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
	log.Println("Client for the Hello World Service Server is generated!")

	cc := CreateConnection()
	defer cc.Close()

	client := CreateClient(cc)
	doBidirectionalStreaming(client)

	log.Println("Client for the Hello World Service Server is shut down!")
}

func doBidirectionalStreaming(client proto.HelloWorldServiceClient) {
	log.Println("Starting to do a bidirectional streaming RPC...")

	stream, e := client.HelloWorldEveryone(context.Background())
	if e != nil {
		log.Fatalf("Error while creating stream: %v", e)
	}

	requests := []*proto.HelloWorldEveryoneRequest{
		{
			Greeting: &proto.Greeting{
				FirstName: "Dennis",
				LastName:  "Ritchie",
			},
		},
		{
			Greeting: &proto.Greeting{
				FirstName: "Ken",
				LastName:  "Thompson",
			},
		},
		{
			Greeting: &proto.Greeting{
				FirstName: "Edgar F.",
				LastName:  "Codd",
			},
		},
	}

	waitChannel := make(chan struct{})

	go func() {
		for _, request := range requests {
			log.Printf("Sending request: %v", request)
			stream.Send(request)
			time.Sleep(500 * time.Millisecond)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			response, e := stream.Recv()
			if e == io.EOF {
				break
			}
			if e != nil {
				log.Fatalf("Error while receiving data: %v", e)
			}
			fmt.Printf("Received: %v\n", response.GetResult())
		}
		close(waitChannel)
	}()

	<-waitChannel
	log.Println("End of the bidirectional streaming RPC...")
}
