package main

// Swap two integers, Pointer style.

import "fmt"

func swap(xPtr *int, yPtr *int) {
	// Swap Integers
	temp := *yPtr
	*yPtr = *xPtr
	*xPtr = temp
	
}

func main() {
	x, y := 1, 2

	fmt.Println(x, y)
	swap(&x, &y)
	fmt.Println(x, y)
}
