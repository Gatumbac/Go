package main

import "fmt"

func sign(n int) int {
	if n < 0 {
		return -1
	}
	if n > 0 {
		return 1
	}
	return 0
}
func main() { fmt.Println(sign(-9), sign(0), sign(9)) }
