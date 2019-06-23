package main

import (
	"fmt"
	"sync"
)

var waitGroup sync.WaitGroup

func main() {
	waitGroup.Add(2)
	go fooTest()
	go barTest()
	waitGroup.Wait()
}

func fooTest() {
	fmt.Println("Foo")
	waitGroup.Done()
}

func barTest() {
	fmt.Println("Bar")
	waitGroup.Done()
}
