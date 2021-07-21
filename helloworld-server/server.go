package main

import (
	"fmt"
	"net/http"
	"os"
)

func getHostname() string {
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "ERROR"
	}
	return hostname
}

func main() {

	fmt.Println("starting server...")
	http.HandleFunc("/", HandlerFunction)
	http.ListenAndServe(":8080", nil)
	fmt.Println("server ends")
}

func HandlerFunction(w http.ResponseWriter, r *http.Request) {
	fmt.Println("http req received")
	fmt.Fprintln(w, "hostname: ", getHostname())
}
