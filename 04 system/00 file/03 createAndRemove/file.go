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

func createFile(fileName string) *os.File {
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return file
}

func removeFile(fileName string) {
	if err := os.Remove(fileName); err != nil {
		fmt.Println(err)
	}
}
