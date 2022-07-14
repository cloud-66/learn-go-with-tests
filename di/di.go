package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	Greet(os.Stdout, "Flodie")
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(MyGreeterHandler)))
}

func Greet(w io.Writer, name string) {
	fmt.Fprintf(w, "Hello, %s", name)
}

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}
