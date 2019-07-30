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
	updateBlog(client, blogId)

	log.Println("Client for the Blog Service Server is shut down!")
}

func updateBlog(client proto.BlogServiceClient, blogId string) {
	log.Println("UpdateBlog called")

	updateRequest := proto.UpdateBlogRequest{
		Blog: &proto.Blog{
			Id:       blogId,
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

	log.Printf("Blog has been updated: %v", updateResponse)
	log.Println("UpdateBlog finished")
}
