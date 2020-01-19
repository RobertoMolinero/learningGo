package main

import (
	"bufio"
	"os"
	"strings"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	w := os.Stdout

	w.WriteString("Basic Terminal (Exit with 'exit')!\n")

	for {
		s.Scan()

		input := strings.Split(string(s.Bytes()), " ")
		cmd := input[0]
		args := input[1:]

		w.WriteString("Your input: [")
		w.WriteString("command: " + cmd)
		w.WriteString(", args: ")

		for _, arg := range args {
			w.WriteString(arg + " ")
		}

		w.WriteString("]\n")

		if s.Text() == "exit" {
			return
		}
	}
}
