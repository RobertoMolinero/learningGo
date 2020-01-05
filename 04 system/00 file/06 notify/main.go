package main

import (
	"github.com/fsnotify/fsnotify"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	subfolder := "test"

	if err := os.Mkdir(subfolder, 0777); err != nil {
		log.Fatal(err)
	}

	_, done := monitoring(subfolder)
	<- done
}

func monitoring (subfolder string) (error, chan bool) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-c
		os.RemoveAll(subfolder)
		os.Exit(1)
	}()

	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				log.Printf("Operation: [%s] detected on: \"%s\"", event.Op.String(), event.Name)
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println(err)
			}
		}
	}()

	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	if err := watcher.Add(wd + "/" + subfolder); err != nil {
		log.Fatal(err)
	}

	done <- true
	return nil, done
}
