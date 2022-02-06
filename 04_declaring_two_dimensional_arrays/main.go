package main

import (
  "fmt"
)

func main() {
  var a [4][2]int
  fmt.Println(a)
  a = [4][2]int{{1, 10}, {2, 20}, {3, 30}, {4, 40}}
  fmt.Println(a)
  a = [4][2]int{1: {1, 10}, 3: {4, 40}}
  fmt.Println(a)
  a = [4][2]int{1: {0: 20}, 3: {1: 41}}
  fmt.Println(a)
  a[0][0] = 1
  a[0][1] = 10
  a[1][1] = 25
  a[2][1] = 30
  a[3][0] = 49
  fmt.Println(a)
}
