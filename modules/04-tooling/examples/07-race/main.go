package main

import "fmt"

func main() {
	results := make(chan int, 2)
	go func() { results <- 4 }()
	go func() { results <- 6 }()
	a, b := <-results, <-results
	fmt.Println(a + b)
}
