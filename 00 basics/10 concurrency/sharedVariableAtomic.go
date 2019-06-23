package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	fmt.Println("CPUs:", runtime.NumCPU())

	var counter int64
	const gs = 100

	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			fmt.Printf("Counter:\t%d\tGo Routines:\t%d\n", atomic.LoadInt64(&counter), runtime.NumGoroutine())
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Final Counter:\t", counter)
}
