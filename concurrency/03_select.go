// 03_select.go demonstrates the use of the select statement in Go for handling multiple channel operations.
// It covers:
// - Basic select usage
// - select with time.After for timeouts
// - select with default case (non-blocking)
//
// This file is part of the concurrency package, which shows different
// ways to work with channels and concurrency in Go.
package concurrency

import (
	"fmt"
	"time"
)

// basicSelectExample demonstrates using select to wait on multiple channels
func basicSelectExample() {
	fmt.Println("Basic Select Example:")
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(3 * time.Second)
		ch1 <- "      -> Message from ch1"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "      -> Message from ch2"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println("  * Received:", msg1)
		case msg2 := <-ch2:
			fmt.Println("  * Received:", msg2)
		}
	}
}

// selectTimeoutExample demonstrates using select with time.After for timeouts
func selectTimeoutExample() {
	fmt.Println("\nSelect Timeout Example:")
	ch := make(chan string)
	go func() {
		time.Sleep(3 * time.Second)
		ch <- "      -> Data from goroutine"
	}()

	select {
	case msg := <-ch:
		fmt.Println("  * Received:", msg)
	case <-time.After(2 * time.Second):
		fmt.Println("  * Timeout! No message received within 2 seconds.")
	}
}

// selectDefaultExample demonstrates using select with a default case for non-blocking operations
func selectDefaultExample() {
	fmt.Println("\nSelect Default (Non-blocking) Example:")
	ch := make(chan int)
	select {
	case val := <-ch:
		fmt.Println("  * Received:", val)
	default:
		fmt.Println("  * No value received, moving on (non-blocking)")
	}
}

// DemoSelect runs all the select statement examples
func DemoSelect() {
	basicSelectExample()
	selectTimeoutExample()
	selectDefaultExample()
}
