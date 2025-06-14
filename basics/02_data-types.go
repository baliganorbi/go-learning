// 02_data-types.go demonstrates the fundamental data types available in Go
// and how to work with variables of different types. This example shows:
// - Declaration of variables with explicit types
// - Basic Go data types (numeric, string, boolean, complex)
// - Group declaration of variables using var blocks
//
// This file is part of the basics package, which contains fundamental
// examples of Go programming concepts.
// DataTypes() - A function that demonstrates Go's basic data types.
package basics

import "fmt"

func DataTypes() {
	// Go has several built-in data types, including:
	// - int, float64, string, bool
	// - complex numbers: complex64, complex128

	// Example of declaring variables with different data types
	var (
		a int        = 42
		b float64    = 3.14
		c string     = "Hello, Go!"
		d bool       = true
		e complex128 = 1 + 2i
	)

	// Print the values and types of the variables
	fmt.Printf("Integer: %d, Type: %T\n", a, a)
	fmt.Printf("Float: %f, Type: %T\n", b, b)
	fmt.Printf("String: %s, Type: %T\n", c, c)
	fmt.Printf("Boolean: %t, Type: %T\n", d, d)
	fmt.Printf("Complex: %v, Type: %T\n", e, e)
}
