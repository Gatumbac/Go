package main

import "fmt"

func main() {
	ch := make(chan string, 2)
	go func() { ch <- "alpha" }()
	go func() { ch <- "beta" }()
	for range 2 {
		fmt.Println(<-ch)
	}
}
