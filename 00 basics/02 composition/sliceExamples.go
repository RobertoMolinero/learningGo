package main

import "fmt"

func main() {
	xi := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	fmt.Println(append(xi[:5]))
	fmt.Println(append(xi[5:]))
	fmt.Println(append(xi[2:7]))
	fmt.Println(append(xi[1:6]))
	fmt.Println(append(xi[:3], xi[6:]...))
	fmt.Println(xi)

	xi = append(xi, 52)
	fmt.Println(xi)

	xi = append(xi, 53, 54, 55)
	fmt.Println(xi)

	yi := []int{56, 57, 58, 59, 60}
	xi = append(xi, yi...)
	fmt.Println(xi)

	states := make([]string, 30, 30)
	states = []string{
		"Alabama",
		"Alaska",
		"Arizona",
		"Arkansas",
		"Colorado",
		"Connecticut",
		"Delaware",
		"Florida",
		"Georgia",
		"Hawaii",
		"Idaho",
		"Illinois",
		"Indiana",
		"Iowa",
		"Kalifornien",
		"Kansas",
		"Kentucky",
		"Lincoln (vorgeschlagener Nordwest-Staat)",
		"Louisiana",
		"Maine",
		"Maryland",
		"Massachusetts",
		"Michigan",
		"Minnesota",
		"Mississippi (Bundesstaat)",
		"Missouri",
		"Montana",
		"Nebraska",
		"Nevada",
		"New Hampshire",
		"New Jersey",
		"New Mexico",
		"New York (Bundesstaat)",
		"North Carolina",
		"North Dakota",
		"Ohio",
		"Oklahoma",
		"Oregon",
		"Pennsylvania",
		"Rhode Island",
		"South Carolina",
		"South Dakota",
		"Tennessee",
		"Texas",
		"Utah",
		"Vermont",
		"Virginia",
		"Washington (Bundesstaat)",
		"West Virginia",
		"Wisconsin",
		"Wyoming",
	}

	fmt.Println(states)
	fmt.Println(len(states))
	fmt.Println(cap(states))
}
