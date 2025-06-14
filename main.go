package main

// Import statements bring in additional functionality.
// baliganorbi/learning/basics contaains our custom packages for basics
// fmt package provides formatting and printing functions
import (
	"baliganorbi/learning/basics"
)

// main() is the entry point of our program.
// Every executable Go program must have a main function in package main
func main() {

	// 01_hello - Hello World example
	basics.Hello()

	// 02_data-types - Demonstrates Go's basic data types
	basics.DataTypes()
}
