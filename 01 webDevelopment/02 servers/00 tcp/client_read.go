package main

import (
	"fmt"
	"io/ioutil"
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

	bytes, e := ioutil.ReadAll(conn)

	if e != nil {
		log.Fatalln(e)
	}

	fmt.Println(string(bytes))
}
