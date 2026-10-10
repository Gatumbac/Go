package main

import "testing"

func TestSign(t *testing.T) {
	cases := []struct {
		name     string
		in, want int
	}{{"negative", -2, -1}, {"zero", 0, 0}, {"positive", 3, 1}}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := sign(tc.in); got != tc.want {
				t.Fatalf("sign(%d)=%d want %d", tc.in, got, tc.want)
			}
		})
	}
}
