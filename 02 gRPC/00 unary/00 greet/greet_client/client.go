package main

import (
	"../greetpb"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
)

func main() {
	fmt.Println("Hello, I'm a client.")

	cc, e := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if e != nil {
		log.Fatalf("Failed to dial: %v", e)
	}

	defer cc.Close()

	client := greetpb.NewGreetServiceClient(cc)
	// log.Printf("Created client: %f", client)

	doUnary(client)
}

func doUnary(client greetpb.GreetServiceClient) {

	fmt.Println("Starting to do a Unary RPC...")

	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Roberto",
			LastName:  "Molinero",
		},
	}

	res, e := client.Greet(context.Background(), req)

	if e != nil {
		log.Fatalf("Failed to call: %v", e)
	}

	log.Printf("Response from Greet: %v", res.Result)
}
