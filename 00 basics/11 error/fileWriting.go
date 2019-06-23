package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	f, e := os.Create("file.txt")

	if e != nil {
		fmt.Println(e)
		return
	}

	defer f.Close()

	r := strings.NewReader("The weather is nice!")
	io.Copy(f, r)
}
