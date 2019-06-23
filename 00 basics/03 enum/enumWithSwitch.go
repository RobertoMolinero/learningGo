package main

import "fmt"

type WindDirection int

const (
	North WindDirection = iota
	East
	South
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
