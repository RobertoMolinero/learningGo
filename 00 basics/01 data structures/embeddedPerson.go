package main

import "fmt"

type embeddedPerson struct {
	first string
	last  string
	age   int
}

type clown struct {
	embeddedPerson
	clownShoes bool
}

func CreatePerson() {
	p1 := embeddedPerson{
		first: "Bart",
		last:  "Simpson",
		age:   32,
	}

	p2 := embeddedPerson{
		first: "Lisa",
		last:  "Simpson",
		age:   27,
	}

	p3 := clown{
		embeddedPerson: embeddedPerson{
			first: "Krusty",
			last:  "Clown",
			age:   40,
		},
		clownShoes: true,
	}

	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p3)

	fmt.Println(p1.first, p1.last, p1.age)
	fmt.Println(p2.first, p2.last, p2.age)
	fmt.Println(p3.first, p3.last, p3.age, p3.clownShoes)
}
