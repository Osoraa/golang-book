package main

import "fmt"

func main() {
	// Collects a number as input and doubles it. No type checking yet!!

	var num rune
	// num := new(rune)

	fmt.Println("Enter a number: ")

	fmt.Scanf("%d", &num)
	// fmt.Scanf("%d", num)

	fmt.Println(num * 2)
	// fmt.Println(*num * 2)
}
