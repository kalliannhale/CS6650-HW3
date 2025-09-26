package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	fmt.Println("💀 FILE I/O HORROR SHOW 💀")
	fmt.Println(strings.Repeat("🩸", 25))

	const iterations = 100000
	lineContent := "The ghosts are writing their stories...\n"

	// Run each test 3 times for averages
	var unbufferedTimes, bufferedTimes []time.Duration

	for round := 1; round <= 3; round++ {
		fmt.Printf("\n🕯️ SÉANCE ROUND %d 🕯️\n", round)

		// Test 1: Unbuffered (Direct to Hell)
		unbufferedTime := testUnbuffered(iterations, lineContent)
		unbufferedTimes = append(unbufferedTimes, unbufferedTime)

		// Test 2: Buffered (With Protection Circle)
		bufferedTime := testBuffered(iterations, lineContent)
		bufferedTimes = append(bufferedTimes, bufferedTime)

		fmt.Printf("\n⚡ Speed Difference: %.2fx faster with buffer\n",
			float64(unbufferedTime)/float64(bufferedTime))
	}

	// Show the cursed truth
	displayResults(unbufferedTimes, bufferedTimes)
}

func testUnbuffered(iterations int, content string) time.Duration {
	fmt.Print("\n👻 UNBUFFERED WRITES (straight to the underworld)... ")

	// Open the cursed tome
	file, err := os.Create("unbuffered_horror.txt")
	if err != nil {
		panic("Failed to open portal to disk dimension!")
	}
	defer file.Close()

	startTime := time.Now()

	// Each write goes DIRECTLY to disk - like texting Sexyy Red one letter at a time
	for i := 0; i < iterations; i++ {
		file.Write([]byte(content)) // Individual trip to hell each time!
	}

	duration := time.Since(startTime)
	fmt.Printf("Complete! Time: %v\n", duration)

	return duration
}

func testBuffered(iterations int, content string) time.Duration {
	fmt.Print("\n✨ BUFFERED WRITES (collecting souls first)... ")

	// Open another cursed tome
	file, err := os.Create("buffered_magic.txt")
	if err != nil {
		panic("Failed to open portal to disk dimension!")
	}
	defer file.Close()

	// Wrap in protective buffer magic - like Eve Brown's organizational system
	writer := bufio.NewWriter(file)

	startTime := time.Now()

	// Writes go to memory first - like collecting verses before recording
	for i := 0; i < iterations; i++ {
		writer.WriteString(content) // Just adding to the buffer spell
	}

	// The actual summoning - all at once!
	writer.Flush() // Like dropping the whole Ethel Cain album at once

	duration := time.Since(startTime)
	fmt.Printf("Complete! Time: %v\n", duration)

	return duration
}

func displayResults(unbufferedTimes, bufferedTimes []time.Duration) {
	fmt.Println("\n" + strings.Repeat("💀", 25))
	fmt.Println("\n🩸 THE HORRIFYING TRUTH ABOUT DISK I/O 🩸")

	// Calculate averages
	var unbuffAvg, buffAvg time.Duration
	for i := range unbufferedTimes {
		unbuffAvg += unbufferedTimes[i]
		buffAvg += bufferedTimes[i]
	}
	unbuffAvg /= time.Duration(len(unbufferedTimes))
	buffAvg /= time.Duration(len(bufferedTimes))

	fmt.Printf("\n📊 AVERAGE TIMES:\n")
	fmt.Printf("Unbuffered: %v (like walking to hell 100,000 times)\n", unbuffAvg)
	fmt.Printf("Buffered: %v (like taking one bus to hell)\n", buffAvg)
	fmt.Printf("\nSpeed difference: %.2fx faster with buffering!\n",
		float64(unbuffAvg)/float64(buffAvg))

	fmt.Print(`
╔════════════════════════════════════════════════════════════╗
║              🔮 THE DISK I/O NIGHTMARE 🔮                   ║
╠════════════════════════════════════════════════════════════╣
║                                                            ║
║  UNBUFFERED (The Ayesha Erotica Chaos Method):           ║
║  • Each write = Full journey to physical disk             ║
║  • Like texting one letter at a time                      ║
║  • Kernel switches, disk seeks, tears & suffering         ║
║  • Good for: Literally nothing except pain                ║
║                                                            ║
║  BUFFERED (The FKA twigs Choreographed Approach):        ║
║  • Collects writes in memory (usually 4KB buffer)         ║
║  • Like writing your whole message before hitting send    ║
║  • One efficient disk journey when buffer fills/flushes   ║
║  • Good for: Not wanting to die of old age                ║
║                                                            ║
║  WHY THE HORROR?                                          ║
║  • Disk I/O is like traveling to another dimension        ║
║  • Each system call crosses the kernel boundary           ║
║  • Physical disk seeks take MILLISECONDS (eternity!)     ║
║  • Memory is ~100,000x faster than disk                   ║
║                                                            ║
║  It's like the difference between:                        ║
║  - Walking to Wu Zetian's palace 100,000 times           ║
║  - Taking one trip with a truck full of tribute          ║
║                                                            ║
╚════════════════════════════════════════════════════════════╝

🎭 THE LESSON: Batch your trips to hell, bestie! 🎭

Like Charli XCX says: "I don't wanna go to school,
I just wanna break the rules" - but the rule here is
ALWAYS USE BUFFERED I/O unless you enjoy suffering!

The buffer is giving Chappell Roan's "Pink Pony Club" -
gathering all the small-town gays before the big journey!
`)
}
