package main

import (
	"fmt"
	"os"
)

func main() {
	file := createFile("test.txt")
	defer file.Close()
	defer removeFile(file.Name())
}

func createFile(filename string) *os.File {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return file
}

func removeFile(filename string) {
	if err := os.Remove(filename); err != nil {
		fmt.Println(err)
	}
}
