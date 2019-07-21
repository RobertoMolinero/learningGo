package main

import (
	"../proto"
	"context"
	"log"
)

func main() {
	log.Println("Client for the Math Service Server is generated!")

	cc := CreateConnection()
	defer cc.Close()

	client := CreateClient(cc)
	doUnary(client)

	log.Println("Client for the Math Service Server is shut down!")
}

func doUnary(client proto.MathServiceClient) {
	log.Println("Starting to do an unary RPC...")

	request := &proto.SumRequest{
		S: &proto.Summands{
			First:  3,
			Second: 7,
		},
	}

	log.Printf("Sending request: %v", request)
	response, e := client.Sum(context.Background(), request)
	if e != nil {
		log.Fatalf("Failed to call: %v", e)
	}

	log.Printf("Response from Sum: %v", response.Result)
	log.Println("End of an unary RPC...")
}
