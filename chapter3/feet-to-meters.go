package main

import "fmt"

func main() {
	// Basic Function to convert Feet measurement to metres.
	// Conversion is 1ft = 0.3048 meters.

	var ft float64

	fmt.Println("Please input a measurement in foot:")
	fmt.Scanf("%f", &ft)

	fmt.Println(ft, "ft is", ft * 0.3048, "in metres.")
}