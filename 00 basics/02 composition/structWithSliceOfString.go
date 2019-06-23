package main

import "fmt"

type individual struct {
	firstName               string
	lastName                string
	favoriteIceCreamFlavors []string
}

func main() {
	roberto := individual{
		firstName:               "El Patito",
		lastName:                "de la Banera",
		favoriteIceCreamFlavors: []string{"Malagha", "Chocolate"},
	}

	rob := individual{
		firstName:               "Rob",
		lastName:                "Pike",
		favoriteIceCreamFlavors: []string{"Orange", "Vanilla"},
	}

	fmt.Println(roberto.firstName, roberto.lastName)

	for i, v := range roberto.favoriteIceCreamFlavors {
		fmt.Printf("Index: %d\tFlavor: %s\n", i, v)
	}

	m := map[string]individual{
		roberto.lastName: roberto,
		rob.lastName:     rob,
	}

	for k, v := range m {
		fmt.Printf("Key: %s\n", k)
		fmt.Println(v.firstName)
		fmt.Println(v.lastName)
		for i, s := range v.favoriteIceCreamFlavors {
			fmt.Printf("Index: %d\tFavor: %s\n", i, s)
		}
	}
}
