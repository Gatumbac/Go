package main

import "fmt"

type Book struct {
	Title string
	Pages int
}

func main() {
	a := Book{Title: "Systems", Pages: 120}
	b := a
	b.Pages = 90
	fmt.Println(a.Pages, b.Pages)
}
