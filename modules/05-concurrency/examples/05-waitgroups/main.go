package main

import (
	"fmt"
	"sync"
)

func main() {
	values := make([]int, 3)
	var wg sync.WaitGroup
	for i := range values {
		wg.Go(func() { values[i] = i + 1 })
	}
	wg.Wait()
	fmt.Println(values)
}
