package main

import "fmt"

func main() {
	ch := make(chan string, 2)
	ch <- "north"
	ch <- "south"
	close(ch)
	for item := range ch {
		fmt.Println(item)
	}
	v, ok := <-ch
	fmt.Printf("%q %v\n", v, ok)
}
