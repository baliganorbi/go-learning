// 02_channels.go demonstrates various aspects of channels in Go, including:
// - Basic channel creation and usage
// - Sending and receiving values
// - Buffered channels
// - Channel closing and range
// - Synchronization with channels
//
// This file is part of the concurrency package, which shows different
// ways to work with channels and concurrency in Go.
package concurrency

import (
	"fmt"
	"time"
)

// basicChannelExample demonstrates sending and receiving using an unbuffered channel
func basicChannelExample() {
	fmt.Println("Basic Channel Example:")
	ch := make(chan string)
	go func() {
		time.Sleep(3000 * time.Millisecond)
		ch <- "      -> Hello from goroutine via channel!"
	}()
	// Receive from the channel
	fmt.Println("  * Waiting for message from goroutine...")
	msg := <-ch
	fmt.Println(msg)
}

// bufferedChannelExample demonstrates buffered channels
func bufferedChannelExample() {
	fmt.Println("\nBuffered Channel Example:")
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Printf("  * Received: %d\n", <-ch)
	fmt.Printf("  * Received: %d\n", <-ch)
}

// channelRangeExample demonstrates closing a channel and ranging over it
func channelRangeExample() {
	fmt.Println("\nChannel Range Example:")
	ch := make(chan string)
	go func() {
		for _, word := range []string{"Go", "channels", "are", "cool"} {
			ch <- word
			time.Sleep(1500 * time.Millisecond)
		}
		close(ch)
	}()
	fmt.Println("  * Receiving messages from channel:")
	// Range over the channel to receive messages until it's closed
	// This will block until the channel is closed
	// and all messages are received
	for msg := range ch {
		fmt.Printf("      -> Received: %s\n", msg)
	}
}

// syncWithChannel demonstrates using a channel for synchronization
func syncWithChannel() {
	fmt.Println("\nChannel Synchronization Example:")
	done := make(chan bool)
	go func() {
		fmt.Println("     -> Working in goroutine...")
		time.Sleep(4000 * time.Millisecond)
		done <- true
	}()
	// Wait for the goroutine to signal completion
	fmt.Println("  * Waiting for goroutine to finish...")
	<-done
	fmt.Println("  * Goroutine completed, main continues.")
}

// DemoChannels runs all the channel examples
func DemoChannels() {
	basicChannelExample()
	bufferedChannelExample()
	channelRangeExample()
	syncWithChannel()
}
