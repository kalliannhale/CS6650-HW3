package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// runWithAtomicInt demonstrates the correct, synchronized way using atomic operations.
func runWithAtomicInt() {
	var ops atomic.Uint64
	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for c := 0; c < 1000; c++ {
				ops.Add(1)
			}
		}()
	}
	wg.Wait()
	fmt.Println("Atomic ops:", ops.Load())
}

// runWithRegularInt demonstrates the incorrect, unsynchronized way, prone to race conditions.
func runWithRegularInt() {
	var ops uint64
	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for c := 0; c < 1000; c++ {
				// This is the problematic line!
				ops++
			}
		}()
	}
	wg.Wait()
	fmt.Println("Regular ops:", ops)
}

func main() {
	fmt.Println("--- Running experiment with Atomic Integer ---")
	runWithAtomicInt()

	fmt.Println("\n--- Running experiment with Regular Integer (potential race condition) ---")
	runWithRegularInt()
}
