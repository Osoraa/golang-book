package main

// Demonstrates methods in Go.

import ("fmt")

type Rectangle struct {
	l, b int
}

func (rec *Rectangle) area() int {
	// Area method made callable on a Recangle object.
	return rec.l * rec.b
}

func main() {
	rec := new(Rectangle)
	rec.l, rec.b = 10, 5

	rect := Rectangle{l: 20, b: 10}

	fmt.Println(rec.area())
	fmt.Println(rect.area())
}