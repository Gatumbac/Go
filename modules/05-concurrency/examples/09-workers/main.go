package main

import (
	"fmt"
	"sort"
	"strings"
	"sync"
)

func upperBatch(input []string) []string {
	jobs := make(chan string, 2)
	results := make(chan string)
	var wg sync.WaitGroup
	for range 2 {
		wg.Go(func() {
			for s := range jobs {
				results <- strings.ToUpper(s)
			}
		})
	}
	go func() {
		defer close(jobs)
		for _, s := range input {
			jobs <- s
		}
	}()
	go func() { wg.Wait(); close(results) }()
	out := []string{}
	for s := range results {
		out = append(out, s)
	}
	sort.Strings(out)
	return out
}
func main() { fmt.Println(upperBatch([]string{"b", "a", "c"})) }
