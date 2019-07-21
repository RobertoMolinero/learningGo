package main

import (
	"../proto"
	"context"
	"fmt"
	"io"
	"log"
)

func main() {
	log.Println("Client for the Hello World Service Server is generated!")

	cc := CreateConnection()
	defer cc.Close()

	client := CreateClient(cc)
	doServerStreaming(client)

	log.Println("Client for the Hello World Service Server is shut down!")
}

func doServerStreaming(client proto.HelloWorldServiceClient) {
	log.Println("Starting to do a server streaming RPC...")

	request := &proto.HelloWorldRepeatedRequest{
		Greeting: &proto.Greeting{
			FirstName: "Roberto",
			LastName:  "Molinero",
		},
		Repeat: 8,
	}

	log.Printf("Sending request: %v", request)
	stream, e := client.HelloWorldRepeated(context.Background(), request)
	if e != nil {
		log.Fatalf("Error while calling 'GreetManyTimes RPC': %v", e)
	}

	for {
		msg, e := stream.Recv()
		if e == io.EOF {
			break
		}
		if e != nil {
			log.Fatalf("Error while reading stream: %v", e)
		}
		fmt.Printf("Response from GreetManyTimes: %v\n", msg.GetResult())
	}

	log.Println("End of a server streaming RPC...")
}