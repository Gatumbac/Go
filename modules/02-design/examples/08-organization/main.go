package main

import (
	"fmt"
	"io"
	"os"
)

func announce(w io.Writer, text string) error { _, err := fmt.Fprintln(w, text); return err }
func main() {
	if err := announce(os.Stdout, "ready"); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
