package main

import (
	"fmt"
	"os"
)

func main() {
	directoryName := "abc"
	directoryNameParent := "def"
	fileName := "test.txt"

	// Create a first folder
	os.Mkdir(directoryName, 0777)

	// Create a file and move it to the first folder
	file := createFile(fileName)
	moveFile(file.Name(), directoryName)

	// Create a second folder and move the first one to the second
	os.Mkdir(directoryNameParent, 0777)
	moveFile(directoryName, directoryNameParent)
	defer removeDirectory(directoryNameParent)
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
