package main

// Generates odd numbers - Exercise 4

import "fmt"

func makeOddGenerator() (oddGen func() int) {
	// Returns a function named odd_gen

	x := 1

	// Function variable
	oddGen = func() int {
		ret := x
		x += 2
		return ret
	}

	// Lone return due to named return variable in function signature
	return
}

func main() {
	nextOdd := makeOddGenerator()

	fmt.Println(nextOdd())
	fmt.Println(nextOdd())
	fmt.Println(nextOdd())
}