package main

import "fmt"

func same[T comparable](a, b T) bool { return a == b }
func main()                          { fmt.Println(same(3, 3), same("a", "b")) }
