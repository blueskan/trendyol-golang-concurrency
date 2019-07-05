package main

import (
	"fmt"
	"time"
)

// Range loop breaks only one case, close the channel..
func main() {
	ch := make(chan string, 3)

	for i := 0; i < 3; i++ {
		ch <- fmt.Sprintf("trendyol %d", i)
	}

	// if we don't use go routine here, then guess what deadlock!
	// but why simply you don't close channel our next slide please
	go rangeOverMyChannelPlease(ch)

	time.Sleep(1 * time.Second)
}

func rangeOverMyChannelPlease(ch <-chan string) {
	for val := range ch {
		fmt.Println(val)
	}
}
