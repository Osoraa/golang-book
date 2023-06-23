package main

import "fmt"

func main() {
	// Collects a number as input and doubles it. No type checking yet!!

	var num rune

	fmt.Println("Enter a number: ")
	fmt.Scanf("%d", &num)
	fmt.Println(num * 2)
}
