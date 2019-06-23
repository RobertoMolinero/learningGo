package main

import "fmt"

func main()  {

	var varColors map[string]string
	fmt.Println(varColors)

	makeColors := make(map[string]string)
	makeColors["white"] = "#ffffff"
	fmt.Println(makeColors)

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
	}

	colors["white"] = "#ffffff"
	fmt.Println(colors)

	// delete(colors, "green")
	// fmt.Println(colors)

	printMap(colors)
}

func printMap(m map[string]string) {
	for key, value := range m {
		fmt.Println(key + ": " + value)
	}
}