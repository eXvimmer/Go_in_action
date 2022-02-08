package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	numberOfGoroutines = 4
	taskLoad           = 10 // amount of work to process
)

var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	tasks := make(chan string, taskLoad) // a buffered channel
	wg.Add(numberOfGoroutines)

	// launch goroutines to handle the work
	for gr := 1; gr <= numberOfGoroutines; gr++ {
		go worker(tasks, gr)
	}

	// add a bunch of work to get done
	for post := 1; post <= taskLoad; post++ {
		tasks <- fmt.Sprintf("Task: %d", post)
	}
	close(tasks)
	// Channel is closed, waiting for all the work to be completed.
	// NOTE: When a channel is closed, goroutines can still perform receives on
	// the channel but can no longer send on the channel.
	wg.Wait()
}

func worker(tasks chan string, worker int) {
	defer wg.Done()

	for {
		task, ok := <-tasks
		if !ok {
			fmt.Printf("Worker %d, shutting down\n", worker)
			return
		}

		fmt.Printf("Worker %d started: %s\n", worker, task)

		sleep := rand.Int63n(100)
		time.Sleep(time.Duration(sleep) * time.Microsecond)
		fmt.Printf("Worker %d completed: %s\n", worker, task)
	}
}
