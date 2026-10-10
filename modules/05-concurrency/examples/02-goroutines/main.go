package main

import "fmt"

func main() { done := make(chan int); go func(n int) { done <- n * n }(5); fmt.Println(<-done) }
