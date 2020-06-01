package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("starting server...")
	http.HandleFunc("/", HandlerFunction)
	http.ListenAndServe(":8080", nil)
	fmt.Println("server ends")
}

func HandlerFunction(w http.ResponseWriter, r *http.Request) {
	fmt.Println("http req received")
	fmt.Fprintln(w, "Hello, world!")
}
