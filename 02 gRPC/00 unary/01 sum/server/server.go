package main

import (
	"../proto"
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
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
