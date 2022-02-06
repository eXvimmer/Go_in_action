package main

import (
  "fmt"
)

func main() {
  s := []int{10, 20, 30, 40}
  n := append(s, 50) // NOTE: n has a new underlying array
  fmt.Println(s, n)
  n[0] = 100
  fmt.Println(s, n)
}
