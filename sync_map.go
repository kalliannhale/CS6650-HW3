package main

import (
	"fmt"
	"sync"
	"time"
)

// SafeMap wraps a map with a mutex for thread-safe access
type SafeMap struct {
	mu sync.Mutex
	m  map[int]int
}

// NewSafeMap creates a new thread-safe map
func NewSafeMap() *SafeMap {
	return &SafeMap{
		m: make(map[int]int),
	}
}

// Set safely writes a key-value pair to the map
func (sm *SafeMap) Set(key, value int) {
	sm.mu.Lock()
	sm.m[key] = value
	sm.mu.Unlock()
}

// Len safely returns the length of the map
func (sm *SafeMap) Len() int {
	sm.mu.Lock()
	defer sm.mu.Unlock()
	return len(sm.m)
}

func main() {
	fmt.Println("=== Mutex-Protected Map Experiment ===\n")

	// Run the safe version 3 times
	var totalTime time.Duration
	for run := 1; run <= 3; run++ {
		duration := runSafeExperiment(run)
		totalTime += duration
		time.Sleep(100 * time.Millisecond)
	}

	fmt.Printf("\nMean time: %.2fms\n", float64(totalTime.Milliseconds())/3.0)

	// Compare with single-threaded version
	fmt.Println("\n=== Single-Threaded Comparison ===")
	runSingleThreaded()

	// Bonus: Show what happens without mutex (commented out to avoid crash)
	// fmt.Println("\n=== Unsafe Version (will crash) ===")
	// runUnsafeExperiment()
}

func runSafeExperiment(runNumber int) time.Duration {
	sm := NewSafeMap()
	var wg sync.WaitGroup

	start := time.Now()

	// Spawn 50 goroutines
	for g := 0; g < 50; g++ {
		wg.Add(1)
		go func(goroutineID int) {
			defer wg.Done()
			// Each goroutine writes 1,000 entries
			for i := 0; i < 1000; i++ {
				sm.Set(goroutineID*1000+i, i)
			}
		}(g)
	}

	// Wait for all goroutines to finish
	wg.Wait()
	duration := time.Since(start)

	fmt.Printf("Run %d: len(m) = %d, time: %.2fms\n",
		runNumber, sm.Len(), float64(duration.Microseconds())/1000.0)

	return duration
}

func runSingleThreaded() {
	m := make(map[int]int)

	start := time.Now()

	// Single thread writes all 50,000 entries
	for g := 0; g < 50; g++ {
		for i := 0; i < 1000; i++ {
			m[g*1000+i] = i
		}
	}

	duration := time.Since(start)

	fmt.Printf("Single-threaded: len(m) = %d, time: %.2fms\n",
		len(m), float64(duration.Microseconds())/1000.0)
}

func runMixedWorkload() {
	sm := NewSafeMapRW()
	var wg sync.WaitGroup

	// First, populate the map
	for i := 0; i < 10000; i++ {
		sm.Set(i, i)
	}

	start := time.Now()

	// 10 writer goroutines
	for w := 0; w < 10; w++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for i := 0; i < 100; i++ {
				sm.Set(id*1000+i, i)
			}
		}(w)
	}

	// 40 reader goroutines
	for r := 0; r < 40; r++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 1000; i++ {
				_ = sm.Len() // Simulate read operations
			}
		}()
	}

	wg.Wait()
	fmt.Printf("Mixed workload time: %.2fms\n",
		float64(time.Since(start).Microseconds())/1000.0)
}

// Uncomment this to see the unsafe version crash
/*
func runUnsafeExperiment() {
	m := make(map[int]int)
	var wg sync.WaitGroup

	for g := 0; g < 50; g++ {
		wg.Add(1)
		go func(goroutineID int) {
			defer wg.Done()
			for i := 0; i < 1000; i++ {
				m[goroutineID*1000+i] = i
			}
		}(g)
	}

	wg.Wait()
	fmt.Printf("len(m) = %d\n", len(m))
}
*/
