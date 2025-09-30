package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(MyGreetingHandler)))
}

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func MyGreetingHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}
