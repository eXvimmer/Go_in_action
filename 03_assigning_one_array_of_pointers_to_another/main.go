package main

import (
	"fmt"
)

func main() {
	a := [3]*string{}
	b := [3]*string{new(string), new(string), new(string)}
	*b[0] = "Mustafa"
	*b[1] = "Malena"
	*b[2] = "Rachel"
	fmt.Println(a, b)
	a = b // NOTE: copy values from b to a
	fmt.Println(a, b)
}
