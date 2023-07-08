package main

// Deferring functions in Go.

import "fmt"

func name() {
	fmt.Println("I'm a Go Dev!!")
}

func deferred() {
	fmt.Println("And I deferred a func!")
}

func hello() {
	fmt.Println("Hello")
}

func main() {
	defer deferred()
	defer name()
	hello()
}
