package main

import (
	"errors"
	"fmt"
)

var errMissing = errors.New("missing data")

func load() error { return fmt.Errorf("load config: %w", errMissing) }
func main()       { err := load(); fmt.Println(errors.Is(err, errMissing)) }
