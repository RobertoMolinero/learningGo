package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Printf("Hey, i am %s %s and i am %d years old.\n", p.first, p.last, p.age)
}

func main() {
	x := person{
		first: "James",
		last:  "Bond",
		age:   32,
	}

	x.speak()
}
