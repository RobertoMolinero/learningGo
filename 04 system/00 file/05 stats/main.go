package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("test.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer os.Remove(file.Name())

	stat, err := os.Stat(file.Name())
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Name:\t %s \n", stat.Name())
	fmt.Printf("IsDir:\t %t \n", stat.IsDir())
	fmt.Printf("Size:\t %d \n", stat.Size())
	fmt.Printf("Sys:\t %v \n", stat.Sys())
	fmt.Printf("Mode:\t %v \n", stat.Mode())
	fmt.Printf("ModTime:\t %v \n", stat.ModTime())
}
