package main

import (
	"fmt"
	"sync"
)

func main() {
	// This map is UNPROTECTED like the Necronomicon just sitting there
	m := make(map[int]int)
	var wg sync.WaitGroup

	fmt.Println("🕯️ Attempting to summon 50 goroutines into one map...")
	fmt.Println("(This is how horror movies start)")

	for g := 0; g < 50; g++ {
		wg.Add(1)
		go func(goroutineID int) {
			defer wg.Done()
			for i := 0; i < 1000; i++ {
				// THE CURSED OPERATION
				m[goroutineID*1000+i] = i
			}
		}(g)
	}

	wg.Wait()
	fmt.Printf("✨ Somehow survived! Map length: %d\n", len(m))
}
