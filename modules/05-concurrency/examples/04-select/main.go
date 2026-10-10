package main

import "fmt"

func main() {
	ch := make(chan int, 1)
	ch <- 8
	select {
	case n := <-ch:
		fmt.Println(n)
	default:
		fmt.Println("empty")
	}
	select {
	case <-ch:
		fmt.Println("data")
	default:
		fmt.Println("empty")
	}
}
