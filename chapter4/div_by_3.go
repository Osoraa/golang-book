package main

import "fmt"

func main() {
	// Prints numbers divisible by 3 between 1 and 100

	for i := 1; i <= 100; i++ {
		if i % 3 == 0 {
			fmt.Print(i, " ")
		}
	}
	fmt.Println()
}