package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type instruction struct {
	Name   string
	Help   string
	Action func(args ...string)
}

func (c instruction) match(s string) bool {
	return c.Name == s
}

func (c instruction) run(args ...string) {
	c.Action(args...)
}

func initInstructionMap() map[string]instruction {
	versionCmd := instruction{
		Name: "version",
		Help: "Shows the current version label of the terminal",
		Action: func(args ...string) {
			fmt.Printf("Actual version: %f\n", 0.1)
		},
	}

	exitCmd := instruction{
		Name: "exit",
		Help: "closes the currently used terminal",
		Action: func(args ...string) {
			os.Exit(0)
		},
	}

	instructionMap := make(map[string]instruction)
	instructionMap["version"] = versionCmd
	instructionMap["exit"] = exitCmd

	return instructionMap
}

func readScan(bytes []byte) (string, []string) {
	input := strings.Split(string(bytes), " ")
	cmd := input[0]
	args := input[1:]
	return cmd, args
}

func main() {
	s := bufio.NewScanner(os.Stdin)
	w := os.Stdout
	instructionMap := initInstructionMap()

	w.WriteString("Basic Terminal (Exit with 'exit')!\n")

	for {
		s.Scan()
		cmd, args := readScan(s.Bytes())

		if instruction, ok := instructionMap[cmd]; ok {
			instruction.run(args...)
		} else {
			fmt.Printf("Command '%s' not found.\n", cmd)
			os.Exit(1)
		}
	}
}
