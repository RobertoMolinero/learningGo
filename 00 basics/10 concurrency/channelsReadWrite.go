package main

import "fmt"

func main() {
	c := make(chan int)
	cr := make(<-chan int)
	cs := make(chan<- int)

	fmt.Printf("c\t%T\n", c)
	fmt.Printf("cr\t%T\n", cr)
	fmt.Printf("csR\t%T\n", cs)
}
