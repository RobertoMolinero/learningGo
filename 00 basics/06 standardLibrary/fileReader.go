package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Println("More or less than 1 Argument")
		fmt.Println("Please Use:")
		fmt.Println("go run concurrency.go <filename>")
		os.Exit(1)
	}

	file, err := os.Open(os.Args[1])

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, file)
	fmt.Println()
}
