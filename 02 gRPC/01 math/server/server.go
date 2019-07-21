package main

import (
	"../proto"
	"context"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	"time"
)

type server struct {
}

func (*server) Sum(ctx context.Context, req *proto.SumRequest) (*proto.SumResponse, error) {
	log.Printf("Sum function was invoked with request %v", req)

	x := req.GetS().First
	y := req.GetS().Second
	z := x + y

	res := &proto.SumResponse{
		Result: z,
	}

	log.Printf("Sum function sends result %v", res)
	log.Println("Sum function is shut down...")

	return res, nil
}

func (*server) PrimeNumberDecomposition(req *proto.PrimeNumberDecompositionRequest, stream proto.MathService_PrimeNumberDecompositionServer) error {
	log.Printf("PrimeNumberDecomposition function was invoked with request %v", req)

	n := req.GetNumber()
	k := int32(2)

	for {
		if n%k == 0 {
			res := &proto.PrimeNumberDecompositionResponse{
				Divider: k,
			}

			log.Printf("PrimeNumberDecomposition function sends result %v", res)
			stream.Send(res)
			time.Sleep(500 * time.Millisecond)
			n = n / k
		} else {
			k = k + 1
		}
		if n <= 1 {
			break
		}
	}

	log.Println("PrimeNumberDecomposition function is shut down...")

	return nil
}

func (*server) ComputeAverage(stream proto.MathService_ComputeAverageServer) error {
	log.Printf("ComputeAverage function was invoked with stream %v", stream)

	result := float64(0)
	sum := int32(0)
	counter := int32(0)

	for {
		req, e := stream.Recv()
		log.Printf("ComputeAverage received request %v", req)

		if e == io.EOF {
			result = float64(sum) / float64(counter)
			res := &proto.ComputeAverageResponse{
				Average: result,
			}

			log.Printf("HelloWorldLong function sends result %v", res)
			log.Println("HelloWorldLong function is shut down...")

			return stream.SendAndClose(res)
		}
		if e != nil {
			log.Fatalf("Error while reading client stream: %v", e)
		}

		sum += req.GetNumber()
		counter++
	}
}

func (*server) FindMaximum(stream proto.MathService_FindMaximumServer) error {
	log.Printf("FindMaximum function was invoked with stream %v", stream)

	result := int32(0)

	for {
		req, e := stream.Recv()
		if e == io.EOF {
			log.Println("FindMaximum function is shut down...")
			return nil
		}
		if e != nil {
			log.Fatalf("Error while reading client stream: %v", e)
			return e
		}

		log.Printf("FindMaximum received request %v", req)
		number := req.GetNumber()

		if number > result {
			result = number
		}

		res := &proto.FindMaximumResponse{
			Maximum: result,
		}

		log.Printf("FindMaximum function sends result %v", res)
		e = stream.Send(res)
		if e != nil {
			log.Fatalf("Error while sending data to the client: %v", e)
			return e
		}
	}
}

func main() {
	log.Println("The Math Service Server is starting!")

	listener, e := net.Listen("tcp", "0.0.0.0:50051")
	if e != nil {
		log.Fatalf("Failed to listen: %v", e)
	}

	s := grpc.NewServer()
	proto.RegisterMathServiceServer(s, &server{})

	e = s.Serve(listener)
	if e != nil {
		log.Fatalf("Failed to serve: %v", e)
	}

	log.Println("The Math Service Server is shut down!")
}
