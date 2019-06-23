#10 Concurrency

##waitgroup

Two functions should be exuted. The main function should not terminate until the two functions are completed.

Therefore i use a variable of type sync.WaitGroup. The variable is set to the number of parallel functions. Every parallel function decrease the counter of the wait group. The wait method on the end of the main function waits until the counter reaches zero.

```
func main() {
	waitGroup.Add(<NUMBER OF PARALLEL FUNCTIONS>)        
	go ...                
	waitGroup.Wait()
}

func parallelFunction() {
	...
	waitGroup.Done()
}
```

##twoParallelFunctions

Similiar example like [waitgroup].

##sharedVariableMutex

To analyse race condition problems in the code:

```
go run -race sharedVariableMutex.go
```

or: 

```
go build -race sharedVariableMutex.go
./sharedVariableMutex 
```

##sharedVariableAtomic

Here is the shared variable of an atomic type (int64)

```
var counter int64
...
atomic.AddInt64(&counter, 1)
```

##channels.go

##channelsDeadlock.go

##channelsSuccess.go

Function:
```
c := make(chan int)
go func() { c <- 42 }()
fmt.Println("Channel with Function:", <-c)
```

Buffer:
```
c := make(chan int, 2)
c <- 42
fmt.Println("Channel with Buffer:", <-c)
```

Buffer (Read only):
```
c := make(<-chan int, 2)
```

Buffer (Write only):
```
c := make(chan<- int, 2)
```
