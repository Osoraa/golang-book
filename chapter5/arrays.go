package main

import "fmt"

func main() {
	// Declares an array of type int and assigns the last element 100
	var arr [5]int

	arr[4] = 100

	fmt.Println(arr)
}