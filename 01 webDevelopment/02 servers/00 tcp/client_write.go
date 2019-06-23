package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {

	conn, e := net.Dial("00 tcp", "localhost:8080")

	if e != nil {
		log.Fatalln(e)
	}

	defer conn.Close()

	fmt.Fprintln(conn, "I dialed you.")
}
