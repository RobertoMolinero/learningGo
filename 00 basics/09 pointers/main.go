package main

import "fmt"

func main() {
	a := 42
	fmt.Printf("%d (%T)\n", a, a)
	fmt.Printf("%d (%T)\n", &a, &a)

	b := &a
	fmt.Printf("%d (%T)\n", b, b)
	fmt.Printf("%d (%T)\n", *b, *b)

	*b = 43
	fmt.Println(a)
	fmt.Println(*b)
}
