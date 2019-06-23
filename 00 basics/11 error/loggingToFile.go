package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	logFile, e := os.Create("log.txt")

	if e != nil {
		fmt.Println(e)
	}

	defer logFile.Close()
	log.SetOutput(logFile)

	file, e := os.Open("no-file.txt")

	if e != nil {
		//fmt.Println(e)
		log.Println(e)
		//log.Fatalln(e)
		//panic(e)
	}

	defer file.Close()
	fmt.Println("Check the log.txt in the directory!")
}
