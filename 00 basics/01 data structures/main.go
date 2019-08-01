package main

import "fmt"

func main() {
	fmt.Println("*** Create and Update a Struct ***")
	CreateAndUpdateStruct()

	fmt.Println("\n*** Create an embedded Struct (clown -> embeddedPerson) ***")
	CreatePerson()

	fmt.Println("\n*** Create an embedded Struct (vehicle -> truck / vehicle -> sedan) ***")
	CreateVehicle()

	fmt.Println("\n*** Create an anonymous Struct ***")
	CreateAnomymous()
}
