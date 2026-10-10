package main

import "testing"

func TestTwice(t *testing.T) {
	if got := twice(6); got != 12 {
		t.Fatalf("twice(6)=%d, want 12", got)
	}
}
