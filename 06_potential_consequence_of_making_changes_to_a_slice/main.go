package main

import (
	"fmt"
)

func main() {
	s := []int{10, 20, 30, 40, 50}
	n := s[1:3]
	n[1] = 99 // NOTE: changed in s and n
	fmt.Println(s, n)
}
