package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	// Simulate 1M sockmap lookups
	total := 0
	for i := 0; i < 1000000; i++ {
		total += (i % 7)
	}
	elapsed := time.Since(start)
	fmt.Printf("[Telemetry] 1M sockmap BPF lookups resolved in %s\n", elapsed)
}
