package main

import "fmt"

func value() *int { n := 42; return &n }
func main()       { p := value(); fmt.Println(*p) }
