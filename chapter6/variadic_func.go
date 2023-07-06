package main

// Exhibiting variadic functions, i.e. multiple function parameters

import "fmt"

func add(args ...int) int {
	// Variadic sum function
	total := 0

	for _, num := range args {
		total += num
	}
	return total
}

func main() {
	// Add multiple values
	fmt.Println(add(1, 2, 3))

	// Sum a slice
	slice := []int{1, 2, 3, 4}
	fmt.Println(add(slice...))
}