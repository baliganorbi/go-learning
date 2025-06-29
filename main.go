package main

// Import statements bring in additional functionality.
// baliganorbi/learning/basics contaains our custom packages for basics
// fmt package provides formatting and printing functions
import (
	"baliganorbi/learning/basics"
	"baliganorbi/learning/concurrency"
	datastructures "baliganorbi/learning/data-structures"
	"baliganorbi/learning/functions"
	"fmt"
	"regexp"
	"strings"
)

// Extracts the short name from the lesson title using regex.
// Example: "== Basics 01_hello - Hello World example ==" -> "Basics 01_hello"
func shortNameFromTitle(title string) string {
	re := regexp.MustCompile(`==\s*(.*?) \s*-`)
	matches := re.FindStringSubmatch(title)
	if len(matches) >= 2 {
		return strings.TrimSpace(matches[1])
	}
	return title
}

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
		{"== Concurrency 01_goroutines - Demonstrates Goroutines ==", concurrency.DemoGoroutines},
		{"== Concurrency 01_url-status-check - Demonstrates URL status check using Goroutines ==", concurrency.DemoURLStatusCheck},
		{"== Concurrency 02_channels - Demonstrates Channels ==", concurrency.DemoChannels},
		{"== Concurrency 02_worker-pool - Demonstrates Worker Pool pattern ==", concurrency.DemoWorkerPool},
	}

	fmt.Println("Available Lessons:")
	for i, lesson := range lessons {
		fmt.Printf("%d. %s\n", i+1, shortNameFromTitle(lesson.Title))
	}

	var choice int
	fmt.Print("Enter the lesson number to run (0 to exit): ")
	_, err := fmt.Scan(&choice)
	if err != nil || choice < 0 || choice > len(lessons) {
		fmt.Println("Invalid input. Exiting.")
		return
	}
	if choice == 0 {
		fmt.Println("Exiting.")
		return
	}

	selected := lessons[choice-1]
	fmt.Printf("\n%s\n\n", selected.Title)
	selected.Run()
}
