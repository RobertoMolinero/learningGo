package main

import "fmt"

func main() {
	successfulFunc()
	successfulBuffer()
}

func successfulFunc() {
	c := make(chan int)

	go func() {
		c <- 42
	}()

	v, ok := <-c

	fmt.Println("Channel (open) with Function:")
	fmt.Printf("Value:\t%d\n", v)
	fmt.Printf("Status:\t%t\n", ok)

	close(c)

	v, ok = <-c

	fmt.Println("Channel (closed) with Function:")
	fmt.Printf("Value:\t%d\n", v)
	fmt.Printf("Status:\t%t\n", ok)
}

func successfulBuffer() {
	c := make(chan int, 2)
	c <- 42
	c <- 43
	fmt.Println("Channel with Buffer:", <-c)
	fmt.Println("Channel with Buffer:", <-c)
}
