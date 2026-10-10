package main

import "fmt"

func bounds(a, b int) (int, int) {
	if a < b {
		return a, b
	}
	return b, a
}
func main() { low, high := bounds(9, 3); fmt.Println(low, high) }
