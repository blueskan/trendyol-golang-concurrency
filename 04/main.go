package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 3; i++ {
		go func(i int) {
			fmt.Printf("Run: %d\n", i)
		}(i)
	}

	time.Sleep(1 * time.Second)
}
