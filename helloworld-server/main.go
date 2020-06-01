package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", HandlerFunction)
	http.ListenAndServe(":8080", nil)
}

func HandlerFunction(w http.ResponseWriter, r *http.Request) {
	fmt.Println("http req received")
	fmt.Fprintln(w, "Hello, world!")
}
