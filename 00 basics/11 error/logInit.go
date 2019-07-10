package main

import (
	"log"
	"os"
)

var (
	// Trace Fine granular information about the functions used
	Trace   *log.Logger
	// Info Standard information about the functions used
	Info    *log.Logger
	// Warning Warnings about potential sources of error
	Warning *log.Logger
	// Error Messages about operations that could not be completed successfully
	Error   *log.Logger
)

func init() {
	logfile, e := os.Create("withInit.log")
	if e != nil {
		panic(e)
	}

	Trace = log.New(logfile, "TRACE: ", log.Ldate|log.Ltime|log.Lshortfile)
	Info = log.New(logfile, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	Warning = log.New(logfile, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(logfile, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {
	Trace.Println("I have something standard to say")
	Info.Println("Special Information")
	Warning.Println("There is something you need to know about")
	Error.Println("Something has failed")
}
