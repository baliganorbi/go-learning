// 01_hello.go demonstrates how to create a simple function in Go that
// takes a parameter and uses string formatting to print output.
//
// This file is part of the basics package, which contains fundamental
// examples of Go programming concepts.
// This is a slightly modified version of the Hello World program
// through the Hello(name) function.
package basics

// Import statements bring in additional functionality.
// fmt package provides formatting and printing functions
import "fmt"

func Hello(name string) {
	// Printf allows formatted output with placeholders:
	// %s - for strings
	fmt.Printf("Hello %s, welcome to the Go learning journey.\n", name)
}
