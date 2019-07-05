package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(50 * time.Millisecond)
		ch <- "hello"

		time.Sleep(70 * time.Millisecond)
		ch2 <- "hello 2"
	}()

	blockingWay(ch, ch2)
	// multiplexingWay(ch, ch2)
	// multiplexingWayWithDefault(ch, ch2)
}

func blockingWay(ch <-chan string, ch2 <-chan string) {
	// This codes blocking each other sad
	fmt.Println(<-ch)
	fmt.Println(<-ch2)
}

func multiplexingWay(ch <-chan string, ch2 <-chan string) {
	// This select statement by oneself blocking
	// but we listen 3 signals in this select statement
	// in short we just multiplexed this values

	// I put inside select to infinite loop
	// When single signal achieve then select blocking loop again
	// After timeout phase finish i break this loop

	for {
		select {
		case value := <-ch:
			fmt.Printf("Get channel 1 value: %s\n", value)
			break
		case <-ch2:
			fmt.Println("Get channel 2 signal but don't interest with value")
			break
		case <-time.After(2 * time.Second): // Return channel
			fmt.Println("timeout")
			return
		}
	}
}

func multiplexingWayWithDefault(ch <-chan string, ch2 <-chan string) {
	// in this example i should create here because
	// everytime for loops run, it should create new timeout
	timeout := time.After(500 * time.Millisecond)

	for {
		select {
		case value := <-ch:
			fmt.Printf("Get channel 1 value: %s\n", value)
			break
		case <-ch2:
			fmt.Println("Get channel 2 signal but don't interest with value")
			break
		case <-timeout: // Return channel
			fmt.Println("timeout")
			return
		default:
			time.Sleep(35 * time.Millisecond) // Prevent stdout flooding
			fmt.Println("Nothing new")
		}
	}
}
