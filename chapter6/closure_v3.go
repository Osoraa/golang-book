package main

// Demonstrtes a function returning a function

import "fmt"

func test() func() int {
	i := int(0)

	// Returns a function that increments
	return func() int {
		i += 2
		return i
	}
}

func main() {
	tes := test()
	fmt.Println(tes())
	fmt.Println(tes())
	fmt.Println(tes())
}