package main

import (
	"context"
	"fmt"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	done := make(chan error)
	go func() { <-ctx.Done(); done <- ctx.Err() }()
	cancel()
	fmt.Println(<-done)
}
