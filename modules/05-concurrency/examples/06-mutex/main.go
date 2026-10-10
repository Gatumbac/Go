package main

import (
	"fmt"
	"sync"
)

type counter struct {
	mu sync.Mutex
	n  int
}

func (c *counter) add()       { c.mu.Lock(); defer c.mu.Unlock(); c.n++ }
func (c *counter) value() int { c.mu.Lock(); defer c.mu.Unlock(); return c.n }
func main() {
	var c counter
	var wg sync.WaitGroup
	for range 4 {
		wg.Go(func() {
			for range 100 {
				c.add()
			}
		})
	}
	wg.Wait()
	fmt.Println(c.value())
}
