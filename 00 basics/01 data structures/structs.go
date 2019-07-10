package main

import "fmt"

type contactInfo struct {
	email string
	zip   int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {

	lisa := person{
		firstName: "Lisa",
		lastName:  "Simpson",
	}

	fmt.Println(lisa)
	fmt.Printf("%+v\n", lisa)

	var bart person

	bart.firstName = "Bart"
	bart.lastName = "Simpson"

	fmt.Printf("%+v\n", bart)

	homer := person{
		firstName: "Homer",
		lastName:  "Simpson",
		contactInfo: contactInfo{
			email: "homer@simpson.com",
			zip:   12345,
		},
	}

	homer.print()

	homer.updateFirstName("Homerino")
	homer.print()

	fmt.Println(&homer)
}

func (p *person) updateFirstName(newFirstName string) {
	(*p).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}
