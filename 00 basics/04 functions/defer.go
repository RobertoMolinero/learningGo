package _3_functions

import "fmt"

func main() {
	defer d()
	n()
}

func n() {
	fmt.Println("Normal Func")
}

func d() {
	defer func() {
		fmt.Println("Defer Func (defered)")
	}()

	fmt.Println("Defer Func")
}