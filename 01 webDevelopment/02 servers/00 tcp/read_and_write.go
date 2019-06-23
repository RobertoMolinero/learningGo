package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
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

		go handleReadWrite(conn)
	}
}

func handleReadWrite(conn net.Conn) {
	deadline := conn.SetDeadline(time.Now().Add(10 * time.Second))

	if deadline != nil {
		log.Fatalln(deadline)
	}

	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		fmt.Fprintf(conn, "I heard you say: %s\n", ln)
	}

	defer conn.Close()

	fmt.Println("Code got here.")
}
