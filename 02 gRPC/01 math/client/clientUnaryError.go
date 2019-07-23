package main

import (
	"../proto"
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

func main() {
	log.Println("Client for the Math Service Server is generated!")

	cc := CreateConnection()
	defer cc.Close()

	client := CreateClient(cc)
	doUnaryError(9, client)
	doUnaryError(-1, client)

	log.Println("Client for the Math Service Server is shut down!")
}

func doUnaryError(number int32, client proto.MathServiceClient) {
	log.Println("Starting to do an unary RPC...")

	request := &proto.SquareRootRequest{
		Number: number,
	}

	log.Printf("Sending request: %v", request)
	response, e := client.SquareRoot(context.Background(), request)
	if e != nil {
		responseError, ok := status.FromError(e)
		if ok {
			log.Println(responseError.Message())
			log.Println(responseError.Code())
			if responseError.Code() == codes.InvalidArgument {
				log.Println("We probably sent a negative number!")
			}
		} else {
			log.Fatalf("Error calling SquareRoot: %v", responseError)
		}
	}

	fmt.Printf("Response from SquareRoot: %v", response.GetRoot())
	log.Println("End of an unary RPC...")
}
