package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {

	listener, e := net.Listen("00 tcp", ":8080")

	if e != nil {
		log.Fatalln(e)
	}

	for {
		conn, e := listener.Accept()

		if e != nil {
			log.Fatalln(e)
		}

		io.WriteString(conn, "\nHello from TCP server\n")
		fmt.Fprintln(conn, "How is your day?")
		fmt.Fprintln(conn, "%v", "Well, I hope!")

		conn.Close()
	}
}
