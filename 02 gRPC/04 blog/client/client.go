package main

import (
	"../proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
)

func main() {
	log.Println("Client for the Hello World Service Server is generated!")

	cc := CreateConnection()
	defer cc.Close()

	client := CreateClient(cc)

	createRequest := proto.CreateBlogRequest{
		Blog: &proto.Blog{
			AuthorId: "Rob",
			Title:    "Hey Ho, Let's Go!",
			Content:  "Lorem ipsum dolor sit amet, consetetur sadipscing elitr, sed diam nonumy eirmod tempor invidunt ut labore et dolore magna aliquyam erat, sed diam voluptua. At vero eos et accusam et justo duo dolores et ea rebum. Stet clita kasd gubergren, no sea tak",
		},
	}

	createResponse, e := client.CreateBlog(context.Background(), &createRequest)
	if e != nil {
		log.Fatalf("Failed to call: %v", e)
	}

	log.Printf("Blog has been created: %v", createResponse)

	readRequest := proto.ReadBlogRequest{
		BlogId: createResponse.GetBlog().GetId(),
	}

	readResponse, e := client.ReadBlog(context.Background(), &readRequest)
	if e != nil {
		log.Fatalf("Failed to call: %v", e)
	}

	blog := readResponse.GetBlog()
	fmt.Printf("Blog Id: %s | Author: %s | Title: %s | Content: %s\n", blog.GetId(), blog.GetAuthorId(), blog.GetTitle(), blog.GetContent())

	updateRequest := proto.UpdateBlogRequest{
		Blog: &proto.Blog{
			Id:       blog.GetId(),
			AuthorId: "Rob (2)",
			Title:    "Hey Ho, Let's Go! (2)",
			Content:  "Lorem ipsum dolor sit... (2)",
		},
	}

	updateResponse, e := client.UpdateBlog(context.Background(), &updateRequest)
	if e != nil {
		log.Fatalf("Failed to call: %v", e)
	}

	updated := updateResponse.GetBlog()
	fmt.Printf("Blog Id: %s | Author: %s | Title: %s | Content: %s\n", updated.GetId(), updated.GetAuthorId(), updated.GetTitle(), updated.GetContent())

	deleteRequest := proto.DeleteBlogRequest{
		BlogId: blog.GetId(),
	}

	deleteResponse, e := client.DeleteBlog(context.Background(), &deleteRequest)
	if e != nil {
		log.Fatalf("Failed to call: %v", e)
	}

	fmt.Printf("Blog Id: %s\n", deleteResponse.GetBlogId())

	log.Println("Client for the Hello World Service Server is shut down!")
}

//CreateConnection Function for establishing the connection to the server.
func CreateConnection() *grpc.ClientConn {
	log.Println("Client will be started next")

	dialOption := grpc.WithInsecure()
	cc, e := grpc.Dial("localhost:50051", dialOption)
	if e != nil {
		log.Fatalf("Failed to dial: %v", e)
	}

	log.Println("Client is started")
	return cc
}

//CreateClient Function for constructing a client.
func CreateClient(cc *grpc.ClientConn) proto.BlogServiceClient {
	log.Println("Connection will be established soon")

	client := proto.NewBlogServiceClient(cc)
	log.Printf("Created client: %f", client)

	log.Println("Connection is established")
	return client
}
