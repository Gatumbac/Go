package main

import "fmt"

func main() {
	total := 0
	for _, n := range []int{2, 4, 6} {
		if n > 3 {
			total += n
		}
	}
	fmt.Println(total)
}
