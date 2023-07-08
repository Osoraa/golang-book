package main

// Demonstrating closures.

import "fmt"

func main() {
	// Main function.

	// The decrement function and i variable form a closure
	i := uint(0)
	decrement := func () uint {
		i--;
		return i
	}

	fmt.Println(decrement())
	fmt.Println(decrement())

}
