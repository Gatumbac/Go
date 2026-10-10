package main

import (
	"bytes"
	"testing"
)

func FuzzRoundtrip(f *testing.F) {
	f.Add([]byte{})
	f.Add([]byte{0, 255})
	f.Fuzz(func(t *testing.T, in []byte) {
		out, err := roundtrip(in)
		if err != nil || !bytes.Equal(in, out) {
			t.Fatalf("roundtrip: %v", err)
		}
	})
}
