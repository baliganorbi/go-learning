package main

// Import statements bring in additional functionality.
// baliganorbi/learning/basics contaains our custom packages for basics
// fmt package provides formatting and printing functions
import (
	"baliganorbi/learning/basics"
	"baliganorbi/learning/functions"
	"fmt"
)

// main() is the entry point of our program.
// Every executable Go program must have a main function in package main
func main() {

	// 01_hello - Hello World example
	fmt.Println("== 01_hello - Hello World example ==")
	basics.Hello()
	fmt.Println()

	// 02_data-types - Demonstrates Go's basic data types
	fmt.Println("== 02_data-types - Demonstrates Go's basic data types ==")
	basics.DataTypes()
	fmt.Println()

	// 03_controls - Demonstrates control structures
	fmt.Println("== 03_controls - Demonstrates control structures ==")
	basics.Controls()
	fmt.Println()

	// 01_func - Demonstrates function features
	fmt.Println("== 01_func - Demonstrates function features ==")
	functions.DemoFunctions()
	fmt.Println()
}
