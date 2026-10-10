package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "ñA"
	fmt.Println(len(s), utf8.RuneCountInString(s))
	for offset, r := range s {
		fmt.Printf("%d %U\n", offset, r)
	}
}
