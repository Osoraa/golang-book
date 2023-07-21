package main

// Generate the fibonacci sequence recursively.

import "fmt"

func fib(num int) int {
	if num == 0 {
		return 0
	}

	if num == 1 {
		return 1
	}

	return fib(num - 1) + fib(num - 2)
}

func main() {
	fmt.Println(fib(7))
}
