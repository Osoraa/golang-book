package main

import "fmt"

func main() {
	// Prints Fizz for numbers divisible by only 3,
	// Buzz for numbers divisible by only 5,
	// and FizzBuzz for numbers divisible by both.

	for i:=1; i<=100; i++ {
		if i % 3 == 0 && i % 5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i % 5 == 0 {
			fmt.Println("Buzz")
		} else if i % 3 == 0 {
			fmt.Println("Fizz")
		} else {
			fmt.Println(i)
		}
	}
}