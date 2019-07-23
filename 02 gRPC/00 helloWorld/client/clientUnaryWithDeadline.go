package main

import (
	"../proto"
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"time"
)

func main() {
	log.Println("Client for the Hello World Service Server is generated!")

	cc := CreateConnection()
	defer cc.Close()

	client := CreateClient(cc)
	doUnaryWithDeadline(client, 5*time.Second)
	doUnaryWithDeadline(client, 1*time.Second)

	log.Println("Client for the Hello World Service Server is shut down!")
}

func doUnaryWithDeadline(client proto.HelloWorldServiceClient, timeout time.Duration) {
	log.Println("Starting to do an unary RPC...")

	request := &proto.HelloWorldWithDeadlineRequest{
		Greeting: &proto.Greeting{
			FirstName: "Roberto",
			LastName:  "Molinero",
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	log.Printf("Sending request: %v", request)
	response, e := client.HelloWorldWithDeadline(ctx, request)

	if e != nil {
		statusError, ok := status.FromError(e)
		if ok {
			if statusError.Code() == codes.DeadlineExceeded {
				log.Println("Timeout was hit! Deadline was exceeded!")
			} else {
				log.Printf("Unexpected error: %v", statusError)
			}
		} else {
			log.Fatalf("Failed to call 'Greet RPC': %v", statusError)
		}
	}

	fmt.Printf("Response from HelloWorldWithDeadline: %v\n", response.GetResult())
	log.Println("End of an unary RPC...")
}
