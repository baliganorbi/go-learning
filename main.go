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

// Lesson represents a single example/lesson with a title and a function to run
type Lesson struct {
	Title string
	Run   func()
}

// main() is the entry point of our program.
// Every executable Go program must have a main function in package main
func main() {
	lessons := []Lesson{
		{"== Basics 01_hello - Hello World example ==", basics.Hello},
		{"== Basics 02_data-types - Demonstrates Go's basic data types ==", basics.DataTypes},
		{"== Basics 03_controls - Demonstrates control structures ==", basics.Controls},
		{"== Functions 01_func - Demonstrates function features ==", functions.DemoFunctions},
		{"== Functions 02_methods - Demonstrates method declarations and usage ==", functions.DemoMethods},
		{"== Functions 03_interfaces - Demonstrates interfaces ==", functions.DemoInterfaces},
		{"== Data Structures 01_arrays - Demonstrates arrays and slices ==", datastructures.DemoArraysAndSlices},
		{"== Data Structures 02_maps - Demonstrates maps ==", datastructures.DemoMaps},
		{"== Data Structures 03_structs - Demonstrates structs ==", datastructures.DemoStructs},
	}

	for _, lesson := range lessons {
		fmt.Println(lesson.Title)
		lesson.Run()
		fmt.Println()
	}
}
