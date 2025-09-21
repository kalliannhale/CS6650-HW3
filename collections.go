package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// Run the experiment 3 times
	for run := 1; run <= 3; run++ {
		fmt.Printf("Run %d: ", run)
		runExperiment()
		time.Sleep(100 * time.Millisecond) // Small delay between runs
	}
}

func runExperiment() {
	m := make(map[int]int)
	var wg sync.WaitGroup

	// Spawn 50 goroutines
	for g := 0; g < 50; g++ {
		wg.Add(1)
		go func(goroutineID int) {
			defer wg.Done()
			// Each goroutine writes 1,000 entries
			for i := 0; i < 1000; i++ {
				m[goroutineID*1000+i] = i
			}
		}(g)
	}

	// Wait for all goroutines to finish
	wg.Wait()
	fmt.Printf("len(m) = %d\n", len(m))
}
