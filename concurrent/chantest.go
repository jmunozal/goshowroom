package main

import (
	"fmt"
)

func main() {

	done := make(chan string)
	defer close(done)

	go func() {
		done <- "desde func"
	}()

	t := <-done
	fmt.Printf("Cadena: %s\n", t)

}
