package main

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	var c counter
	var wg sync.WaitGroup
	for range 8 {
		wg.Go(func() {
			for range 100 {
				c.add()
			}
		})
	}
	wg.Wait()
	if got := c.value(); got != 800 {
		t.Fatalf("got %d want 800", got)
	}
}
