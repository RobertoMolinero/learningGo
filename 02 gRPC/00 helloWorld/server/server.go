package main

import (
	"../proto"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
	"io"
	"log"
	"net"
	"strconv"
	"time"
)

type server struct{}

func (*server) HelloWorld(ctx context.Context, request *proto.HelloWorldRequest) (*proto.HelloWorldResponse, error) {
	log.Printf("HelloWorld function was invoked with request %v", request)

	firstName := request.GetGreeting().GetFirstName()
	lastName := request.GetGreeting().GetLastName()

	result := "Hello " + firstName + " " + lastName
	response := &proto.HelloWorldResponse{
		Result: result,
	}

	log.Printf("HelloWorld function sends result %v", response)
	log.Println("HelloWorld function is shut down...")

	return response, nil
}

func (*server) HelloWorldRepeated(request *proto.HelloWorldRepeatedRequest, stream proto.HelloWorldService_HelloWorldRepeatedServer) error {
	log.Printf("HelloWorldRepeated function was invoked with request %v", request)

	firstName := request.Greeting.GetFirstName()
	lastName := request.GetGreeting().GetLastName()
	repeat := request.GetRepeat()

	for i := 0; i < int(repeat); i++ {
		result := "Hello " + firstName + " " + lastName + " number " + strconv.Itoa(i)
		response := &proto.HelloWorldRepeatedResponse{
			Result: result,
		}

		log.Printf("HelloWorldRepeated function sends result %v", response)

		e := stream.Send(response)
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
		request, e := stream.Recv()
		log.Printf("HelloWorldLong received request %v", request)

		if e == io.EOF {
			response := &proto.HelloWorldLongResponse{
				Result: result,
			}

			log.Printf("HelloWorldLong function sends result %v", response)
			log.Println("HelloWorldLong function is shut down...")

			return stream.SendAndClose(response)
		}
		if e != nil {
			log.Fatalf("Error while reading client stream: %v", e)
		}

		firstName := request.GetGreeting().GetFirstName()
		lastName := request.GetGreeting().GetLastName()
		result += "Hello " + firstName + " " + lastName + "!\n"
	}
}

func (*server) HelloWorldEveryone(stream proto.HelloWorldService_HelloWorldEveryoneServer) error {
	log.Printf("HelloWorldEveryone function was invoked with stream %v", stream)

	for {
		request, e := stream.Recv()
		if e == io.EOF {
			log.Println("HelloWorldEveryone function is shut down...")
			return nil
		}
		if e != nil {
			log.Fatalf("Error while reading client stream: %v", e)
			return e
		}

		log.Printf("HelloWorldEveryone received request %v", request)

		firstName := request.GetGreeting().GetFirstName()
		lastName := request.GetGreeting().GetLastName()
		response := "Hello " + firstName + " " + lastName + "!\n"

		log.Printf("HelloWorldEveryone function sends result %v", response)
		e = stream.Send(&proto.HelloWorldEveryoneResponse{
			Result: response,
		})
		if e != nil {
			log.Fatalf("Error while sending data to the client: %v", e)
			return e
		}
	}
}

func (*server) HelloWorldWithDeadline(ctx context.Context, request *proto.HelloWorldWithDeadlineRequest) (*proto.HelloWorldWithDeadlineResponse, error) {
	log.Printf("HelloWorldWithDeadline function was invoked with request %v", request)

	for i := 0; i < 3; i++ {
		if ctx.Err() == context.DeadlineExceeded {
			log.Println("The client canceled the request!")
			return nil, status.Error(codes.Canceled, "The client canceled the request!")
		}

		time.Sleep(1 * time.Second)
	}

	firstName := request.GetGreeting().GetFirstName()
	lastName := request.GetGreeting().GetLastName()

	result := "Hello " + firstName + " " + lastName
	response := &proto.HelloWorldWithDeadlineResponse{
		Result: result,
	}

	log.Printf("HelloWorldWithDeadline function sends result %v", response)
	log.Println("HelloWorldWithDeadline function is shut down...")

	return response, nil
}

func main() {
	log.Println("The Hello World Service Server is starting!")

	listener, e := net.Listen("tcp", "0.0.0.0:50051")
	if e != nil {
		log.Fatalf("Failed to listen: %v", e)
	}

	certFile := "../../02 ssl/ssl/server.crt"
	keyFile := "../../02 ssl/ssl/server.pem"

	transportCredentials, e := credentials.NewServerTLSFromFile(certFile, keyFile)
	if e != nil {
		log.Fatalf("Failed loading certificates: %v", e)
	}

	serverOption := grpc.Creds(transportCredentials)
	s := grpc.NewServer(serverOption)
	proto.RegisterHelloWorldServiceServer(s, &server{})

	if s.Serve(listener); e != nil {
		log.Fatalf("Failed to serve: %v", e)
	}

	log.Println("The Hello World Service Server is shut down!")
}
