package main

import "fmt"

type Label struct{ Text string }

func (l Label) Render() string { return "[" + l.Text + "]" }

type Badge struct {
	Label
	Number int
}

func main() { b := Badge{Label: Label{Text: "guest"}, Number: 7}; fmt.Println(b.Render(), b.Number) }
