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

// SafeMapRW uses RWMutex for potentially better read performance
type SafeMapRW struct {
	mu sync.RWMutex
	m  map[int]int
}

func NewSafeMapRW() *SafeMapRW {
	return &SafeMapRW{
		m: make(map[int]int),
	}
}

func (sm *SafeMapRW) Set(key, value int) {
	sm.mu.Lock()
	sm.m[key] = value
	sm.mu.Unlock()
}

func (sm *SafeMapRW) Len() int {
	sm.mu.RLock()
	defer sm.mu.RUnlock()
	return len(sm.m)
}

func main() {
	fmt.Println("=== Comparing Map Synchronization Approaches ===\n")

	// 1. Regular Mutex
	fmt.Println("1. Regular Mutex:")
	var mutexTotal time.Duration
	for run := 1; run <= 3; run++ {
		duration := runMutexExperiment(run)
		mutexTotal += duration
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Printf("Mean time: %.2fms\n\n", float64(mutexTotal.Milliseconds())/3.0)

	// 2. RWMutex
	fmt.Println("2. RWMutex:")
	var rwMutexTotal time.Duration
	for run := 1; run <= 3; run++ {
		duration := runRWMutexExperiment(run)
		rwMutexTotal += duration
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Printf("Mean time: %.2fms\n\n", float64(rwMutexTotal.Milliseconds())/3.0)

	// 3. sync.Map
	fmt.Println("3. sync.Map:")
	var syncMapTotal time.Duration
	for run := 1; run <= 3; run++ {
		duration := runSyncMapExperiment(run)
		syncMapTotal += duration
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Printf("Mean time: %.2fms\n\n", float64(syncMapTotal.Milliseconds())/3.0)

	// Single-threaded baseline
	fmt.Println("=== Single-Threaded Baseline ===")
	runSingleThreaded()
}

func runMutexExperiment(runNumber int) time.Duration {
	sm := NewSafeMap()
	var wg sync.WaitGroup

	start := time.Now()

	for g := 0; g < 50; g++ {
		wg.Add(1)
		go func(goroutineID int) {
			defer wg.Done()
			for i := 0; i < 1000; i++ {
				sm.Set(goroutineID*1000+i, i)
			}
		}(g)
	}

	wg.Wait()
	duration := time.Since(start)

	fmt.Printf("Run %d: len(m) = %d, time: %.2fms\n",
		runNumber, sm.Len(), float64(duration.Microseconds())/1000.0)

	return duration
}

func runRWMutexExperiment(runNumber int) time.Duration {
	sm := NewSafeMapRW()
	var wg sync.WaitGroup

	start := time.Now()

	for g := 0; g < 50; g++ {
		wg.Add(1)
		go func(goroutineID int) {
			defer wg.Done()
			for i := 0; i < 1000; i++ {
				sm.Set(goroutineID*1000+i, i)
			}
		}(g)
	}

	wg.Wait()
	duration := time.Since(start)

	fmt.Printf("Run %d: len(m) = %d, time: %.2fms\n",
		runNumber, sm.Len(), float64(duration.Microseconds())/1000.0)

	return duration
}

func runSyncMapExperiment(runNumber int) time.Duration {
	var m sync.Map
	var wg sync.WaitGroup

	start := time.Now()

	for g := 0; g < 50; g++ {
		wg.Add(1)
		go func(goroutineID int) {
			defer wg.Done()
			for i := 0; i < 1000; i++ {
				m.Store(goroutineID*1000+i, i)
			}
		}(g)
	}

	wg.Wait()
	duration := time.Since(start)

	// Count entries
	count := 0
	m.Range(func(_, _ interface{}) bool {
		count++
		return true
	})

	fmt.Printf("Run %d: entries = %d, time: %.2fms\n",
		runNumber, count, float64(duration.Microseconds())/1000.0)

	return duration
}

func runSingleThreaded() {
	m := make(map[int]int)

	start := time.Now()

	for g := 0; g < 50; g++ {
		for i := 0; i < 1000; i++ {
			m[g*1000+i] = i
		}
	}

	duration := time.Since(start)

	fmt.Printf("Single-threaded: len(m) = %d, time: %.2fms\n",
		len(m), float64(duration.Microseconds())/1000.0)
}
