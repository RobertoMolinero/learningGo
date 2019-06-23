package main

import (
	"bufio"
	"fmt"
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

		go handleRead(conn)
	}
}

func handleRead(conn net.Conn) {

	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
	}

	defer conn.Close()

	fmt.Println("Code got here.")
}
