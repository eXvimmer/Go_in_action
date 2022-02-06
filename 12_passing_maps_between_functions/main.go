package main

import (
	"fmt"
)

func main() {
	colors := map[string]string{
		"AliceBlue":   "#f0f8ff",
		"Coral":       "#ff7F50",
		"DarkGray":    "#a9a9a9",
		"ForestGreen": "#228b22",
	}

	for k, v := range colors {
		fmt.Printf("key: %s, value: %s\n", k, v)
	}

	removeColor(colors, "Coral") // NOTE: colors is passed by reference
	// NOTE: reference types in Go: slices, maps, interfaces, functions, channels
  fmt.Println()
	for k, v := range colors {
		fmt.Printf("key: %s, value: %s\n", k, v)
	}
}

func removeColor(m map[string]string, k string) {
	delete(m, k)
}
