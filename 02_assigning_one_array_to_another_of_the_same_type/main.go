package main

import (
  "fmt"
)

func main() {
  a := [5]string{"red", "green", "bule", "orange", "yellow"}
  b := a // NOTE: copy the values from a to b
  fmt.Println(b)
  b[2] = "violet" // NOTE: a is not changed
  fmt.Println(a, b)
}
