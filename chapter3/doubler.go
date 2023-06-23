package main

import "fmt"

func main() {
	fmt.Println("Enter a number: ")
	var num rune
	fmt.Scanf("%d", &num)
	fmt.Println(num * 2)
}
