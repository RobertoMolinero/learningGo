package main

import "fmt"

func main() {
	m := map[string][]string{
		"james":      {"Martinis", "Women"},
		"moneypenny": {"James", "Literature", "Computer Science"},
		"drNo":       {"Evil", "Ice Cream", "Sunsets"},
	}

	m["goldfinger"] = []string{"Gold"}

	for k, v := range m {
		fmt.Printf("Key: %s\n", k)
		for i, s := range v {
			fmt.Printf("Index: %d\tValue: %s\n", i, s)
		}
	}
}
