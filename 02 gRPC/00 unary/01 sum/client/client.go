package main

import (
	"../proto"
	"context"
	"google.golang.org/grpc"
	"log"
)

func main() {

	cc, e := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if e != nil {
		log.Fatalf("Failed to dial: %v", e)
	}

	defer cc.Close()

	client := proto.NewSumServiceClient(cc)

	request := &proto.SumRequest{
		S: &proto.Summands{
			First:  3,
			Second: 7,
		},
	}

	response, e := client.Sum(context.Background(), request)

	if e != nil {
		log.Fatalf("Failed to call: %v", e)
	}

	log.Printf("Response from Sum: %v", response.Result)
}
