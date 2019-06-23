package main

import "fmt"

/*
 * Problem 001: Multiples of 3 and 5
 *
 * If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.
 * Find the sum of all the multiples of 3 or 5 below 1000.
 */
func main() {

	//	gauss := additionOfGauss(100)
	//gauss := additionOfGaussWithFactor(2, 10)
	gauss := additionOfGaussWithMultipleFactors(1000, 3, 5)
	fmt.Println(gauss)

	xi := []int{1, 2, 3, 4}
	asfsa(xi...)
	asfsa()
}

func asfsa(x ...int) {

}

func additionOfGaussWithMultipleFactors(limit int, factor1 int, factor2 int) int {

	f := factor1 * factor2
	return additionOfGaussWithFactor(1000, factor1) + additionOfGaussWithFactor(1000, factor2) - additionOfGaussWithFactor(1000, f)
}

func additionOfGaussWithFactor(limit int, factor int) int {
	limitWithFactor := limit / factor
	return factor * additionOfGauss(limitWithFactor)
}

func additionOfGauss(limit int) int {
	if limit%2 == 0 {
		return (1 + limit) * (limit / 2)
	}

	return (1+(limit-1))*((limit-1)/2) + limit
}
