package main

/*
 * Problem 001: Multiples of 3 and 5
 *
 * If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.
 * Find the sum of all the multiples of 3 or 5 below 1000.
 */
func AdditionOfGauss(limit int, factorOne int, factorTwo int) int {
	leastCommonMultiple := factorOne * factorTwo
	return additionOfGaussWithFactor(limit, factorOne) + additionOfGaussWithFactor(limit, factorTwo) - additionOfGaussWithFactor(limit, leastCommonMultiple)
}

func additionOfGaussWithFactor(limit int, factor int) int {
	// The limit itself is not included
	limit -= 1
	limitWithFactor := limit / factor
	return factor * additionOfGauss(limitWithFactor)
}

func additionOfGauss(limit int) int {
	if limit%2 == 0 {
		return (1 + limit) * (limit / 2)
	}
	return (1+(limit-1))*((limit-1)/2) + limit
}
