package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type person struct {
	First  string
	Last   string
	Age    int
	Drinks []string
}

func main() {
	marshal()
	unmarshal()
}

func marshal() {
	sledge := person{
		First:  "Sledge",
		Last:   "Hammer",
		Age:    35,
		Drinks: []string{"Brown Ale", "Water"},
	}

	bytes, e := json.Marshal(sledge)

	if e != nil {
		fmt.Println(e)
	}

	n, e := os.Stdout.Write(bytes)

	if e != nil {
		fmt.Println(e)
	}

	fmt.Println()
	fmt.Printf("Number of Bytes written: %d\n", n)
}

var jsonBlob = []byte(`[
	{"First":"Sledge","Last":"Hammer","Age":35,"Drinks":["Brown Ale","Water"]},
	{"First":"A","Last":"B","Age":25,"Drinks":["Whisky","Beer"]}
]`)

func unmarshal() {
	var persons []person
	e := json.Unmarshal(jsonBlob, &persons)

	if e != nil {
		fmt.Println(e)
	}

	for i, p := range persons {
		fmt.Printf("Person %d\n", i)
		fmt.Printf("\tFirst: %s\n", p.First)
		fmt.Printf("\tLast: %s\n", p.Last)
		fmt.Printf("\tAge: %d\n", p.Age)
		fmt.Println("\tDrinks:")
		for j, d := range p.Drinks {
			fmt.Printf("\t\tIndex: %d\tDrink: %s\n", j, d)
		}
	}

	fmt.Printf("%+v\n", persons)
}
