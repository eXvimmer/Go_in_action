package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	runtime.GOMAXPROCS(1)
	wg.Add(2)
	fmt.Println("Create Goroutines")

	go printPrime("A")
	go printPrime("B")

	fmt.Println("Waiting to Finish")
	wg.Wait()
	fmt.Println("\nTerminating Program")
}

// printPrimes displays prime numbers for the first 5000 numbers
func printPrime(prefix string) {
	defer wg.Done()

next:
	for outer := 2; outer < 5000; outer++ {
		for inner := 2; inner < outer; inner++ {
			if outer%inner == 0 { // not a prime number
				continue next
			}
		}
		fmt.Printf("%s: %d\n", prefix, outer) // a prime number
	}
	fmt.Println("Completed", prefix)
}
