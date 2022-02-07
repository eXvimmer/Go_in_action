package main

import (
	"fmt"

	"github.com/eXvimmer/exporting_and_unexporting_identifires/counters"
)

func main() {
	// counter := counters.alertCounter(13) // error
	// error: cannot refer to unexported name counters.alertCounter

	counter := counters.New(13)
	fmt.Printf("Counter: %d\n", counter)
}
