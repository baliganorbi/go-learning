// 03_interfaces.go demonstrates interface declarations and usage in Go, including:
// - Interface definition and implementation
// - Multiple interface implementation
// - Empty interface
// - Type assertions and type switches
// - Interface composition
//
// This file is part of the functions package, which shows different
// ways to work with interfaces in Go.
package functions

import "fmt"

// Shape is an interface that defines methods for calculating area and perimeter
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Printer is an interface for types that can be printed
type Printer interface {
	Print()
}

// Circle implements both Shape and Printer interfaces
type Circle struct {
	radius float64
}

// Circle methods implementing Shape interface
func (c Circle) Area() float64 {
	return 3.14159 * c.radius * c.radius
}

func (c Circle) Perimeter() float64 {
	return 2 * 3.14159 * c.radius
}

// Circle method implementing Printer interface
func (c Circle) Print() {
	fmt.Printf("  * Circle with radius: %.2f\n", c.radius)
}

// Rectangle implements Shape interface only
type Rectangle struct {
	width, height float64
}

// Rectangle methods implementing Shape interface
func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

// Rectangle does not implement Printer, so it won't have a Print method

// printShapeInfo demonstrates interface usage
func printShapeInfo(s Shape) {
	// Type assertion to check if the shape is also a Printer
	if p, ok := s.(Printer); ok {
		p.Print()
	} else {
		fmt.Println("  * Shape does not implement Printer interface")
	}

	fmt.Printf("    - Area: %.2f\n", s.Area())
	fmt.Printf("    - Perimeter: %.2f\n", s.Perimeter())
}

// demonstrateEmptyInterface shows how empty interfaces work
func demonstrateEmptyInterface(items ...interface{}) {
	fmt.Println("  Type switch demonstration:")
	for _, item := range items {
		switch v := item.(type) {
		case string:
			fmt.Printf("  * String: %s\n", v)
		case int:
			fmt.Printf("  * Integer: %d\n", v)
		case Shape:
			fmt.Printf("  * Shape with area: %.2f\n", v.Area())
		default:
			fmt.Printf("  * Unknown type: %T\n", v)
		}
	}
}

// DemoInterfaces runs examples of interfaces in Go
func DemoInterfaces() {
	// Create shapes
	circle := Circle{radius: 5}
	rectangle := Rectangle{width: 4, height: 6}

	fmt.Println("Shape Interface Example:")
	fmt.Println("\nCircle:")
	printShapeInfo(circle)

	fmt.Println("\nRectangle:")
	printShapeInfo(rectangle)

	// Demonstrate empty interface and type assertions
	fmt.Println("\nEmpty Interface Example:")
	demonstrateEmptyInterface("Hello", 42, circle, 3.14)

	// Demonstrate interface slice
	fmt.Println("\nInterface Slice Example:")
	shapes := []Shape{circle, rectangle}
	fmt.Println("  Iterating over shapes:")
	for _, shape := range shapes {
		if p, ok := shape.(Printer); ok {
			p.Print()
			fmt.Printf("    - Area: %.2f\n", shape.Area())
		} else {
			fmt.Println("  * Shape does not implement Printer interface")
			fmt.Printf("    - Area: %.2f\n", shape.Area())
		}
	}
}
