package main

import (
	"fmt"
	"github.com/pkg/errors"
	"log"
)

// ErrMath An error that occurs within a mathematical operation
var ErrMath = errors.New("Error: f can not be lower than 0!")

func main() {
	fmt.Printf("%T\n", ErrMath)
	_, e := sqrt(-2)

	if e != nil {
		log.Fatalln(e)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, ErrMath
	}
	return 42, nil
}
