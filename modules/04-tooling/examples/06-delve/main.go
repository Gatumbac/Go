package main

import "fmt"

func discount(cents int) int { result := cents - cents/10; return result }
func main()                  { fmt.Println(discount(200)) }
