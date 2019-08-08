package main

import (
	"fmt"
	"github.com/urfave/cli"
	"log"
	"os"
)

var app = cli.NewApp()

func info() {
	app.Name = "Euler Problem CLI"
	app.Usage = "A command line to display and execute tasks from the euler problem database."
	app.Author = "Roberto Molinero"
	app.Version = "1.0.0"
}

func commands() {
	app.Commands = []cli.Command{
		{
			Name:  "001",
			Usage: "Problem 001: Multiples of 3 and 5",
			Action: func(c *cli.Context) {
				result := AdditionOfGauss(1000, 3, 5)
				fmt.Println(result)
			},
		},
	}
}

func main() {
	info()
	commands()
	e := app.Run(os.Args)
	if e != nil {
		log.Fatal(e)
	}
}
