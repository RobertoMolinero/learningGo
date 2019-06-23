package main

import (
	"fmt"
	"log"
)

func main() {
	_, e := sqrtWithOwnError(-2)

	if e != nil {
		log.Fatalln(e)
	}
}

func sqrtWithOwnError(f float64) (float64, error) {
	if f < 0 {
		m := fmt.Errorf("error: value %v is lower than 0!\n", f)
		return 0, m
	}
	return 42, nil
}
