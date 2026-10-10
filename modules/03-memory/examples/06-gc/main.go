package main

import (
	"fmt"
	"runtime"
)

func main() {
	var stats runtime.MemStats
	runtime.ReadMemStats(&stats)
	fmt.Printf("heap=%d gc-cycles=%d\n", stats.HeapAlloc, stats.NumGC)
}
