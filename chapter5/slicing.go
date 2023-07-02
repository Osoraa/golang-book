package main

import "fmt"

func main() {
	var slice []float64		// Creating a slice

	// Using make to create a slice.
	//The slice below is associated with a float64 array of length 5.
	slic := make([]float64, 5)

	// Specifying a third parameter in make
	arr := make([]int, 5, 10)

	// Using the [low:high] expression to create slices
	arra := [5]int{1, 2, 3, 4, 5}
	sli := arra[0:2]

	fmt.Println(sli)
}