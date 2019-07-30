package main

import (
	"../proto"
	"context"
	"fmt"
	"io"
	"log"
)

func main() {
	log.Println("Client for the Blog Service Server is generated!")

	cc := CreateConnection()
	defer cc.Close()

	client := CreateClient(cc)
	listBlog(client)

	log.Println("Client for the Blog Service Server is shut down!")
}

func listBlog(client proto.BlogServiceClient) {
	log.Println("ListBlog called")

	listRequest := proto.ListBlogRequest{}
	stream, e := client.ListBlog(context.Background(), &listRequest)
	if e != nil {
		log.Fatalf("Failed to call: %v", e)
	}

	for {
		listResponse, e := stream.Recv()
		if e == io.EOF {
			break
		}
		if e != nil {
			log.Fatalf("Failed to receive data from stream: %v", e)
		}
		streamItem := listResponse.GetBlog()
		fmt.Printf("Stream: Blog Id: %s | Author: %s | Title: %s | Content: %s\n", streamItem.GetId(), streamItem.GetAuthorId(), streamItem.GetTitle(), streamItem.GetContent())
	}

	log.Println("ListBlog finished")
}
