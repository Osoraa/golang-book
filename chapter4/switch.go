package main

import "fmt"

func main() {
	// Using the Switch statement
	// collects a number and Prints it's digits backwards
	var num rune

	fmt.Println("Enter a number:")
	fmt.Scanf("%d", &num)
	fmt.Println("The digits backwards are:")

	for ; num > 0; num /= 10 {
		i := num % 10

		switch i {
		case 0: fmt.Println("Zero")
		case 1: fmt.Println("One")
		case 2: fmt.Println("Two")
		case 3: fmt.Println("Three")
		case 4: fmt.Println("Four")
		case 5: fmt.Println("Five")
		case 6: fmt.Println("Six")
		case 7: fmt.Println("Seven")
		case 8: fmt.Println("Eight")
		case 9: fmt.Println("Nine")
		default: fmt.Println("Unknown number")
		}
	}
}