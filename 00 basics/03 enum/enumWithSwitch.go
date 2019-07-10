package main

import "fmt"

// WindDirection The numeric coding of the Wind Directions
type WindDirection int

const (
	// North The Direction North (Up)
	North WindDirection = iota
	// East The Direction East (Right)
	East
	// South The Direction South (Down)
	South
	// West The Direction West (Left)
	West
)

func (d WindDirection) String() string {
	return [...]string{
		"North",
		"East",
		"South",
		"West",
	}[d]
}

func main() {
	d := North
	fmt.Print(d)

	switch d {
	case North:
		fmt.Println(" goes up.")
	case South:
		fmt.Println(" goes down.")
	case East:
		fmt.Println(" goes right.")
	case West:
		fmt.Println(" goes left.")
	default:
		fmt.Println(" stays put.")
	}
}
