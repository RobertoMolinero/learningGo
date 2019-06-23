package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := fanIn(boring("Joe"), boring("Ann"))

	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}

	fmt.Println("You're both boring; I'm leaving.")
}

func fanIn(x, y <-chan string) <-chan string {
	c := make(chan string)

	go func() {
		for {
			c <- <-x
		}
	}()

	go func() {
		for {
			c <- <-y
		}
	}()

	return c
}

func boring(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}
