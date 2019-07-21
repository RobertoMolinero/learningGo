package main

import (
	"../proto"
	"context"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	"strconv"
	"time"
)

type server struct{}

func (*server) HelloWorld(ctx context.Context, req *proto.HelloWorldRequest) (*proto.HelloWorldResponse, error) {
	log.Printf("HelloWorld function was invoked with request %v", req)

	firstName := req.GetGreeting().GetFirstName()
	lastName := req.GetGreeting().GetLastName()

	result := "Hello " + firstName + " " + lastName
	res := &proto.HelloWorldResponse{
		Result: result,
	}

	log.Printf("HelloWorld function sends result %v", res)
	log.Println("HelloWorld function is shut down...")

	return res, nil
}

func (*server) HelloWorldRepeated(req *proto.HelloWorldRepeatedRequest, stream proto.HelloWorldService_HelloWorldRepeatedServer) error {
	log.Printf("HelloWorldRepeated function was invoked with request %v", req)

	firstName := req.Greeting.GetFirstName()
	lastName := req.GetGreeting().GetLastName()
	repeat := req.GetRepeat()

	for i := 0; i < int(repeat); i++ {
		result := "Hello " + firstName + " " + lastName + " number " + strconv.Itoa(i)
		res := &proto.HelloWorldRepeatedResponse{
			Result: result,
		}

		log.Printf("HelloWorldRepeated function sends result %v", res)

		e := stream.Send(res)
		if e != nil {
			log.Fatalf("Error while sending data to client: %v", e)
		}

		time.Sleep(500 * time.Millisecond)
	}

	log.Println("HelloWorldRepeated function is shut down...")

	return nil
}

func (*server) HelloWorldLong(stream proto.HelloWorldService_HelloWorldLongServer) error {
	log.Printf("HelloWorldLong function was invoked with stream %v", stream)
	result := ""

	for {
		req, e := stream.Recv()
		log.Printf("HelloWorldLong received request %v", req)

		if e == io.EOF {
			res := &proto.HelloWorldLongResponse{
				Result: result,
			}

			log.Printf("HelloWorldLong function sends result %v", res)
			log.Println("HelloWorldLong function is shut down...")
			return stream.SendAndClose(res)
		}
		if e != nil {
			log.Fatalf("Error while reading client stream: %v", e)
		}

		firstName := req.GetGreeting().GetFirstName()
		lastName := req.GetGreeting().GetLastName()
		result += "Hello " + firstName + " " + lastName + "!\n"
	}
}

func (*server) HelloWorldEveryone(stream proto.HelloWorldService_HelloWorldEveryoneServer) error {
	log.Printf("HelloWorldEveryone function was invoked with stream %v", stream)
	for {
		req, e := stream.Recv()
		if e == io.EOF {
			log.Println("HelloWorldEveryone function is shut down...")
			return nil
		}
		if e != nil {
			log.Fatalf("Error while reading client stream: %v", e)
			return e
		}

		log.Printf("HelloWorldLong received request %v", req)

		firstName := req.GetGreeting().GetFirstName()
		lastName := req.GetGreeting().GetLastName()
		res := "Hello " + firstName + " " + lastName + "!\n"

		log.Printf("HelloWorldEveryone function sends result %v", res)
		e = stream.Send(&proto.HelloWorldEveryoneResponse{
			Result: res,
		})
		if e != nil {
			log.Fatalf("Error while sending data to the client: %v", e)
			return e
		}
	}
}

func main() {
	log.Println("The Hello World Service Server is starting!")

	listener, e := net.Listen("tcp", "0.0.0.0:50051")
	if e != nil {
		log.Fatalf("Failed to listen: %v", e)
	}

	s := grpc.NewServer()
	proto.RegisterHelloWorldServiceServer(s, &server{})

	if s.Serve(listener); e != nil {
		log.Fatalf("Failed to serve: %v", e)
	}

	log.Println("The Hello World Service Server is shut down!")
}
