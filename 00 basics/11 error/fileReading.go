package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	f, e := os.Open("file.txt")

	if e != nil {
		fmt.Println(e)
		return
	}

	defer f.Close()

	bytes, e := ioutil.ReadAll(f)

	if e != nil {
		fmt.Println(e)
	}

	fmt.Println(string(bytes))
}
