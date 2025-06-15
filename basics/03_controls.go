// 03_controls.go demonstrates Go's control structures including
// if statements, for loops, and switch statements. This example shows:
// - Different types of for loops (standard, while-like, infinite)
// - If statements with initialization
// - Switch statements with and without expressions
//
// This file is part of the basics package, which contains fundamental
// examples of Go programming concepts.
// Controls() - A function that demonstrates Go's control structures.
package basics

import (
	"fmt"
)

func Controls() {
	// 1. Standard for loop
	fmt.Println("Standard for loop (1 to 3):")
	for i := 1; i <= 3; i++ {
		fmt.Printf("  Count: %d\n", i)
	}

	// 2. For loop as while
	fmt.Println("\nWhile-style loop (count down from 3):")
	count := 3
	for count > 0 {
		fmt.Printf("  * Count: %d\n", count)
		count--
	}

	// 3. For loop with range
	fmt.Println("\nFor-range loop over slice:")
	fruits := []string{"apple", "banana", "orange"}
	for index, value := range fruits {
		fmt.Printf("  * fruits[%d] = %s\n", index, value)
	}

	// 4. Basic if
	fmt.Println("\nBasic if statement:")
	number := 7
	if number%2 == 0 {
		fmt.Printf("  %d is even\n", number)
	} else {
		fmt.Printf("  %d is odd\n", number)
	}

	// 5. If with initialization
	fmt.Println("\nIf statement with initialization:")
	if score := 85; score >= 70 {
		fmt.Printf("  Score %d: Passed!\n", score)
	} else {
		fmt.Printf("  Score %d: Failed\n", score)
	}

	// 6. Basic switch
	fmt.Println("\nBasic switch statement:")
	day := "Wednesday"
	fmt.Printf("  Today is %s\n", day)
	switch day {
	case "Monday":
		fmt.Println("  * Start of work week")
	case "Wednesday":
		fmt.Println("  * Middle of work week")
	case "Friday":
		fmt.Println("  * End of work week")
	default:
		fmt.Println("  * Another day")
	}

	// 7. Switch with condition
	age := 18
	fmt.Println("\nSwitch with condition:")
	fmt.Printf("  Age: %d\n", age)
	switch {
	case age < 13:
		fmt.Println("  * Child")
	case age < 20:
		fmt.Println("  * Teenager")
	case age < 65:
		fmt.Println("  * Adult")
	default:
		fmt.Println("  * Senior")
	}

	// 8. Switch with type
	fmt.Println("\nSwitch with type:")
	myType := func(i interface{}) {
		switch t := i.(type) {
		case int:
			fmt.Printf("  * Integer: %d\n", t)
		case string:
			fmt.Printf("  * String: %s\n", t)
		case bool:
			fmt.Printf("  * Boolean: %t\n", t)
		default:
			fmt.Printf("  * Unknown type: %T\n", t)
		}
	}
	myType(42)
	myType("Hello")
	myType(true)
	myType(3.14)
}
