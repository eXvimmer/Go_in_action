package main

import (
	"fmt"
)

func main() {
	s := []string{99: ""} // NOTE: length and capacity is set to 100
	fmt.Println(len(s), cap(s), s)
	var n []int // a nil slice: pointer: nil, length: 0, capacity: 0
	fmt.Println(n)
	e := make([]int, 0) // an empty slice
	f := []int{}        // an empty slice
	fmt.Println(e, f)
}
