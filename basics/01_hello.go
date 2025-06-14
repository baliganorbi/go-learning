// 01_hello.go is a slightly modified version of the Hello World program.
// Reads input from the user (stdio) and prints a greeting message (stdout).
//
// This file is part of the basics package, which contains fundamental
// examples of Go programming concepts.
// Hello() - A simple function that greets the user by name.
package basics

// Import statements bring in additional functionality.
// fmt package provides formatting and printing functions
import "fmt"

func Hello() {
	var name string

	// Ask the user for their name
	fmt.Print("Enter your name: ")
	fmt.Scanln(&name) // Read user input into the name variable
	// Printf allows formatted output with placeholders:
	// %s - for strings
	fmt.Printf("Hello %s, welcome to the Go learning journey.\n", name)
}
