package main

import (
	"encoding/hex"
	"fmt"
)

func roundtrip(data []byte) ([]byte, error) { return hex.DecodeString(hex.EncodeToString(data)) }
func main()                                 { out, err := roundtrip([]byte{0, 255}); fmt.Println(out, err) }
