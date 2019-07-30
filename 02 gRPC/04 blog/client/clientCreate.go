package main

import (
	"log"
)

func main() {
	log.Println("Client for the Blog Service Server is generated!")

	cc := CreateConnection()
	defer cc.Close()

	client := CreateClient(cc)
	CreateBlog(client)

	log.Println("Client for the Blog Service Server is shut down!")
}
