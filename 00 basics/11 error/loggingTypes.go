package main

import (
	"os"
)

func main() {
	_, e := os.Open("no-file.txt")

	if e != nil {
		//fmt.Println(e)
		//log.Println(e)
		//log.Fatalln(e)
		panic(e)
	}
}
