package main

import "fmt"

type shape interface {
	getArea() float64
}

type square struct {
	sideLength float64
}

type triangle struct {
	base   float64
	height float64
}

type circle struct {
	diameter float64
}

func printArea(s shape) {
	switch s.(type) {
	case square:
		fmt.Println("There comes a square!")
	case triangle:
		fmt.Println("There comes a triangle!")
	case circle:
		fmt.Println("There comes a circle!")
	}

	fmt.Println(s.getArea())
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (c circle) getArea() float64 {
	return 3.14 * c.diameter * c.diameter
}

func main() {
	s := square{
		sideLength: 5,
	}
	printArea(s)

	t := triangle{
		height: 5,
		base:   3,
	}
	printArea(t)

	c := circle{
		diameter: 5,
	}
	printArea(c)
}
