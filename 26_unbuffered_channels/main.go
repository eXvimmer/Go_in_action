package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	court := make(chan int) // an unbuffered channel
	wg.Add(2)

	go player("Nadal", court)
	go player("Djokovic", court)

	// start the set
	court <- 1

	wg.Wait()
}

func player(name string, court chan int) {
	defer wg.Done()

	for {
		ball, ok := <-court
		// This locks the goroutine until a send is performed on the channel
		if !ok {
			fmt.Printf("%s won\n", name)
			return
		}

		n := rand.Intn(100)
		if n%13 == 0 {
			fmt.Printf("%s missed\n", name)
			// close channel to signal we lost
			close(court)
			return
		}

		fmt.Printf("%s hits %d\n", name, ball)
		ball++ // increment the hit count
		court <- ball
		// At this point, both goroutines are locked until the exchange is made.
	}
}
