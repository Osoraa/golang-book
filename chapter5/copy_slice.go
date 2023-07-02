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
}