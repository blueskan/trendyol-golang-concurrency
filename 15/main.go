package main

import (
	"fmt"
	"sync/atomic"
)

// Global counter, int's default value 0
var counter int64

func main() {
	atomicIncrement()
}

func atomicIncrement() {
	chWait := make(chan bool)

	for i := 0; i < 1000; i++ {
		go criticalSection(chWait)
	}

	// Just wait our main routine
	for i := 0; i < 1000; i++ {
		<-chWait
	}

	fmt.Println(counter)
}

func criticalSection(chWait chan<- bool) {
	atomic.AddInt64(&counter, 1)
	chWait <- true
}
