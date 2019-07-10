package main

import "fmt"

func main() {

	foo()

	func() {
		fmt.Println("anonymous func ran")
	}()

	func(x int) {
		fmt.Printf("anonymous func with parameters(%d) ran\n", x)
	}(42)

	f := func() {
		fmt.Println("a func expression")
	}
	f()

	g := func(x int) {
		fmt.Printf("a func expression with parameter(%d)\n", x)
	}

	g(1984)

	h := bar()
	fmt.Printf("%T\n", h)
	fmt.Println(h())

	fmt.Println(bar()())
}

func foo() {
	fmt.Println("foo ran")
}

func bar() func() int {
	return func() int {
		return 0
	}
}
