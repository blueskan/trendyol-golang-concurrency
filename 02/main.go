package main

import (
	"fmt"
	"time"
)

func routine() {
	fmt.Println("I'm go routine")
}

func main() {
	// Normal function call
	routine()

	// Goroutine call
	go routine()

	time.Sleep(3 * time.Second)
}
