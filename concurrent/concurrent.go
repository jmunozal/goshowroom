package main

import (
	"fmt"
	"time"
)

func main() {
	go spinner(100 * time.Millisecond)
	time.Sleep(3 * time.Second)
	fmt.Printf("\r \n")
}

func spinner(delay time.Duration) {
	for  {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}