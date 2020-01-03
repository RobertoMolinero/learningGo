package main

import "os"

func main() {
	os.Mkdir("abc", 0777)
	defer os.Remove("abc")

	os.MkdirAll("a/b/c", 0777)
	defer os.RemoveAll("a")
}
