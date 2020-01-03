package main

import (
	"fmt"
	"os"
)

func main() {
	directoryName := "abc"
	fileName := "test.txt"

	os.Mkdir(directoryName, 0777)
	defer removeDirectory(directoryName)

	file := createFile(fileName)
	moveFile(file.Name(), directoryName)
}

func createFile(fileName string) *os.File {
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return file
}

func moveFile(fileName string, to string) {
	if err := os.Rename(fileName, to+"/"+fileName); err != nil {
		fmt.Println(err)
	}
}

func removeDirectory(directoryName string) {
	if err := os.RemoveAll(directoryName); err != nil {
		fmt.Println(err)
	}
}
