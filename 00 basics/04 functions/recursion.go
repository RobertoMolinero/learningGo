package _3_functions

import "fmt"

func main() {
	fmt.Println(4 * 3 * 2 * 1)
	fmt.Println(factorial(4))
	fmt.Println(factorialWithoutRecursion(4))
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func factorialWithoutRecursion(n int) int {
	var total = 1
	for ; n > 0; n-- {
		total *= n
	}
	return total
}
