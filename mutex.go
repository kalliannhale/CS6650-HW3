package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

// Regular Mutex Map - EXCLUSIVE access only
type SafeMap struct {
	mu sync.Mutex
	m  map[int]int
}

// RWMutex Map - Multiple readers OR one writer
type RWMap struct {
	mu sync.RWMutex // The dual-nature spell!
	m  map[int]int
}

func main() {
	fmt.Println("⚔️ MUTEX BATTLE: Regular vs RWMutex ⚔️")
	fmt.Println(strings.Repeat("=", 50)) // Fixed it like Eve Brown would!

	// ROUND 1: Regular Mutex
	fmt.Println("\n🔮 REGULAR MUTEX (everyone waits their turn):")
	regularMap := &SafeMap{m: make(map[int]int)}
	testRegularMutex(regularMap)

	// ROUND 2: RWMutex
	fmt.Println("\n✨ RWMUTEX (multiple readers allowed):")
	rwMap := &RWMap{m: make(map[int]int)}
	testRWMutex(rwMap)
}

func testRegularMutex(safeMap *SafeMap) {
	var wg sync.WaitGroup
	startTime := time.Now()

	// Writers - like Ethel Cain recording vocals
	for g := 0; g < 50; g++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for i := 0; i < 1000; i++ {
				safeMap.mu.Lock()
				safeMap.m[id*1000+i] = i
				safeMap.mu.Unlock()
			}
		}(g)
	}

	// Readers - like fans trying to stream the album
	for r := 0; r < 20; r++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 100; i++ {
				safeMap.mu.Lock()
				_ = len(safeMap.m) // Just checking the vibe
				safeMap.mu.Unlock()
				time.Sleep(time.Microsecond) // Small pause between reads
			}
		}()
	}

	wg.Wait()
	writeTime := time.Since(startTime)

	safeMap.mu.Lock()
	finalLen := len(safeMap.m)
	safeMap.mu.Unlock()

	fmt.Printf("📊 Final map size: %d\n", finalLen)
	fmt.Printf("⏱️ Total time: %v\n", writeTime)
}

func testRWMutex(rwMap *RWMap) {
	var wg sync.WaitGroup
	startTime := time.Now()

	// Writers - like Sexyy Red dropping exclusive content
	for g := 0; g < 50; g++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for i := 0; i < 1000; i++ {
				rwMap.mu.Lock() // EXCLUSIVE writer lock
				rwMap.m[id*1000+i] = i
				rwMap.mu.Unlock()
			}
		}(g)
	}

	// Readers - like the girlies sharing the tea simultaneously
	for r := 0; r < 20; r++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 100; i++ {
				rwMap.mu.RLock() // SHARED reader lock!
				_ = len(rwMap.m) // Multiple readers at once!
				rwMap.mu.RUnlock()
				time.Sleep(time.Microsecond)
			}
		}()
	}

	wg.Wait()
	writeTime := time.Since(startTime)

	rwMap.mu.RLock()
	finalLen := len(rwMap.m)
	rwMap.mu.RUnlock()

	fmt.Printf("📊 Final map size: %d\n", finalLen)
	fmt.Printf("⏱️ Total time: %v\n", writeTime)
}
