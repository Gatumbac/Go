package main

import "fmt"

func main() {
	done := make(chan struct{})
	go func() { defer close(done) }()
	<-done
	fmt.Println("joined")
}
