package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

func main() {
	listener, e := net.Listen("tcp", ":8080")
	if e != nil {
		log.Fatalln(e)
	}

	for {
		conn, e := listener.Accept()
		if e != nil {
			log.Fatalln(e)
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()

	io.WriteString(conn,
		"\nIN-MEMORY DATABASE\n"+
			"USE:\n"+
			"SET <key> <value>\n"+
			"GET <key>\n"+
			"DEL <key>\n"+
			"EXAMPLE:\n"+
			"SET name james\n"+
			"GET name\n")

	data := make(map[string]string)
	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		ln := scanner.Text()
		fs := strings.Fields(ln)

		if len(fs) < 1 {
			continue
		}

		switch fs[0] {
		case "GET":
			k := fs[1]
			v := data[k]
			fmt.Fprintf(conn, "%s\n", v)
		case "SET":
			if len(fs) != 3 {
				fmt.Fprintf(conn, "Expected Value!")
				continue
			}
			k := fs[1]
			v := fs[2]
			data[k] = v
		case "DEL":
			k := fs[1]
			delete(data, k)
		default:
			fmt.Fprintf(conn, "Invalid command!")
		}
	}
}
