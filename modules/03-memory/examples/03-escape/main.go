package main

import "fmt"

func newLabel() *string { s := "sample"; return &s }
func main()             { fmt.Println(*newLabel()) }
