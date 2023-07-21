package main

// Takes a variadic param and outputs the greatest number

import "fmt"

func greatest(args ...int) int {
	// A function that finds the greatest number from a variadic param list

	res := 0

	for _, num := range args {
		if num > res {
			res = num
		}
	}

	return res
}

func main() {
	fmt.Println(greatest(1, 2, 10, 350, -1, 4, 500))

	slice := []int{1, 2, 10, 3, 4, 5}

	fmt.Println(greatest(slice...))
}