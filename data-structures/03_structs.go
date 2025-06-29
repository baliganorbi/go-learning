// 03_structs.go demonstrates the use of structs in Go, including:
// - Struct declaration and initialization
// - Accessing and updating struct fields
// - Structs with methods
// - Anonymous structs
// - Struct embedding (composition)
//
// This file is part of the data-structures package, which shows different
// ways to work with collections and custom data types in Go.
package datastructures

import "fmt"

// Person is a basic struct with fields
type Person struct {
	Name string
	Age  int
}

// Method on Person
func (p *Person) Greet() {
	fmt.Printf("  Hello, my name is %s and I am %d years old.\n", p.Name, p.Age)
}

// Employee embeds Person (composition)
type Employee struct {
	Person
	Position string
}

// DemoStructs runs examples of structs in Go
func DemoStructs() {
	fmt.Println("Structs Example:")

	// 1. Struct declaration and initialization
	fmt.Println("  Persons:")
	alice := Person{Name: "Alice", Age: 30}
	bob := Person{"Bob", 25} // positional fields
	fmt.Printf("    * alice: %+v\n", alice)
	fmt.Printf("    * bob:   %+v\n", bob)

	// 2. Accessing and updating fields
	bob.Age = 26
	fmt.Printf("      -> bob after birthday: %+v\n", bob)

	// 3. Calling methods on structs
	fmt.Println("  Calling methods:")
	alice.Greet()
	bob.Greet()

	// 4. Anonymous struct
	car := struct {
		Brand string
		Year  int
	}{Brand: "Toyota", Year: 2020}
	fmt.Printf("  Anonymous car struct: %+v\n", car)

	// 5. Struct embedding (composition)
	e := Employee{Person: alice, Position: "Developer"}
	fmt.Printf("  Employee: %+v\n", e)
	// Access embedded fields directly
	e.Greet() // Calls Person's Greet method
	fmt.Printf("    * %s is a %s\n", e.Name, e.Position)
}
