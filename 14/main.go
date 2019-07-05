package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

func main() {
	// This example even worse with normal mutex instead of rw mutex
	rwMutexMap()
	atomicMap()
}

func rwMutexMap() {
	defer timeTrack(time.Now(), "RWMutexMap")

	rwMutexLock := sync.RWMutex{}

	m := make(map[string]int, 0)
	m["fuzzy"] = 0

	wg := &sync.WaitGroup{}
	wg.Add(100000)

	for i := 0; i < 100000; i++ {
		go func() {
			rwMutexLock.Lock()
			rand.Seed(time.Now().UnixNano())
			randInt := rand.Intn(100000)

			m["fuzzy"] = randInt
			rwMutexLock.Unlock()
		}()
	}

	go func() {
		for i := 0; i < 100000; i++ {
			rwMutexLock.RLock()
			fmt.Println(m["fuzzy"])
			wg.Done()
			rwMutexLock.RUnlock()
		}
	}()

	wg.Wait()
}

func atomicMap() {
	defer timeTrack(time.Now(), "AtomicMap")

	m := &sync.Map{}
	m.Store("fuzzy", 0)

	wg := &sync.WaitGroup{}
	wg.Add(100000)

	for i := 0; i < 100000; i++ {
		go func() {
			rand.Seed(time.Now().UnixNano())
			randInt := rand.Intn(100000)

			m.Store("fuzzy", randInt)
		}()
	}

	go func() {
		for i := 0; i < 100000; i++ {
			value, _ := m.Load("fuzzy")
			fmt.Println(value.(int))
			wg.Done()
		}
	}()

	wg.Wait()
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}
