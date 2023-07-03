package main

import "fmt"

func main() {
	// Function to find the smallest digit in a slice

	x := []int{
		48,96,86,68,
		57,82,63,70,
		37,34,83,27,
		19,97, 9,17,
	}

	small := x[0]

	for _, value := range x {
		if value < small {
			small = value
		}
	}

	fmt.Println(small)
}