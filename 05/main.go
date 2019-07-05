package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	go func() {
		fmt.Println(<-ch)
	}()

	ch <- "Trendyol"

	time.Sleep(1 * time.Second)
}
