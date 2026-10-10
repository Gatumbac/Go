package main

import (
	"reflect"
	"testing"
)

func TestUpperBatch(t *testing.T) {
	got := upperBatch([]string{"z", "a"})
	if !reflect.DeepEqual(got, []string{"A", "Z"}) {
		t.Fatalf("got %v", got)
	}
	if got := upperBatch(nil); len(got) != 0 {
		t.Fatalf("empty got %v", got)
	}
}
