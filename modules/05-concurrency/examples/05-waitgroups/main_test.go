package main

import (
	"sync"
	"testing"
)

func TestJoin(t *testing.T) {
	values := make([]int, 8)
	var wg sync.WaitGroup
	for i := range values {
		wg.Go(func() { values[i] = i })
	}
	wg.Wait()
	for i, n := range values {
		if n != i {
			t.Fatalf("index %d got %d", i, n)
		}
	}
}
