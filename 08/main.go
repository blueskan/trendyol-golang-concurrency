package main

import "fmt"

func main() {
	ch := make(chan string)

	go sendOnly(ch)
	recvOnly(ch)
}

func sendOnly(ch chan<- string) {
	ch <- "trendyol"
}

func recvOnly(ch <-chan string) {
	fmt.Println(<-ch)
}

/*
func doesNotCompile(ch chan<- string) {
	fmt.Println(<-ch)
}

func doesNotCompileToo(ch <-chan string) {
	ch <- "trendyol"
}
*/
