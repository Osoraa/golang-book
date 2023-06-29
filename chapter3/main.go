// A Hello World Program

// 
package main

import "fmt"

func main() {
	// Assign multiple variables, just like unpacking in Python 
	name, greet:= "Osoraa", "Hello"

	const immut = "Bookie"

	// Throws an error
	// immut = "Osoraa"

	fmt.Println(greet, "World from Introduction to Go by", name)
}
