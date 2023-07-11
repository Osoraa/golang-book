package main

// Takes a variadic param and outputs the greatest number

import "fmt"

func greatest(args ...int) int {
	// Comment

	greatest := 0
	for _, num := range args {
		if num > greatest {
			greatest = num
		}
	}

	return greatest
}

func main() {
	// slice := []int{1, 2, 10, 3, 4, 5}

	fmt.Println(greatest(1, 2, 10, 3, 4, 5))
}