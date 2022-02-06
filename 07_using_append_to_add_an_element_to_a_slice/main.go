package main

import (
	"fmt"
)

func main() {
	s := []int{10, 20, 30, 40, 50}
	n := s[1:3]       // length: 2, cap: 5-1
	n = append(n, 60) // changed both
	fmt.Println(n, s)
}
