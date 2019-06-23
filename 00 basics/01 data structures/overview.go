package main

import "fmt"

type human interface {
	speak()
}

type character struct {
	fname string
	lname string
}

type secretAgent struct {
	character
	licenseToKill bool
}

func (c character) speak() {
	fmt.Println(c.fname, c.lname, `says: "Good Morning, James."`)
}

func (s secretAgent) speak() {
	fmt.Println(s.fname, s.lname, `says: "Shaken, not stirred."`)
}

func saySomething(h human) {
	h.speak()
}

func main() {
	// Slice
	xi := []int{1, 2, 3}
	fmt.Println(xi)

	// Map
	m := map[string]int{
		"Roberto": 38,
		"Job":     42,
	}
	fmt.Println(m)

	// Struct
	c := character{
		fname: "Miss",
		lname: "Moneypenny",
	}
	fmt.Println(c)

	s := secretAgent{
		character{
			fname: "Mr.",
			lname: "Bond",
		}, true,
	}
	fmt.Println(s)

	fmt.Println("Call Receiver Function: ")
	c.speak()
	s.speak()
	s.character.speak()

	fmt.Println("Call Receiver Function through Interface Method: ")
	saySomething(c)
	saySomething(s)
	saySomething(s.character)
}
