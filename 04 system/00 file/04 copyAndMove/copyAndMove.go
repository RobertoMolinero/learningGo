package main

import (
	"fmt"
	"io"
	"os"
)

func MoveFile(from string, to string) {
	if err := os.Rename(from, to+"/"+from); err != nil {
		fmt.Println(err)
	}
}

func CopyFile(from string, to string) (int64, error) {
	src, err := os.Open(from)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	defer src.Close()

	dst, err := os.OpenFile(to, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	defer dst.Close()

	return io.Copy(dst, src)
}
