package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	counter int64
	wg      sync.WaitGroup
	mutex   sync.Mutex
)

func main() {
	wg.Add(2)

	go incCounter(1)
	go incCounter(2)

	wg.Wait()
	fmt.Println("Final counter:", counter)
}

func incCounter(id int) {
	defer wg.Done()

	for count := 0; count < 2; count++ {
		// NOTE: only one goroutine can enter the critical section at a time. Not
		// until the call to the Unlock() function is made can another goroutine
		// enter the critical section.
		mutex.Lock()

		// critical section, curly braces are not necessary
		{
			value := counter
			runtime.Gosched()
			value++
			counter = value
		}

		// release the lock and allow any waiting goroutine through.
		mutex.Unlock()
	}
}
