package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	res, errHTTP := http.Get("http://google.com")

	if errHTTP != nil {
		fmt.Println("Error: ", errHTTP)
		os.Exit(1)
	}

	fmt.Println("Status: ", res.Status)
	fmt.Println("Status Code: ", res.StatusCode)

	lw := logWriter{}
	_, errHTTP = io.Copy(lw, res.Body)

	if errHTTP != nil {
		fmt.Println("Error: ", errHTTP)
		os.Exit(1)
	}
}

type logWriter struct{}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println("Length of the Response Body: ", len(bs))
	fmt.Println(string(bs))
	return len(bs), nil
}
