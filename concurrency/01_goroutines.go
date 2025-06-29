// 01_goroutines.go demonstrates various aspects of goroutines in Go, including:
// - Basic goroutine creation
// - Running multiple goroutines
// - Synchronizing goroutines with WaitGroup
// - Anonymous goroutines
//
// This file is part of the concurrency package, which shows different
// ways to work with goroutines and concurrency in Go.
package concurrency

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// printMessage prints a message with a delay
func printMessage(message string, delay time.Duration) {
	time.Sleep(delay)
	fmt.Println(message)
}

// DemoGoroutines runs all the goroutine examples
func DemoGoroutines() {
	fmt.Println("Basic Goroutine:")
	go printMessage("      -> Hello from a goroutine!", 1000*time.Millisecond)
	printMessage("  * Hello from main goroutine!", 0*time.Millisecond)

	time.Sleep(1500 * time.Millisecond) // Ensure the goroutine has time to run

	fmt.Println("\nMultiple Goroutines:")
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 1; i <= 3; i++ {
		delay := time.Duration(rng.Intn(4)+1) * 1000 * time.Millisecond // Random delay between 1 and 4 seconds
		go printMessage(fmt.Sprintf("  * Goroutine #%d (delay: %v)", i, delay), delay)
	}

	time.Sleep(5000 * time.Millisecond) // Give goroutines time to finish

	fmt.Println("\nGoroutines with WaitGroup:")
	var wg sync.WaitGroup
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(id int) { // anonymous goroutine
			defer wg.Done()
			printMessage(fmt.Sprintf("    -> WaitGroup Goroutine #%d", id), time.Duration(4-id)*2000*time.Millisecond)
		}(i)
	}
	wg.Wait()
	// WaitGroup ensures all goroutines finish before proceeding
	fmt.Println("  * All WaitGroup goroutines completed.")
}
