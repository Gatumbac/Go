package main

import "fmt"

func greeting(name string) string { return "Welcome, " + name }
func main()                       { fmt.Println(greeting("Ada")) }
