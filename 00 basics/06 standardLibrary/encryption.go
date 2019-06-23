package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	s := "secret"

	bytes, e := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)

	if e != nil {
		fmt.Println(e)
	}

	fmt.Println("Origin: ", s)
	fmt.Println("Encoded: ", bytes)

	password := "secret"
	e = bcrypt.CompareHashAndPassword(bytes, []byte(password))

	if e != nil {
		fmt.Println(e)
	} else {
		fmt.Println("Login complete!")
	}
}
