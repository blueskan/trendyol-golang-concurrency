package main

import (
	"fmt"
)

func routine() {
	fmt.Println("I'm go routine")
}

func main() {
	// Normal function call
	routine()

	// Goroutine call
	go routine()
}
