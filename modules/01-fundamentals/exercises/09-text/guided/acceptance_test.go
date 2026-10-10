//go:build exercise

package practice

import "testing"

func TestCountRunes(t *testing.T) {
	for _, tc := range []struct {
		in   string
		want int
	}{{"", 0}, {"abc", 3}, {"ñA", 2}, {"🙂", 1}} {
		if got := CountRunes(tc.in); got != tc.want {
			t.Fatalf("input %q: got %d want %d", tc.in, got, tc.want)
		}
	}
}
