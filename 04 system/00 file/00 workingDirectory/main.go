package main

import (
	"fmt"
	"os"
)

func main() {
	wdStart := getWd()
	fmt.Println("Working directory: ", wdStart)

	// Go up one level
	changeWd(getWd() + "/../")
	wdParent := getWd()
	fmt.Println("Working directory: ", wdParent)

	// Go to the root
	changeWd("/")
	wdRoot := getWd()
	fmt.Println("Working directory: ", wdRoot)
}

func getWd() string {
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return wd
}

func changeWd(wd string) {
	if err := os.Chdir(wd); err != nil {
		fmt.Println(err)
		return
	}
}
