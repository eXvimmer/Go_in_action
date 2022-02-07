package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	// fmt.Println(runtime.NumCPU())

	// Allocate 1 logical processor for the program to finish
	runtime.GOMAXPROCS(1)

	var wg sync.WaitGroup
	wg.Add(2) // 2, one for each goroutine

	fmt.Println("Starting Goroutines")

	go func() {
		defer wg.Done()

		for count := 0; count < 3; count++ {
			for char := 'a'; char < 'a'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
	}()

	go func() {
		defer wg.Done()

		for count := 0; count < 3; count++ {
			for char := 'A'; char < 'A'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
	}()

	fmt.Println("Waiting to finish")
	wg.Wait()
	fmt.Println("\nTerminating program.")
}
