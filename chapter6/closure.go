package main

// Demonstrating nested finctions

import "fmt"

func main() {
	// Add function stored as a variable.

	// Here, a single type_name is used for all vars, has to come last.
	add := func (x, y, z int) int {
		return x + y + z
	}

	fmt.Println(add(3, 4, 5))
}
