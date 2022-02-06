package main

import (
  "fmt"
)

func main() {
  // NOTE: range is making a copy of the value, not returning a reference.
  s := []int{10, 20, 30, 40, 50}
  for i, v := range s {
    fmt.Printf("value: %d, value_addr: %X, element_addr: %X\n", v, &v, &s[i])
  }
  // NOTE: the address for the v is always the same, because it's a variable
  // that contains a copy.
}

