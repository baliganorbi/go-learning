// 01_arrays.go demonstrates the use of arrays and slices in Go, including:
// - Array declaration, initialization, and usage
// - Slice creation, slicing, and appending
// - Built-in functions for slices (len, cap, append, copy)
// - Iterating over arrays and slices
//
// This file is part of the data-structures package, which shows different
// ways to work with collections in Go.
package datastructures

import "fmt"

// DemoArraysAndSlices runs examples of arrays and slices in Go
func DemoArraysAndSlices() {
	fmt.Println("Arrays and Slices Example:")

	// 1. Array declaration and initialization
	var arr [3]int = [3]int{1, 2, 3}
	fmt.Printf("  Array: %v\n", arr)
	fmt.Printf("  Array length: %d\n", len(arr))

	// 2. Iterating over an array
	fmt.Println("  Iterating over array:")
	for i, v := range arr {
		fmt.Printf("    * arr[%d] = %d\n", i, v)
	}

	// 3. Slice creation from array
	slice := arr[1:]
	fmt.Printf("  Slice from array (arr[1:]): %v\n", slice)

	// 4. Slice literal and appending
	nums := []int{10, 20, 30}
	fmt.Printf("  Nums slice: %v\n", nums)
	nums = append(nums, 40)
	fmt.Printf("    * Slice after append: %v\n", nums)

	// 5. Built-in functions: len, cap, copy
	fmt.Printf("    * Slice length: %d, capacity: %d\n", len(nums), cap(nums))
	copySlice := make([]int, len(nums))
	copy(copySlice, nums)
	fmt.Printf("  Copied slice: %v\n", copySlice)

	// 6. Slicing a slice
	sub := nums[1:3]
	fmt.Printf("  Sub-slice (nums[1:3]): %v\n", sub)

	// 7. Iterating over a slice
	fmt.Println("  Iterating over slice:")
	for i, v := range nums {
		fmt.Printf("    nums[%d] = %d\n", i, v)
	}
}
