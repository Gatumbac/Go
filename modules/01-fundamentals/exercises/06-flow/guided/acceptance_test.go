//go:build exercise

package practice

import "testing"

func TestCountPositive(t *testing.T) {
	for _, tc := range []struct {
		in   []int
		want int
	}{{nil, 0}, {[]int{-2, 0, 4, 7}, 2}, {[]int{0, 0}, 0}} {
		if got := CountPositive(tc.in); got != tc.want {
			t.Fatalf("input %v: got %d want %d", tc.in, got, tc.want)
		}
	}
}
