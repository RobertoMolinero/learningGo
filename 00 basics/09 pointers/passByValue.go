package main

import "fmt"

func main() {
	x := 0

	foo(x)
	fmt.Println("after foo(): ", x)

	bar(&x)
	fmt.Println("after bar(): ", x)
}

func foo(y int) {
	fmt.Println(y)
	y = 42
	fmt.Println(y)
}

func bar(y *int) {
	fmt.Println(y)
	*y = 42
	fmt.Println(y)
}