package main

import (
	"fmt"
	"github.com/pkg/errors"
	"log"
)

var MathError = errors.New("Error: f can not be lower than 0!")

func main() {
	fmt.Printf("%T\n", MathError)

	_, e := sqrt(-2)

	if e != nil {
		log.Fatalln(e)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, MathError
	}
	return 42, nil
}
