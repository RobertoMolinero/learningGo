package main

import "fmt"

func main() {
	for i := range [11]int{} {
		fmt.Printf("The number %d is %v.\n", i, evenOdd(i))
	}
}

func evenOdd(i int) string {
	if (i % 2) == 0 {
		return "even"
	}

	return "odd"
}
