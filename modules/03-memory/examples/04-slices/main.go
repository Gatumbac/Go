package main

import "fmt"

func main() { a := []int{1, 2, 3, 4}; b := a[:2:2]; b = append(b, 9); b[0] = 7; fmt.Println(a, b) }
