package main

// Demonstrating Recursion in Go.

import "fmt"

func factorial(x int) int {
	if x == 1 {
		return 1
	}
	return x * factorial(x-1)
}

func main() {
	fmt.Println((factorial(6)))
	fmt.Println((factorial(5)))
}
