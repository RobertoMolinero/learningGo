package main

import (
	"../proto"
	"context"
	"log"
)

func main() {
	log.Println("Client for the Blog Service Server is generated!")

	cc := CreateConnection()
	defer cc.Close()

	client := CreateClient(cc)
	blogId := CreateBlog(client)
	deleteBlog(client, blogId)

	log.Println("Client for the Blog Service Server is shut down!")
}

func deleteBlog(client proto.BlogServiceClient, blogId string) {
	log.Println("DeleteBlog called")

	deleteRequest := proto.DeleteBlogRequest{
		BlogId: blogId,
	}

	deleteResponse, e := client.DeleteBlog(context.Background(), &deleteRequest)
	if e != nil {
		log.Fatalf("Failed to call: %v", e)
	}

	log.Printf("Blog has been deleted: %v", deleteResponse)
	log.Println("DeleteBlog finished")
}
