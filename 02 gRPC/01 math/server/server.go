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

func (*server) ComputeAverage(stream proto.SumService_ComputeAverageServer) error {
	result := float64(0)
	sum := int32(0)
	counter := int32(0)

	for {
		request, e := stream.Recv()
		if e == io.EOF {
			result = float64(sum / counter)
			response := &proto.AverageResponse{
				Average: result,
			}
			return stream.SendAndClose(response)
		}
		if e != nil {
			log.Fatalf("Error while reading client stream: %v", e)
		}

		sum += request.GetNumber()
		counter++
	}
}

func (*server) PrimeNumberDecomposition(req *proto.PrimeRequest, stream proto.SumService_PrimeNumberDecompositionServer) error {
	n := req.GetNumber()
	k := int32(2)

	for {
		if n%k == 0 {
			res := &proto.PrimeResponse{
				Divider: k,
			}
			stream.Send(res)
			time.Sleep(1000 * time.Millisecond)
			n = n / k
		} else {
			k = k + 1
		}
		if n <= 1 {
			break
		}
	}

	return nil
}

func (*server) Sum(ctx context.Context, req *proto.SumRequest) (*proto.SumResponse, error) {
	x := req.GetS().First
	y := req.GetS().Second
	z := x + y
	response := &proto.SumResponse{
		Result: z,
	}
	return response, nil
}

func main() {
	listener, e := net.Listen("tcp", "0.0.0.0:50051")
	if e != nil {
		log.Fatalf("Failed to listen: %v", e)
	}

	s := grpc.NewServer()
	proto.RegisterSumServiceServer(s, &server{})

	e = s.Serve(listener)
	if e != nil {
		log.Fatalf("Failed to serve: %v", e)
	}
}
