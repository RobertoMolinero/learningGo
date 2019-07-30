package main

import (
	"../proto"
	"context"
	"fmt"
	"log"
)

func main() {
	log.Println("Client for the Blog Service Server is generated!")

	cc := CreateConnection()
	defer cc.Close()

	client := CreateClient(cc)
	blogId := CreateBlog(client)
	readBlog(client, blogId)

	log.Println("Client for the Blog Service Server is shut down!")
}

func readBlog(client proto.BlogServiceClient, blogId string) {
	log.Println("ReadBlog called")

	readRequest := proto.ReadBlogRequest{
		BlogId: blogId,
	}

	readResponse, e := client.ReadBlog(context.Background(), &readRequest)
	if e != nil {
		log.Fatalf("Failed to call: %v", e)
	}

	blog := readResponse.GetBlog()
	fmt.Printf("Blog Id: %s | Author: %s | Title: %s | Content: %s\n", blog.GetId(), blog.GetAuthorId(), blog.GetTitle(), blog.GetContent())

	log.Printf("Blog has been read: %v", readResponse)
	log.Println("ReadBlog finished")
}
