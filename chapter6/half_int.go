package main

// Exercise 2 Chapter 6

import "fmt"

func half(num rune) (rune, bool) {
	// Halves an integer, returns the result and true if num is even.

	return num / 2, num % 2 == 0
}

func main() {
	// Creates a pointer to an int32
	num := new(rune)

	// Requests a number
	fmt.Print("Input a number: ")
	fmt.Scanf("%d", num)

	// Prints result of half func
	fmt.Println(half(*num))
}