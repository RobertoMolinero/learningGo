package main

import "fmt"

func main() {
	n, e := fmt.Println("Hello World!")

	if e != nil {
		fmt.Println(e)
	}

	fmt.Println(n)
}
