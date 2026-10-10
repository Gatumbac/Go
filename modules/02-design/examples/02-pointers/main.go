package main

import "fmt"

func increase(p *int) { *p += 2 }
func main()           { n := 10; increase(&n); fmt.Println(n) }
