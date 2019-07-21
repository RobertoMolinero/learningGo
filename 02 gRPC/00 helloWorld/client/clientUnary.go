package main

import (
	"../proto"
	"context"
	"log"
)

func main() {
	log.Println("Client for the Hello World Service Server is generated!")

	cc := CreateConnection()
	defer cc.Close()

	client := CreateClient(cc)
	doUnary(client)

	log.Println("Client for the Hello World Service Server is shut down!")
}

func doUnary(client proto.HelloWorldServiceClient) {
	log.Println("Starting to do an unary RPC...")

	request := &proto.HelloWorldRequest{
		Greeting: &proto.Greeting{
			FirstName: "Roberto",
			LastName:  "Molinero",
		},
	}

	log.Printf("Sending request: %v", request)
	response, e := client.HelloWorld(context.Background(), request)
	if e != nil {
		log.Fatalf("Failed to call 'Greet RPC': %v", e)
	}

	log.Printf("Response from HelloWorld: %v", response.Result)
	log.Println("End of an unary RPC...")
}
