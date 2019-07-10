package main

import (
	"fmt"
	"sort"
)

type character struct {
	First string
	Age   int
}

// ByAge A slice of characters that should be sorted by the age of the elements
type ByAge []character

func (a ByAge) Len() int {
	return len(a)
}

func (a ByAge) Less(i, j int) bool {
	return a[i].Age < a[j].Age
}

func (a ByAge) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

// ByName A slice of characters that should be sorted by the name of the elements
type ByName []character

func (n ByName) Len() int {
	return len(n)
}

func (n ByName) Less(i, j int) bool {
	return n[i].First < n[j].First
}

func (n ByName) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}

func main() {
	p1 := character{"James", 32}
	p2 := character{"Moneypenny", 27}
	p3 := character{"Q", 64}
	p4 := character{"M", 56}

	people := []character{p1, p2, p3, p4}

	fmt.Println("Unsorted: ", people)
	sort.Sort(ByAge(people))
	fmt.Println("Sorted by Age: ", people)

	fmt.Println("--- --- ---")

	fmt.Println("Unsorted: ", people)
	sort.Sort(ByName(people))
	fmt.Println("Sorted by Name: ", people)
}
