package main

import "fmt"

func total(values []int) int {
	n := 0
	for _, v := range values {
		n += v
	}
	return n
}
func main() { fmt.Println(total([]int{2, 3, 4})) }
