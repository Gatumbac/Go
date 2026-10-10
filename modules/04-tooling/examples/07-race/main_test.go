package main

import "testing"

func TestTransfer(t *testing.T) {
	ch := make(chan int)
	go func() { ch <- 17 }()
	if got := <-ch; got != 17 {
		t.Fatalf("got %d", got)
	}
}
