package main

import (
	"fmt"
)

func main() {
	// ex1()
	// ex2()
	// ex3() // lets deadlock great again
}

func ex1() {
	ch := make(chan string, 1)

	ch <- "Trendyol"
	fmt.Println(<-ch)
}

func ex2() {
	ch := make(chan string, 5)

	for i := 0; i < 5; i++ {
		ch <- "Trendyol"
	}

	for i := 0; i < 5; i++ {
		fmt.Println(<-ch)
	}
}

func ex3() {
	ch := make(chan string, 5)

	for i := 0; i < 5; i++ {
		ch <- "Trendyol"
	}

	for i := 0; i < 6; i++ {
		fmt.Println(<-ch)
	}
}
