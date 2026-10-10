package main

import "fmt"

type Counter struct{ N int }

func (c *Counter) Add(n int) { c.N += n }
func (c Counter) Value() int { return c.N }
func main()                  { c := Counter{}; c.Add(3); fmt.Println(c.Value()) }
