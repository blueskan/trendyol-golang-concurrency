package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	contextTimeout()
	// contextCancel()
}

func contextTimeout() {
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		break
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}

func contextCancel() {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		time.Sleep(80 * time.Millisecond)
		cancel()
	}()

	select {
	case <-time.After(1 * time.Second):
		break
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}
