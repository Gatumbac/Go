package main

import "fmt"

func main() {
	n := 1
	defer fmt.Println("snapshot", n)
	defer func() { fmt.Println("later", n) }()
	n = 2
	fmt.Println("body", n)
}
