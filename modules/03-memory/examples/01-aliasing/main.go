package main

import "fmt"

func main() { a := []int{3, 7}; b := a; c := append([]int(nil), a...); b[0] = 9; fmt.Println(a, b, c) }
