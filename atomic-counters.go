package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	// The protected one - like Wu Zetian's Iron Widow mech, synchronized and lethal
	var atomicOps atomic.Uint64

	// The unprotected one - like Eve Brown without her sisters, chaotic and vulnerable
	var regularOps uint64

	var wg sync.WaitGroup

	fmt.Println("🕯️ Summoning 50 goroutines to increment 1000 times each...")
	fmt.Println("Expected: 50,000 for both. Reality? *cackles in data race*\n")

	// First, the atomic ritual
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				atomicOps.Add(1) // Protected like Ethel Cain's trailer behind locked doors
			}
		}()
	}
	wg.Wait()

	// Reset for the chaos demon (regular int)
	wg = sync.WaitGroup{}

	// Now the unprotected variable - like going to investigate that noise alone
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				regularOps++ // This is giving "first to die in a horror movie" energy
			}
		}()
	}
	wg.Wait()

	fmt.Printf("⚡ Atomic ops (protected by eldritch synchronization): %d\n", atomicOps.Load())
	fmt.Printf("👻 Regular ops (raw dogging concurrency): %d\n", regularOps)
	fmt.Printf("\n💀 Data corruption level: %d missing increments\n", 50000-int(regularOps))
}
