package main

import (
	"../proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
	"time"
)

func main() {
	cc, e := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if e != nil {
		log.Fatalf("Failed to dial: %v", e)
	}

	defer cc.Close()

	client := proto.NewSumServiceClient(cc)

	// doUnary(client)
	// doServerStreaming(client)
	doClientStreaming(client)
}

func doUnary(client proto.SumServiceClient) {
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

func doServerStreaming(client proto.SumServiceClient) {
	request := &proto.PrimeRequest{
		Number: 638514764,
	}

	resStream, e := client.PrimeNumberDecomposition(context.Background(), request)
	if e != nil {
		log.Fatalf("Failed to stream: %v", e)
	}

	for {
		msg, e := resStream.Recv()
		if e == io.EOF {
			break
		}
		if e != nil {
			log.Fatalf("Error while reading stream: %v", e)
		}
		log.Printf("Response from PrimeNumberDecomposition: %v", msg.GetDivider())
	}
}

func doClientStreaming(client proto.SumServiceClient) {
	stream, e := client.ComputeAverage(context.Background())
	if e != nil {
		log.Fatalf("Error while calling ComputeAverage: %v", e)
	}

	requests := []*proto.AverageRequest{
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
		stream.Send(req)
		time.Sleep(500 * time.Millisecond)
	}

	response, e := stream.CloseAndRecv()
	if e != nil {
		log.Fatalf("Error while receiving from ComputeAverage: %v", e)
	}

	fmt.Printf("ComputeAverage Response: %v\n", response.GetAverage())
}
