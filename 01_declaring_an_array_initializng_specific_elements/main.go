package main

import (
	"fmt"
)

func main() {
	arr := [5]int{1: 10, 2: 20}
	fmt.Println(arr)
	pa := [5]*int{1: new(int), 3: new(int)}
  fmt.Println(pa)
  *pa[3] = 10
  *pa[1] = 40
  // NOTE: values are changed, not memory addresses
  fmt.Println(pa)
}
