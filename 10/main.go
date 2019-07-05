package main

import (
	"fmt"
	"time"
)

// Close channel for break range loop,

// But caution: Always close your channel from publisher
// because simply if one publisher publish closed channel
// occured panic in golang..

func main() {
	ch := make(chan string, 3)

	for i := 0; i < 3; i++ {
		ch <- fmt.Sprintf("trendyol %d", i)
	}

	close(ch)

	// Proof of Concept:
	// Lets publish some string over closed channel
	// ch <- "boom!!!"

	rangeOverMyChannelPlease(ch)

	time.Sleep(1 * time.Second)
}

func rangeOverMyChannelPlease(ch <-chan string) {
	for val := range ch {
		fmt.Println(val)
	}
}
