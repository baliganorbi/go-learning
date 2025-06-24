// 02_maps.go demonstrates the use of maps in Go, including:
// - Map declaration, initialization, and usage
// - Adding, updating, and deleting elements
// - Checking for key existence
// - Iterating over maps
// - Using maps with slices and structs
//
// This file is part of the data-structures package, which shows different
// ways to work with collections in Go.
package datastructures

import "fmt"

// DemoMaps runs examples of maps in Go
func DemoMaps() {
	fmt.Println("Maps Example:")

	// 1. Map declaration and initialization
	var m1 map[string]int = map[string]int{"apple": 5, "banana": 3}
	m2 := make(map[string]int) // using make
	m2["orange"] = 7
	m2["grape"] = 2
	fmt.Printf("  * Map m1: %v\n", m1)
	fmt.Printf("  * Map m2: %v\n", m2)

	// 2. Adding and updating elements
	m1["pear"] = 4
	m1["apple"] = 10 // update
	fmt.Printf("  * m1 after add/update: %v\n", m1)

	// 3. Deleting elements
	delete(m1, "banana")
	fmt.Printf("  * m1 after delete: %v\n", m1)

	// 4. Checking for key existence
	val, ok := m1["apple"]
	if ok {
		fmt.Printf("    - apple exists, value: %d\n", val)
	} else {
		fmt.Println("    - apple does not exist")
	}
	val, ok = m1["banana"]
	if ok {
		fmt.Printf("    - banana exists, value: %d\n", val)
	} else {
		fmt.Println("    - banana does not exist")
	}

	// 5. Iterating over a map
	fmt.Println("\n  Iterating over m1:")
	for k, v := range m1 {
		fmt.Printf("    * %s: %d\n", k, v)
	}

	// 6. Map of slices
	studentGrades := map[string][]int{
		"Alice": {90, 85, 92},
		"Bob":   {78, 80, 88},
	}
	fmt.Println("  Map of slices:")
	for name, grades := range studentGrades {
		fmt.Printf("    %s: %v\n", name, grades)
	}
}
