package main

import (
	"fmt"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

// Regular Mutex Map
type SafeMap struct {
	mu sync.Mutex
	m  map[int]int
}

// RWMutex Map
type RWMap struct {
	mu sync.RWMutex
	m  map[int]int
}

func main() {
	fmt.Println("🔮 THE GREAT MUTEX BATTLE: A Trilogy 🔮")
	fmt.Println(strings.Repeat("=", 50))

	// Run each test 3 times and average
	regularTimes := make([]time.Duration, 3)
	rwTimes := make([]time.Duration, 3)
	syncMapTimes := make([]time.Duration, 3)

	for i := 0; i < 3; i++ {
		fmt.Printf("\n🌙 RITUAL #%d 🌙\n", i+1)

		// Test 1: Regular Mutex
		fmt.Println("\n1. REGULAR MUTEX (the overprotective parent):")
		regularMap := &SafeMap{m: make(map[int]int)}
		regularTimes[i] = testRegularMutex(regularMap)

		// Test 2: RWMutex
		fmt.Println("\n2. RWMUTEX (the smart bouncer):")
		rwMap := &RWMap{m: make(map[int]int)}
		rwTimes[i] = testRWMutex(rwMap)

		// Test 3: sync.Map
		fmt.Println("\n3. SYNC.MAP (the chaos witch):")
		syncMapTimes[i] = testSyncMap()
	}

	// Calculate and display averages
	fmt.Println("\n" + strings.Repeat("🕯️", 25))
	fmt.Println("\n✨ FINAL BATTLE RESULTS ✨")
	fmt.Printf("\n🔒 Regular Mutex Average: %v", average(regularTimes))
	fmt.Printf("\n📖 RWMutex Average: %v", average(rwTimes))
	fmt.Printf("\n🌀 sync.Map Average: %v", average(syncMapTimes))

	displayTradeoffs()
}

func testRegularMutex(safeMap *SafeMap) time.Duration {
	var wg sync.WaitGroup
	startTime := time.Now()

	// 50 writers
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

	wg.Wait()
	duration := time.Since(startTime)

	safeMap.mu.Lock()
	finalLen := len(safeMap.m)
	safeMap.mu.Unlock()

	fmt.Printf("📊 Map size: %d | ⏱️ Time: %v\n", finalLen, duration)
	return duration
}

func testRWMutex(rwMap *RWMap) time.Duration {
	var wg sync.WaitGroup
	startTime := time.Now()

	// 50 writers
	for g := 0; g < 50; g++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for i := 0; i < 1000; i++ {
				rwMap.mu.Lock()
				rwMap.m[id*1000+i] = i
				rwMap.mu.Unlock()
			}
		}(g)
	}

	wg.Wait()
	duration := time.Since(startTime)

	rwMap.mu.RLock()
	finalLen := len(rwMap.m)
	rwMap.mu.RUnlock()

	fmt.Printf("📊 Map size: %d | ⏱️ Time: %v\n", finalLen, duration)
	return duration
}

func testSyncMap() time.Duration {
	var m sync.Map
	var wg sync.WaitGroup
	startTime := time.Now()

	// 50 writers - like 50 witches casting spells simultaneously
	for g := 0; g < 50; g++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for i := 0; i < 1000; i++ {
				m.Store(id*1000+i, i) // Pre-protected chaos magic!
			}
		}(g)
	}

	wg.Wait()
	duration := time.Since(startTime)

	// Count entries using Range (the séance method)
	var count int64
	m.Range(func(key, value interface{}) bool {
		atomic.AddInt64(&count, 1)
		return true // Continue the séance
	})

	fmt.Printf("📊 Map size: %d | ⏱️ Time: %v\n", count, duration)
	return duration
}

func average(times []time.Duration) time.Duration {
	var total time.Duration
	for _, t := range times {
		total += t
	}
	return total / time.Duration(len(times))
}

func displayTradeoffs() {
	fmt.Println("\n\n" + strings.Repeat("💀", 25))
	fmt.Println("\n🩸 THE BLOOD PRICE OF EACH APPROACH 🩸")

	fmt.Println(`
╔════════════════════════════════════════════════════════════╗
║                  ⚰️ MUTEX COMPARISON ⚰️                      ║
╠════════════════════════════════════════════════════════════╣
║ Regular Mutex (The Eve Brown Approach)                     ║
║ 🎭 Speed: SLOWEST (everyone waits, even readers)          ║
║ 🛡️ Safety: MAXIMUM (one at a time, period)                ║
║ 💭 Memory: LOWEST (just one lock)                         ║
║ 🎪 Best for: Simple cases, write-heavy loads              ║
║ 👻 Horror Level: Overprotective parent in horror movie    ║
╠════════════════════════════════════════════════════════════╣
║ RWMutex (The Wu Zetian Strategy)                          ║
║ 🎭 Speed: MEDIUM (readers can party together)             ║
║ 🛡️ Safety: HIGH (smart separation)                        ║
║ 💭 Memory: LOW (slightly more than regular)               ║
║ 🎪 Best for: Read-heavy workloads                         ║
║ 👻 Horror Level: Smart final girl who actually survives   ║
╠════════════════════════════════════════════════════════════╣
║ sync.Map (The Ethel Cain Chaos Magic)                     ║
║ 🎭 Speed: FASTEST (lock-free witchcraft)                  ║
║ 🛡️ Safety: BUILT-IN (but with limitations)                ║
║ 💭 Memory: HIGHEST (duplicate storage, atomic magic)      ║
║ 🎪 Best for: Mostly static keys, few writers              ║
║ 👻 Horror Level: Possessed doll that somehow works        ║
╚════════════════════════════════════════════════════════════╝

🌙 READ-HEAVY SCENARIO PROPHECY 🌙
If reads dominated (like streaming Chappell Roan vs recording):
- Regular Mutex: Still slow (readers wait for each other) 
- RWMutex: SHINES! (multiple readers vibe together)
- sync.Map: Good but not as optimized for pure reading

💀 THE CURSED TRUTH 💀
sync.Map uses copy-on-write and atomic operations - like having
two versions of reality (read-only and dirty) that occasionally
sync up during a séance. It's optimized for:
- Keys written once, read many times
- Multiple goroutines reading disjoint key sets

But it's TERRIBLE for:
- Frequent updates to same keys
- Need to iterate over all entries often
- Memory-constrained environments

🎪 Like Sexyy Red says: "You get what you pay for!" 🎪
`)
}
