package main

import (
	"fmt"
	"sync"
)

// Remember previous example we use dummy channel for waiting
// But there is more elegant solution for this situation
// sync.WaitGroup
func main() {
	wg := &sync.WaitGroup{}

	wg.Add(100)

	for i := 0; i < 100; i++ {
		go printTrendyol(wg)
	}

	fmt.Println("\n\nWaiting\n\n")
	wg.Wait() // This do magic
	fmt.Println("\n\nDone")
}

func printTrendyol(wg *sync.WaitGroup) {
	fmt.Println("Trendyol")
	wg.Done()
}
