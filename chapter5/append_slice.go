package main

import "fmt"

func main() {
	slice0 := []int{0, 1, 2}

	// Appending to a slice
	slice1 := append(slice0, 3, 4, 5)

	fmt.Println(slice0, len(slice0), slice1, len(slice1))
}
