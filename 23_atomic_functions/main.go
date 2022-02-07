package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var (
	counter int64
	wg      sync.WaitGroup
)

func main() {
	// NOTE: Atomic functions provide low-level locking mechanisms for
	// synchronizing access to integers and pointers.
	wg.Add(2)

	go incCounter(1)
	go incCounter(2)

	wg.Wait()
	fmt.Println("Final counter:", counter)
}

func incCounter(id int) {
	defer wg.Done()

	for count := 0; count < 2; count++ {
		// safely add 1 to counter
		atomic.AddInt64(&counter, 1)
		runtime.Gosched()
	}
}
