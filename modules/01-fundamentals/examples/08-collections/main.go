package main

import "fmt"

func main() {
	a := [3]int{4, 5, 6}
	b := a
	b[0] = 99
	s := a[:2]
	s[1] = 8
	counts := map[string]int{"red": 2}
	_, found := counts["blue"]
	fmt.Println(a, b, s, found)
}
