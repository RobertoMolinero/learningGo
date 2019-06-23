package main

import "fmt"

func main() {
	x := []int{1, 2, 3}
	fmt.Println(x)

	y := []int{7, 8, 9}
	x = append(x, y...)
	fmt.Println(x)

	x = append(x[:2], x[4:]...)
	fmt.Println(x)

	m := map[string]int {
		"Bart": 8,
		"Lisa": 6,
		"Maggie": 1,
	}

	fmt.Println(m)

	if v, ok := m["Bart"]; ok {
		fmt.Println(v)
		fmt.Println(ok)
	}

	if v, ok := m["Homer"]; ok {
		fmt.Println(v)
		fmt.Println(ok)
	} else {
		fmt.Println("No Homer Exception!")
	}

	delete(m, "Lisa")
	fmt.Println(m)

	delete(m, "Homer")
	fmt.Println(m)

	delete(m, "Homer")
}
