package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	var b bytes.Buffer
	b.Write([]byte("Hello "))
	fmt.Fprintf(&b, "world\n")
	b.WriteTo(os.Stdout)
}
