package main

import "fmt"

type Celsius float64

func main() {
	var count int
	const ratio = 1.5
	temp := Celsius(21.5)
	fmt.Println(count, float64(count)+ratio, temp)
}
