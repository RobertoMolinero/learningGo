package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var cmdFunc func(args []string)

func exitCmd(args [] string) {
	os.Exit(0)
}

func versionCmd(args [] string) {
	fmt.Printf("Actual version: %f\n", 0.1)
}

func readScan(bytes []byte) (error, string, []string) {
	input := strings.Split(string(bytes), " ")
	cmd := input[0]
	args := input[1:]
	return nil, cmd, args
}

func main() {
	s := bufio.NewScanner(os.Stdin)
	w := os.Stdout

	cmdMap := make(map[string]func(args []string))
	cmdMap["exit"] = exitCmd
	cmdMap["version"] = versionCmd

	w.WriteString("Basic Terminal (Exit with 'exit')!\n")

	for {
		s.Scan()
		_, cmd, args := readScan(s.Bytes())

		cmdFunc = cmdMap[cmd]

		if cmdFunc == nil {
			fmt.Printf("Command '%s' not found.\n", cmd)
			os.Exit(1)
		}

		cmdFunc(args)
	}
}
