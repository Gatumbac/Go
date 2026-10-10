package main

import "fmt"

type Describer interface{ Describe() string }
type Planet string

func (p Planet) Describe() string { return "planet " + string(p) }
func show(d Describer)            { fmt.Println(d.Describe()) }
func main()                       { show(Planet("Mars")) }
