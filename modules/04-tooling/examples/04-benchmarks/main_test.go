package main

import "testing"

var result int

func BenchmarkTotal(b *testing.B) {
	values := make([]int, 100)
	for i := range values {
		values[i] = i
	}
	b.ReportAllocs()
	for b.Loop() {
		result = total(values)
	}
}
