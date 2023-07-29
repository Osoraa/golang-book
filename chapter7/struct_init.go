package main

// Demonstrating Structs in Go

import ("fmt"; "math")

// Declaring a Struct
type Circle struct {
	x, y, r float64
}

func circleArea(c *Circle) float64 {
	// Calculates the area of a circle
	return math.Pi * c.r * c.r
}

func main() {
	// Instantiating a Struct
	var circ Circle
	circl := new(Circle)
	cir := Circle{x: 5, y: 5, r: 2}
	cirl := &Circle{4, 4, 1}	// & to get pointer to the struct

	// Setting Struct values
	circ.x = 10
	circ.y = 11
	circ.r = 5

	circl.x = 20
	circl.y = 21
	circl.r = 10

	fmt.Println(circ)
	fmt.Println(*circl) // Dereferenced pointer
	fmt.Println(cir)
	fmt.Println(*cirl)	// Dereferencing because pointer to struct was created

	fmt.Println(circleArea(&circ))
}
