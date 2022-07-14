package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	Greet(os.Stdout, "Flodie")
}

func Greet(w io.Writer, name string) {
	fmt.Fprintf(w, "Hello, %s", name)
}
