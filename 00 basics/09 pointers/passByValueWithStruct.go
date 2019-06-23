package main

import "fmt"

type person struct {
	first string
	last  string
}

func main() {
	x := person{
		first: "James",
		last:  "Bond",
	}

	fmt.Println("before changeMe(): ", x)
	changeMe(&x)
	fmt.Println("after changeMe(): ", x)
}

func changeMe(p *person) {
	p.first = "A"
	// (*p).first = "A"
}
