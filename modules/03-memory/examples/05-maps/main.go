package main

import (
	"fmt"
	"sort"
)

func main() {
	m := map[string]int{"b": 2, "a": 1}
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Println(k, m[k])
	}
}
