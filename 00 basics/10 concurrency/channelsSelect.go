package main

import "fmt"

func main() {
	even := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	go sendThree(even, odd, quit)
	receiveThree(even, odd, quit)

	fmt.Println("End of the main method.")
}

func sendThree(e, o, q chan<- int) {
	for i := 0; i < 5; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	close(q)
}

func receiveThree(e, o, q <-chan int) {
	for {
		select {
		case v := <-e:
			fmt.Println("From the even channel:", v)
		case v := <-o:
			fmt.Println("From the odd channel:", v)
		case v, ok := <-q:
			if !ok {
				fmt.Println("From the quit channel:", v, ok)
				fmt.Println("The channel is already closed!")
				return
			}
			fmt.Println("From the quit channel:", v)
			return
		}
	}
}
