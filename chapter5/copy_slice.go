package main

import "fmt"

func main() {
	// Create slices
	slice0 := []int{1, 2, 3}
	slice1 := make([]int, 2)

	// Copy slice zero into one, notice how copy stands alone
	// Destination(dst) comes before source(src)...weird!
	copy(slice1, slice0)

	// Only the first two elements from 0 are copied into 1
	fmt.Println("Slice0 is", slice0, "with length", len(slice0))
	fmt.Println("Slice1 is", slice1, "with length", len(slice1))

	// Create and copy slices
	slice2 := []int{1, 2, 3}
	slice3 := make([]int, 4)

	// Remember, destination slice comes before source slice
	copy(slice3, slice2)

	// All elements are copied from 2 into 3
	fmt.Println("Slice2 is", slice2, "with length", len(slice2))
	fmt.Println("Slice3 is", slice3, "with length", len(slice3))
}