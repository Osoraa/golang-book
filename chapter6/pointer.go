package main

// Pointers in Go

import "fmt"

func one(xPtr *int) {
	*xPtr = 100
}

func main() {
	/* x := 0			// Normal way to do it
	one(&x) */

	x := new(int)	// The Go way
	one(x)

	fmt.Println(*x)
}