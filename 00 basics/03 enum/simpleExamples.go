package main

import (
	"fmt"
	"runtime"
)

type hotdog int

var x hotdog
var y int

const (
	MONDAY = iota + 1
	TUESDAY
	_
	THURSDAY
)

func main() {
	fmt.Printf("Value of x: %d\n", x)
	fmt.Printf("Type of x: %T\n", x)

	x = 42
	fmt.Printf("New Value of x: %d\n", x)

	y = int(x)
	fmt.Printf("Value of y: %d\n", y)
	fmt.Printf("Type of y: %T\n", y)

	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)

	s := "James Bond"
	b := []byte(s)

	fmt.Printf("'James Bond' as a String (intern: %T): %v\n", s, s)
	fmt.Printf("'James Bond' as a slice of bytes (intern: %T): %v\n", b, b)

	for i, v := range s {
		fmt.Printf("%d %#U\n", i, v)
	}

	fmt.Println(MONDAY)
	fmt.Println(TUESDAY)
	fmt.Println(THURSDAY)

	fmt.Printf("%d\t\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t\t%b\n", mb, mb)
	fmt.Printf("%d\t\t%b\n", gb, gb)
}

type Direction int

const (
	North Direction = iota
	East
	South
	West
)

func (d Direction) String() string {
	return [...]string{"North", "East", "South", "West"}[d]
}

const (
	_ = iota
	kb = 1 << (iota * 10)
	mb = 1 << (iota * 10)
	gb = 1 << (iota * 10)
)
