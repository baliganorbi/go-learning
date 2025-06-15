// 01_func.go demonstrates various aspects of functions in Go, including:
// - Basic function declaration and parameters
// - Multiple return values
// - Named return values
// - Variadic functions
// - Functions as values (first-class citizens)
// - Closures and anonymous functions
//
// This file is part of the functions package, which shows different
// ways to work with functions in Go.
package functions

import "fmt"

// BasicFunction demonstrates a simple function with parameters and a return value
func BasicFunction(x, y int) int {
	return x + y
}

// MultipleReturns demonstrates a function that returns multiple values
func MultipleReturns(name string, age int) (string, bool) {
	isAdult := age >= 18
	greeting := fmt.Sprintf("Hello, %s!", name)
	return greeting, isAdult
}

// NamedReturns demonstrates named return values
func NamedReturns(radius float64) (area, circumference float64) {
	const pi = 3.14159
	area = pi * radius * radius
	circumference = 2 * pi * radius
	return // naked return using named return values
}

// VariadicFunction demonstrates a function that takes variable number of arguments
func VariadicFunction(prefix string, numbers ...int) []string {
	result := make([]string, len(numbers))
	for i, num := range numbers {
		result[i] = fmt.Sprintf("%s-%d", prefix, num)
	}
	return result
}

// HigherOrderFunction demonstrates returning a function
func HigherOrderFunction(factor int) func(int) int {
	// Returns a closure that multiplies its argument by factor
	return func(x int) int {
		return x * factor
	}
}

// DemoFunctions runs all the function examples
func DemoFunctions() {
	// 1. Basic Function
	fmt.Println("Basic Function:")
	result := BasicFunction(5, 3)
	fmt.Printf("  * 5 + 3 = %d\n", result)

	// 2. Multiple Return Values
	fmt.Println("\nMultiple Return Values:")
	greeting, isAdult := MultipleReturns("Alice", 20)
	fmt.Printf("  * %s\n  Is adult? %v\n", greeting, isAdult)

	// 3. Named Return Values
	fmt.Println("\nNamed Return Values:")
	area, circumference := NamedReturns(5.0)
	fmt.Printf("  Circle with radius 5.0:\n")
	fmt.Printf("  * Area: %.2f\n", area)
	fmt.Printf("  * Circumference: %.2f\n", circumference)

	// 4. Variadic Function
	fmt.Println("\nVariadic Function:")
	items := VariadicFunction("ID", 1, 2, 3)
	fmt.Println("  Generated IDs:")
	for _, item := range items {
		fmt.Printf("  * %s\n", item)
	}

	// 5. Higher-Order Function and Closure
	fmt.Println("\nHigher-Order Function and Closure:")
	doubler := HigherOrderFunction(2)
	tripler := HigherOrderFunction(3)
	fmt.Printf("  * Doubling 5: %d\n", doubler(5))
	fmt.Printf("  * Tripling 5: %d\n", tripler(5))

	// 6. Anonymous Function
	fmt.Println("\nAnonymous Function:")
	func(message string) {
		fmt.Printf("  * Message: %s\n", message)
	}("Hello from anonymous function!")
}
