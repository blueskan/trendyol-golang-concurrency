package main

import (
	"fmt"
)

func main() {
	ch := make(chan string)

	ch <- "Trendyol"
	fmt.Println(<-ch)
}
