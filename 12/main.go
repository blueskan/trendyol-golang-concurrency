package main

import (
	"fmt"
	"sync"
)

// Global counter, int's default value 0
var counter int

func main() {
	withoutMutex()
	// withMutex()
}

func withoutMutex() {
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
	counter++
	chWait <- true
}

func withMutex() {
	chWait := make(chan bool)
	var mutex = &sync.Mutex{}

	for i := 0; i < 1000; i++ {
		go criticalSectionWithMutex(chWait, mutex)
	}

	// Just wait our main routine
	for i := 0; i < 1000; i++ {
		<-chWait
	}

	fmt.Println(counter)
}

func criticalSectionWithMutex(chWait chan<- bool, mutex *sync.Mutex) {
	mutex.Lock()
	counter++
	chWait <- true
	mutex.Unlock()
}
