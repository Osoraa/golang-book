package main

// Demonstrating how to recover from a panic

import "fmt"

func main() {
	defer func() {
		str := recover()
		fmt.Println(str)
	}()

	panic("Not yet Implemented!!!")
}
