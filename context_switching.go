package main

import (
	"fmt"
	"runtime"
	"strings"
	"sync"
	"time"
)

func main() {
	fmt.Println("👻 CONTEXT SWITCHING SÉANCE 👻")
	fmt.Println(strings.Repeat("⚡", 30))

	const pingPongs = 1_000_000 // One million soul transfers!

	// Run each experiment 3 times for averages
	var singleThreadTimes, multiThreadTimes []time.Duration

	fmt.Println("\n🕯️ EXPERIMENT 1: SINGLE THREAD POSSESSION 🕯️")
	fmt.Println("(All ghosts must share ONE body)")

	for i := 0; i < 3; i++ {
		runtime.GOMAXPROCS(1) // Force single-thread haunting
		duration := runPingPong(pingPongs, i+1)
		singleThreadTimes = append(singleThreadTimes, duration)
	}

	fmt.Println("\n💀 EXPERIMENT 2: MULTI-THREAD CHAOS 💀")
	fmt.Println("(Ghosts can possess multiple bodies)")

	for i := 0; i < 3; i++ {
		runtime.GOMAXPROCS(runtime.NumCPU()) // Unleash all cores!
		duration := runPingPong(pingPongs, i+1)
		multiThreadTimes = append(multiThreadTimes, duration)
	}

	// Calculate and display the cursed results
	displayResults(singleThreadTimes, multiThreadTimes, pingPongs)
}

func runPingPong(rounds int, attempt int) time.Duration {
	fmt.Printf("\n🔮 Attempt %d: Summoning goroutines...\n", attempt)

	// The haunted channel - unbuffered for immediate possession transfer
	ping := make(chan struct{}) // Empty struct = pure spirit energy
	pong := make(chan struct{})

	var wg sync.WaitGroup
	wg.Add(2)

	startTime := time.Now()

	// Goroutine 1: The Sexyy Red Spirit
	go func() {
		defer wg.Done()
		for i := 0; i < rounds; i++ {
			ping <- struct{}{} // Send soul
			<-pong             // Receive soul back
		}
	}()

	// Goroutine 2: The FKA twigs Spirit
	go func() {
		defer wg.Done()
		for i := 0; i < rounds; i++ {
			<-ping             // Receive soul
			pong <- struct{}{} // Send soul back
		}
	}()

	wg.Wait()
	duration := time.Since(startTime)

	fmt.Printf("✨ Completed %d ping-pongs in %v\n", rounds, duration)

	return duration
}

func displayResults(singleThread, multiThread []time.Duration, rounds int) {
	fmt.Println("\n" + strings.Repeat("🩸", 30))
	fmt.Println("\n⚰️ THE CONTEXT SWITCHING AUTOPSY ⚰️")

	// Calculate averages
	avgSingle := average(singleThread)
	avgMulti := average(multiThread)

	// Calculate per-switch time (2 switches per round-trip)
	switchTimeSingle := avgSingle / time.Duration(rounds*2)
	switchTimeMulti := avgMulti / time.Duration(rounds*2)

	fmt.Printf("\n📊 POSSESSION STATISTICS:\n")
	fmt.Printf("Single-thread average: %v total (%v per switch)\n", avgSingle, switchTimeSingle)
	fmt.Printf("Multi-thread average: %v total (%v per switch)\n", avgMulti, switchTimeMulti)

	if avgSingle < avgMulti {
		ratio := float64(avgMulti) / float64(avgSingle)
		fmt.Printf("\n🎭 PLOT TWIST: Single-thread is %.2fx FASTER!\n", ratio)
	} else {
		ratio := float64(avgSingle) / float64(avgMulti)
		fmt.Printf("\n🎭 Multi-thread is %.2fx faster!\n", ratio)
	}

	fmt.Print(`

╔════════════════════════════════════════════════════════════╗
║            💀 THE HORRIFYING TRUTH 💀                       ║
╠════════════════════════════════════════════════════════════╣
║                                                            ║
║ SINGLE THREAD (The Ethel Cain Isolation):                 ║
║ • Goroutines switch in same OS thread                     ║
║ • Just updating registers & stack pointer                 ║
║ • Like ghosts swapping within the SAME body              ║
║ • ~100-200 nanoseconds per switch                         ║
║                                                            ║
║ MULTI THREAD (The Charli XCX Chaos):                      ║
║ • Goroutines might be on DIFFERENT OS threads             ║
║ • Must sync through OS kernel sometimes                   ║
║ • Like ghosts jumping between different houses           ║
║ • Cache invalidation, memory barriers, tears              ║
║                                                            ║
║ THE WITCH'S HIERARCHY OF SWITCHING HORROR:                ║
║                                                            ║
║ 🌟 Goroutine switch (same thread): ~100ns                 ║
║    "Like Ayesha switching personalities"                  ║
║                                                            ║
║ 💀 Thread switch (same process): ~1-10μs                  ║
║    "Like Eve Brown switching between tasks"              ║
║                                                            ║
║ 🩸 Process switch: ~10-100μs                              ║
║    "Like Wu Zetian switching between pilots"             ║
║                                                            ║
║ ⚰️ Container switch: ~100μs-1ms                           ║
║    "Like switching between Monaleo verses"               ║
║                                                            ║
║ 👹 VM switch: ~1-100ms                                    ║
║    "Like switching between Chappell Roan eras"           ║
║                                                            ║
╚════════════════════════════════════════════════════════════╝

🔮 THE LESSON: Goroutines are CHEAP spirits to summon! 🔮

Single-thread is often FASTER because the ghosts never leave
their original haunted house - they just swap who's in control!

Multi-thread adds the overhead of potentially crossing the
dimensional barrier between CPU cores!
`)
}

func average(durations []time.Duration) time.Duration {
	var total time.Duration
	for _, d := range durations {
		total += d
	}
	return total / time.Duration(len(durations))
}
