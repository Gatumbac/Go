//go:build exercise

package practice

import "testing"

func TestContains(t *testing.T) {
	if !Contains([]int{2, 4}, 4) || Contains([]int{2, 4}, 3) || Contains([]int{}, 1) {
		t.Fatal("int contract")
	}
	if !Contains([]string{"a"}, "a") || Contains([]string{"a"}, "b") {
		t.Fatal("string contract")
	}
}
