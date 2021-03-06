package main

import (
	"fmt"
	"sync"
)

//START
var syncedCounter struct {
	lock sync.Mutex
	n    int
}

func inc() {
	syncedCounter.lock.Lock()
	syncedCounter.n++
	syncedCounter.lock.Unlock()
}

func main() {
	fmt.Printf("n: %d\n", syncedCounter.n)
	for i := 1; i < 10; i++ {
		go func() { // goroutine + anonymous function
			inc()
			fmt.Printf("n: %d\n", syncedCounter.n)
		}() // invoke
	}
	fmt.Printf("n: %d\n", syncedCounter.n)
}

//END
