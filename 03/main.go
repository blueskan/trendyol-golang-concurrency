package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 3; i++ {
		go func() {
			fmt.Printf("Run: %d\n", i)
		}()
	}

	time.Sleep(1 * time.Second)
}
