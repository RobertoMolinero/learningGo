package main

import (
	"bufio"
	"os"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	w := os.Stdout

	w.WriteString("Basic Terminal Start (Exit with CTRL-C)!\n")

	for {
		s.Scan()

		if s.Text() == "exit" {
			return
		}

		w.WriteString("Your input:")
		w.Write(s.Bytes())
		w.WriteString("\n")
	}
}
