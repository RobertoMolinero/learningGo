package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	name := "Roberto Molinero"

	tpl := `
	<!DOCTYPE html>
	<html lang="en">
		<head>
			<meta charset="UTF-8">
			<title>Hello World!</title>
		</head>
		<body>
			<h1>` + name + `</h1>
		</body>
	</html>
	`

	fmt.Println(tpl)

	file, e := os.Create("index.html")

	if e != nil {
		log.Fatalln("Error creating File",e)
	}

	defer file.Close()

	io.Copy(file, strings.NewReader(tpl))
}
