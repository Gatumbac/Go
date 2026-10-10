package main

import (
	"fmt"
	"testing"
)

var sink []int

func main() {
	n := testing.AllocsPerRun(20, func() { sink = make([]int, 100) })
	fmt.Printf("allocations/run=%.1f\n", n)
}
