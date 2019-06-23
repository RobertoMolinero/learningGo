package main

import "fmt"

func main() {
	xs := []string{"James", "Bond", "Shaken, not stirred"}
	xt := []string{"Miss", "Moneypenny", "Hellooooo, James"}
	x := [][]string{xs, xt}

	fmt.Println(x)

	for i, sArr := range x {
		fmt.Printf("Record: %d\n", i)
		for j, s := range sArr {
			fmt.Printf("Index: %d\tValue: %s\n", j, s)
		}
	}
}
