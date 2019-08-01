package main

import "fmt"

func CreateAnomymous() {
	x := struct {
		s  string
		xi []string
		m  map[int]string
		mx map[int][]string
	}{
		s:  "a",
		xi: []string{"1", "2", "3"},
		m: map[int]string{
			1: "a",
			2: "b",
			3: "c",
		},
		mx: map[int][]string{
			1: {"a", "b", "c"},
			2: {"d", "e", "f"},
		},
	}

	fmt.Println(x)

	anonymous := struct {
		first string
		last  string
	}{
		first: "A",
		last:  "X",
	}

	fmt.Println(anonymous)
}
