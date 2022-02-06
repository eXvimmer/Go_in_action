package main

/*
 * This third index gives you control over the capacity of the new slice. The
 * purpose is not to increase capacity, but to restrict the capacity. As you'll
 * see, being able to restrict the capacity of a new slice provides a level of
 * protection to the underlying array and gives you more control over append
 * operations.
 */

import (
	"fmt"
)

func main() {
	s := []string{"Apple", "Orange", "Plum", "Banana", "Grape"}
	fmt.Println(len(s), cap(s), s)
	// NOTE: By having the option to set the capacity of a new slice to be the
	// same as the length, you can force the first append operation to detach the
	// new slice from the underlying array. Detaching the new slice from its
	// original source array makes it safe to change.
	n := s[2:3:4] // restrict the capacity
	fmt.Println(len(n), cap(n), n)
	// NOTE: for slice[i:j:k]
	// length: j - i         -> 3 - 2 = 1
	// capacity: k - i       -> 4 - 1 = 2

	// NOTE: if you set a capacity that's larger than the available capacity,
	// you'll get a runtime error.

	m := s[2:3:3] // length: 1, capacity: 1
	fmt.Println(m, len(m), cap(m))
	m = append(m, "Kiwi") // length: 3, capacity: 2, s is not changed
	fmt.Println(m, len(m), cap(m))
	m = append(m, "Cherry") // length: 3, capacity: 4
	fmt.Println(m, len(m), cap(m))
	fmt.Println(s, len(s), cap(s))
}
