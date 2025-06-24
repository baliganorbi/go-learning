package main

// Import statements bring in additional functionality.
// baliganorbi/learning/basics contaains our custom packages for basics
// fmt package provides formatting and printing functions
import (
	"baliganorbi/learning/basics"
	datastructures "baliganorbi/learning/data-structures"
	"baliganorbi/learning/functions"
	"fmt"
)

// main() is the entry point of our program.
// Every executable Go program must have a main function in package main
func main() {

	// 01_hello - Hello World example
	fmt.Println("== Basics 01_hello - Hello World example ==")
	basics.Hello()
	fmt.Println()

	// 02_data-types - Demonstrates Go's basic data types
	fmt.Println("== Basics 02_data-types - Demonstrates Go's basic data types ==")
	basics.DataTypes()
	fmt.Println()

	// 03_controls - Demonstrates control structures
	fmt.Println("== Basics 03_controls - Demonstrates control structures ==")
	basics.Controls()
	fmt.Println()

	// 01_func - Demonstrates function features
	fmt.Println("== Functions 01_func - Demonstrates function features ==")
	functions.DemoFunctions()
	fmt.Println()

	// 02_methods - Demonstrates method declarations and usage
	fmt.Println("== Functions 02_methods - Demonstrates method declarations and usage ==")
	functions.DemoMethods()
	fmt.Println()

	// 03_interfaces - Demonstrates interfaces
	fmt.Println("== Functions 03_interfaces - Demonstrates interfaces ==")
	functions.DemoInterfaces()
	fmt.Println()

	// 01_arrays - Demonstrates arrays and slices
	fmt.Println("== Data Structures 01_arrays - Demonstrates arrays and slices ==")
	datastructures.DemoArraysAndSlices()
	fmt.Println()

	// 02_maps - Demonstrates maps
	fmt.Println("== Data Structures 02_maps - Demonstrates maps ==")
	datastructures.DemoMaps()
	fmt.Println()

	// 03_structs - Demonstrates structs
	fmt.Println("== Data Structures 03_structs - Demonstrates structs ==")
	datastructures.DemoStructs()
	fmt.Println()

}
